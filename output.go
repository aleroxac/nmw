package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type CardOutput struct {
	Translation  map[string]string
	Difficulty   string    `yaml:"difficulty,omitempty"`
	ReviewPeriod string    `yaml:"review_period,omitempty"`
	LastReviewed time.Time `yaml:"last_reviewed,omitempty"`
	NextReview   time.Time `yaml:"next_review,omitempty"`
}

type Output struct {
	Cards []CardOutput `yaml:"cards"`
}

func loadOutputFile(config Config) *Output {
	var output Output
	outputFileAbsolutePath := fmt.Sprintf("%s/%s.%s", config.ConfigOutput.FilePath, config.ConfigOutput.FileName, config.ConfigOutput.FileFormat)

	if _, err := os.Stat(outputFileAbsolutePath); err == nil {
		outputFile, err := os.ReadFile(outputFileAbsolutePath)
		if err != nil {
			// fmt.Printf("Error reading %s: %v", outputFileAbsolutePath, err)
			return &output
		}

		err = yaml.Unmarshal(outputFile, &output)
		if err != nil {
			fmt.Printf("Error unmarshalling %s: %v", outputFileAbsolutePath, err)
		}
	} else {
		return &output
	}
	return &output
}

func discoverOutputLangs() (string, string) {
	data, err := os.ReadFile("output.yaml")
	if err != nil {
		// log.Fatalf("Error reading YAML file: %v\n", err)
		return "", ""
	}

	var rawDeck struct {
		Cards []struct {
			Translation  map[string]string
			Difficulty   string
			ReviewPeriod string
			LastReviewed time.Time
			NextReview   time.Time
		}
	}

	err = yaml.Unmarshal(data, &rawDeck)
	if err != nil {
		log.Fatalf("Error parsing YAML: %v\n", err)
	}

	var output Output
	for _, rawCard := range rawDeck.Cards {
		card := CardOutput{Translation: rawCard.Translation}
		output.Cards = append(output.Cards, card)
	}

	// Assuming the first card has the structure we want:
	sourceLang := "unknown"
	destineLang := "unknown"
	for k := range output.Cards[0].Translation {
		if sourceLang == "unknown" {
			sourceLang = k
		} else {
			destineLang = k
			break
		}
	}

	return sourceLang, destineLang
}
