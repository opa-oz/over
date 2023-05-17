package cmd

import (
	"fmt"
	"github.com/hashicorp/go-version"
	"github.com/opa-oz/over/pkg/config"
	"github.com/opa-oz/over/pkg/fileutils"
	"github.com/opa-oz/over/pkg/versionutils"
	"strings"

	"github.com/spf13/cobra"
)

var (
	patch   = false
	minor   = false
	major   = false
	inplace = false
)

var b2i = map[bool]int8{false: 0, true: 1}

var upCmd = &cobra.Command{
	Use:   "up",
	Short: "Up package version",
	Long:  `TBD`,
	RunE: func(cmd *cobra.Command, args []string) error {

		sum := b2i[minor] + b2i[major] + b2i[patch]

		if sum > 1 {
			return fmt.Errorf("please select ONLY one of (patch/minor/major) flags")
		}
		if sum == 0 {
			return fmt.Errorf("please select one of patch/minor/major flags")
		}

		cfg, err := config.ParseConfig()

		if err != nil {
			return err
		}

		currentVersion := cfg.Package.Version

		if Verbose {
			fmt.Println("Current version:", currentVersion)
		}

		var v *version.Version

		v, err = version.NewSemver(currentVersion)
		if err != nil {
			return err
		}

		v, err = versionutils.Increase(v, patch, minor, major)
		if err != nil {
			return err
		}

		if !inplace {
			fmt.Println(fmt.Sprintf("%s", v.String()))
			return nil
		}

		files := append(cfg.Package.Files, config.File{
			Name: ".over.yaml",
			Templates: []string{
				"version: __VERSION__",
				"version: '__VERSION__'",
				`version: "__VERSION__"`,
			},
		})

		for _, file := range files {
			if Verbose {
				fmt.Println(fmt.Sprintf("Preparing templates for %s", file.Name))
			}

			replacements := make([]fileutils.Replacement, len(file.Templates))

			for index, template := range file.Templates {
				replacements[index] = fileutils.Replacement{
					To:   strings.ReplaceAll(template, "__VERSION__", v.String()),
					From: strings.ReplaceAll(template, "__VERSION__", currentVersion),
				}

				if Verbose {
					fmt.Println(fmt.Sprintf("\t[%s] -> [%s]", replacements[index].From, replacements[index].To))
				}
			}

			if len(replacements) > 0 {
				err = fileutils.ReplaceInFile(file.Name, replacements)
				if err != nil {
					return err
				}
			}
		}

		return nil
	},
}
