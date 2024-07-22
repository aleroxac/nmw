package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type ConfigDeck struct {
	FilePath   string `yaml:"file_path,omitempty"`
	FileName   string `yaml:"file_name,omitempty"`
	FileFormat string `yaml:"file_format,omitempty"`
}

type ConfigOutput struct {
	FilePath   string `yaml:"file_path,omitempty"`
	FileName   string `yaml:"file_name,omitempty"`
	FileFormat string `yaml:"file_format,omitempty"`
}

type DifficultyLevel struct {
	ReviewPeriod string `yaml:"review_period,omitempty"`
}

type ConfigDifficultyLevels struct {
	Level map[string]DifficultyLevel `yaml:",inline"`
}

type Config struct {
	ConfigDeck             `yaml:"deck,omitempty"`
	ConfigOutput           `yaml:"output,omitempty"`
	ConfigDifficultyLevels `yaml:"difficulty_levels,omitempty"`
}

func loadConfigFile() Config {
	var config Config
	if _, err := os.Stat("config.yaml"); err == nil {
		configFile, err := os.ReadFile("config.yaml")
		if err != nil {
			fmt.Println("Error reading config.yaml:", err)
		}

		err = yaml.Unmarshal(configFile, &config)
		if err != nil {
			fmt.Println("Error unmarshalling config.yaml:", err)
		}
	}
	return config
}
