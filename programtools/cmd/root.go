/*
Copyright © 2025 Arvid Kristoffersson arvid.kristoffersson@icloud.com
*/

package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "programtools",
	Short: "Tools for problem creation for Hundis",
	Long: `Programtools contains CLI Tools that make problem creation
and verification for Hundis easier.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
