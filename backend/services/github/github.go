package github

import (
	"encoding/json"
	"errors"
	"hundis/config"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type ErrorResponse struct {
	ErrorName        string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
}

type UserResponse struct {
	Login     string `json:"login"`
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	AvatarURL string `json:"avatar_url"`
}

type EmailResponse []struct {
	Email      string `json:"email"`
	Primary    bool   `json:"primary"`
	Verified   bool   `json:"verified"`
	Visibility string `json:"visibility,omitempty"`
}

type UserInfo struct {
	ID           int
	Username     string
	Name         string
	PrimaryEmail string
	AvatarURL    string
}

func HandleAPIResponse(resp *http.Response) ([]byte, error) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var errorResp ErrorResponse
	if err := json.Unmarshal(body, &errorResp); err == nil && errorResp.ErrorName != "" {
		return nil, errors.New(errorResp.ErrorDescription)
	}

	return body, nil
}

func RequestToken(code string) (*TokenResponse, error) {
	config := config.Config()

	data := url.Values{}
	data.Set("client_id", config.GitHubClientID)
	data.Set("client_secret", config.GitHubClientSecret)
	data.Set("code", code)

	req, err := http.NewRequest(http.MethodPost, "https://github.com/login/oauth/access_token", strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := HandleAPIResponse(resp)
	if err != nil {
		return nil, err
	}

	var tokenResponse TokenResponse
	if err := json.Unmarshal(body, &tokenResponse); err != nil {
		return nil, err
	}

	if tokenResponse.AccessToken == "" {
		return nil, errors.New("no access token returned from GitHub")
	}

	return &tokenResponse, nil
}

func GetUserInfo(accessToken string) (*UserInfo, error) {
	user, err := getUserProfile(accessToken)
	if err != nil {
		return nil, err
	}

	emails, err := getUserEmails(accessToken)
	if err != nil {
		return nil, err
	}

	primaryEmail := ""
	for _, email := range emails {
		if email.Primary && email.Verified {
			primaryEmail = email.Email
			break
		}
	}

	if primaryEmail == "" {
		return nil, errors.New("GitHub primary email not found or not verified")
	}

	return &UserInfo{
		Username:     user.Login,
		ID:           user.ID,
		Name:         user.Name,
		AvatarURL:    user.AvatarURL,
		PrimaryEmail: primaryEmail,
	}, nil
}

func getUserProfile(accessToken string) (*UserResponse, error) {
	req, err := http.NewRequest(http.MethodGet, "https://api.github.com/user", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := HandleAPIResponse(resp)
	if err != nil {
		return nil, err
	}

	var userResponse UserResponse
	if err := json.Unmarshal(body, &userResponse); err != nil {
		return nil, err
	}

	return &userResponse, nil
}

func getUserEmails(accessToken string) (EmailResponse, error) {
	req, err := http.NewRequest(http.MethodGet, "https://api.github.com/user/emails", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := HandleAPIResponse(resp)
	if err != nil {
		return nil, err
	}

	var emailResponse EmailResponse
	if err := json.Unmarshal(body, &emailResponse); err != nil {
		return nil, err
	}

	return emailResponse, nil
}
