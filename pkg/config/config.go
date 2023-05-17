package config

import (
	"github.com/spf13/viper"
)

type File struct {
	Name      string   `yaml:"name"`
	Templates []string `yaml:"templates"`
}

type Package struct {
	Name      string `yaml:"name"`
	Version   string `yaml:"version"`
	IsDefault bool   `yaml:"default"`
	Files     []File `yaml:"files,mapstructure"`
}

type Config struct {
	Package Package `yaml:"package"`
}

func ParseConfig() (*Config, error) {
	config := Config{}

	viper.SetDefault("package.default", true)

	err := viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
