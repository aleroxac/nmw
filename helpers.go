package main

import (
	"fmt"
	"reflect"
	"time"

	"gopkg.in/yaml.v3"
)

func getReviewPeriod(difficulty int) string {
	switch difficulty {
	case 1:
		return "0d:0h:5m"
	case 2:
		return "0d:0h:4m"
	case 3:
		return "0d:0h:3m"
	case 4:
		return "0d:0h:2m"
	default:
		return "0d:0h:1m"
	}
}

func addReviewPeriod(reviewPeriod string, lastReviewed time.Time) time.Time {
	var days, hours, minutes int
	fmt.Sscanf(reviewPeriod, "%dd:%dh:%dm", &days, &hours, &minutes)
	duration := time.Duration(days)*24*time.Hour + time.Duration(hours)*time.Hour + time.Duration(minutes)*time.Minute
	return lastReviewed.Add(duration)
}

func cardExistsInOutputDeck(deck Deck, output Output) (bool, int) {
	for idx := range output.Cards {
		deckSourceLang, deckDestineLang := discoverDeckLangs()
		outputSourceLang, outputDestineLang := discoverOutputLangs()

		if reflect.DeepEqual(deckSourceLang, outputSourceLang) && reflect.DeepEqual(deckDestineLang, outputDestineLang) {
			return true, idx
		}
	}
	return false, -1
}

// quoteNode recursively sets the style of the YAML node to be double quoted for values only
func quoteNode(node *yaml.Node) {
	if node.Kind == yaml.MappingNode {
		for i, child := range node.Content {
			// We only want to quote the values, not the keys.
			// In a mapping node, the even indices are keys and the odd indices are values.
			if i%2 == 1 && child.Kind == yaml.ScalarNode && !child.IsZero() {
				// If the parent key is "difficulty", don't quote
				if node.Content[i-1].Value != "difficulty" {
					child.Style = yaml.DoubleQuotedStyle
				}
			}
			quoteNode(child)
		}
	} else {
		for _, child := range node.Content {
			quoteNode(child)
		}
	}
}
