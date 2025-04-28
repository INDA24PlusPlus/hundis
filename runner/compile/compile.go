package compile

import (
	"fmt"
	"os"
	"runner/nsjail"

	"github.com/google/uuid"
)

type Environment struct {
	SourceFile *os.File
	SourceName string
	Path       string
	OutputName string
}

func NewEnvironment(sourceName, outputName string) (*Environment, error) {
	path, err := os.MkdirTemp("", "*")
	if err != nil {
		return nil, err
	}

	file := path + "/" + sourceName
	f, err := os.Create(file)
	if err != nil {
		return nil, err
	}

	return &Environment{
		SourceName: sourceName,
		SourceFile: f,
		Path:       path,
		OutputName: outputName,
	}, nil
}

func (ce *Environment) WriteCode(code string) error {
	defer ce.SourceFile.Close()
	_, err := ce.SourceFile.WriteString(code)
	if err != nil {
		return err
	}
	return nil
}

func (ce *Environment) CleanUp() {
	if ce.SourceFile != nil {
		ce.SourceFile.Close()
	}
	if ce.Path != "" {
		os.RemoveAll(ce.Path)
	}
}

func (ce *Environment) GetCompiledCode() ([]byte, error) {
	out, err := os.ReadFile(ce.Path + "/" + ce.OutputName)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func Compile(code string, config *Config) (string, error) {
	if !config.ShouldCompile {
		return "", nil
	}

	env, err := NewEnvironment(config.SourceName, config.OutputName)
	if err != nil {
		return "", err
	}
	defer env.CleanUp()
	if err := env.WriteCode(code); err != nil {
		return "", err
	}

	out, err := nsjail.RunNsjail(config.NsjailConfigName, env.Path)
	fmt.Println("nsjail:\n", string(out))
	if err != nil {
		return "", err
	}

	data, err := env.GetCompiledCode()
	if err != nil {
		return "", err
	}

	uuid := uuid.New().String()
	path := "/compiled/" + uuid

	if err := os.WriteFile(path, data, 0644); err != nil {
		return "", err
	}

	return path, nil
}
