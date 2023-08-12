package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Get `over` version",
	Long:  `Get version of package you are using right now`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("0.1.0")
		return nil
	},
}
