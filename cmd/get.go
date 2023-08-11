package cmd

import (
	"fmt"
	"github.com/opa-oz/over/pkg/config"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get package version",
	Long:  `Get package version from .over.yaml file of current directory`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.ParseConfig()

		if err != nil {
			return err
		}

		currentVersion := cfg.Package.Version

		fmt.Println(currentVersion)
		return nil
	},
}
