/*
Copyright © 2025 Saurav Upadhyay
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	Version = "dev"
	Commit  = "none"
	Date    = "unknown"
)

var rootCmd = &cobra.Command{
	Use:   "bootstrap",
	Short: "Scaffold and manage production-ready Go services.",
	Long: `BootstrapCLI is a project lifecycle tool for Go services,
It helps you create, configure, and evolve Go projects using
opinionated defaults, a clear project structure, and
declarative configuration.`,
	Version: Version,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.SetVersionTemplate(
		fmt.Sprintf(
			"bootstrap-cli %s\nCommit: %s\nBuilt:  %s\n",
			Version,
			Commit,
			Date,
		),
	)
}
