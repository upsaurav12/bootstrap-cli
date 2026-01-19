/*
Copyright © 2025 Saurav Upadhyay
*/
package cmd

import (
	"fmt"
	"os"
	"runtime/debug"

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
	Long: `BootstrapCLI is a project lifecycle tool for Go services.
It helps you create, configure, and evolve Go projects using
opinionated defaults and a clear project structure.`,
	Run: func(cmd *cobra.Command, args []string) {
		v, _ := cmd.Flags().GetBool("version")
		if v {
			fmt.Print(cmd.VersionTemplate())
			return
		}
		_ = cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func resolvedVersion() string {
	if info, ok := debug.ReadBuildInfo(); ok {
		if info.Main.Version != "(devel)" {
			return info.Main.Version
		}
	}
	return Version
}

func init() {
	rootCmd.Flags().BoolP("version", "v", false, "print version information")

	rootCmd.SetVersionTemplate(
		fmt.Sprintf(
			"bootstrap-cli %s\nCommit: %s\nBuilt:  %s\n",
			resolvedVersion(),
			Commit,
			Date,
		),
	)
}
