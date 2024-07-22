package main

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type CardDeck struct {
	Translation map[string]string
}

type Deck struct {
	Cards []CardDeck `yaml:"cards"`
}

func loadDeckFile(config Config) Deck {
	_, err := os.Stat("deck.yaml")
	if err != nil {
		log.Fatalf("Error reading YAML file: %v\n", err)
	}

	deckFile, err := os.ReadFile("deck.yaml")
	if err != nil {
		log.Fatalf("Error reading YAML file: %v\n", err)
	}

	var rawDeck struct {
		Cards []map[string]string `yaml:"cards"`
	}
	err = yaml.Unmarshal(deckFile, &rawDeck)
	if err != nil {
		log.Fatalf("Error parsing YAML: %v\n", err)
	}

	var deck Deck
	for _, rawCard := range rawDeck.Cards {
		card := CardDeck{Translation: rawCard}
		deck.Cards = append(deck.Cards, card)
	}

	return deck
}

func discoverDeckLangs() (string, string) {
	data, err := os.ReadFile("deck.yaml")
	if err != nil {
		log.Fatalf("Error reading YAML file: %v\n", err)
	}

	var rawDeck struct {
		Cards []map[string]string `yaml:"cards"`
	}
	err = yaml.Unmarshal(data, &rawDeck)
	if err != nil {
		log.Fatalf("Error parsing YAML: %v\n", err)
	}

	var deck Deck
	for _, rawCard := range rawDeck.Cards {
		card := CardDeck{Translation: rawCard}
		deck.Cards = append(deck.Cards, card)
	}

	// Assuming the first card has the structure we want:
	sourceLang := "unknown"
	destineLang := "unknown"
	for k := range deck.Cards[0].Translation {
		if sourceLang == "unknown" {
			sourceLang = k
		} else {
			destineLang = k
			break
		}
	}

	return sourceLang, destineLang
}
