package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "over",
	Short: "Control version everywhere",
	Long:  `TBD`,
}

var (
	cfgFile = ""
	Verbose = false
)

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is .over.yaml)")
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")

	upCmd.Flags().BoolVarP(&patch, "patch", "p", false, "increase patch version")
	upCmd.Flags().BoolVarP(&minor, "minor", "m", false, "increase minor version")
	upCmd.Flags().BoolVarP(&major, "major", "M", false, "increase major version")
	upCmd.Flags().BoolVarP(&inplace, "inplace", "i", false, "change version in all the files")

	rootCmd.AddCommand(upCmd)
	rootCmd.AddCommand(getCmd)
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigName(".over")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(".")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		if Verbose {
			fmt.Println("Using config file:", viper.ConfigFileUsed())
		}
	}
}
