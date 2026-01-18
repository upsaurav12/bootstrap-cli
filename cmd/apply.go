package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/upsaurav12/bootstrap/pkg/parser"
)

var applyCmd = cobra.Command{
	Use:   "apply",
	Short: "Apply project configuration from YAML",
	Run: func(cmd *cobra.Command, args []string) {

		if yamlPath == "" {
			fmt.Fprintln(cmd.OutOrStdout(), "Error: --yaml is required")
			return
		}

		yamlConfig, err := parser.ReadYAML(yamlPath)
		if err != nil {
			fmt.Fprintln(cmd.OutOrStdout(), "error reading yaml:", err)
			return
		}

		YAMLPath = yamlPath
		projectRouter = yamlConfig.Project.Router
		projectPort = strconv.Itoa(yamlConfig.Project.Port)
		DBType = yamlConfig.Project.Database
		Entities = yamlConfig.Entities

		createNewProject(
			yamlConfig.Project.Name,
			projectRouter,
			yamlConfig.Project.Type,
			cmd.OutOrStdout(),
		)
	},
}

var yamlPath string

type Config struct {
	Project     Project  `yaml:"project"`
	Entities    []string `yaml:"entities"`
	CustomLogic []string `yaml:"custom_logic"`
}

type Project struct {
	Name     string `yaml:"name"`
	Type     string `yaml:"type"`
	Port     int    `yaml:"port"`
	Location string `yaml:"location"`
	Database string `yaml:"db"`
	Router   string `yaml:"router"`
}

func init() {

	applyCmd.Flags().StringVar(&yamlPath, "yaml", "", "yaml configuation")
	rootCmd.AddCommand(&applyCmd)
}
