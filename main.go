package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"gopkg.in/yaml.v3"
)

func main() {
	// Always read config.yaml
	config := loadConfigFile()
	// fmt.Printf("config:\t%v\n", config)

	// Always read deck.yaml
	deck := loadDeckFile(config)
	// fmt.Printf("deck:\t%v\n", deck)

	// If output.yaml exists, read it
	output := loadOutputFile(config)

	// get langs
	sourceLang, destineLang := discoverDeckLangs()

	// get min and max gifficulties
	minDifficulty, maxDifficulty := getDifficulties(config)

	now := time.Now()
	allCardsAfterNextReview := true

	for _, card := range deck.Cards {
		exists, index := cardExistsInOutputDeck(deck, *output)

		if exists {
			if output.Cards[index].NextReview.After(now) {
				continue
			}

			// Update difficulty if card exists in outputDeck
			difficultyInt := ask(sourceLang, destineLang, minDifficulty, maxDifficulty, card.Translation)

			// fmt.Printf("%s: %s", sourceLang, card.Translation[sourceLang])
			// fmt.Print("\nPress 'Enter' to see the translation...")
			// var enter string
			// fmt.Scanln(&enter)
			// fmt.Printf("%s: %s", destineLang, card.Translation[destineLang])

			// var difficultyInt int
			// fmt.Printf("\nDifficulty (%s-%s): ", minDifficulty, maxDifficulty)
			// fmt.Scan(&difficultyInt)

			output.Cards[index].Difficulty = strconv.Itoa(difficultyInt)
			oldDifficulty := output.Cards[index].Difficulty
			difficultyStr := strconv.Itoa(difficultyInt)
			if oldDifficulty != difficultyStr {
				output.Cards[index].ReviewPeriod = getReviewPeriod(difficultyInt)
			}
			output.Cards[index].LastReviewed = now
			output.Cards[index].NextReview = addReviewPeriod(output.Cards[index].ReviewPeriod, now)
		} else {
			// Add card if it doesn't exist in outputDeck
			difficultyInt := ask(sourceLang, destineLang, minDifficulty, maxDifficulty, card.Translation)
			// fmt.Printf("%s: %s", sourceLang, card.Translation[sourceLang])
			// fmt.Print("\nPress 'Enter' to see the translation...")
			// var enter string
			// fmt.Scanln(&enter)
			// fmt.Printf("%s: %s", destineLang, card.Translation[destineLang])

			// var difficultyInt int
			// fmt.Printf("\nDifficulty (%s-%s): ", minDifficulty, maxDifficulty)
			// fmt.Scan(&difficultyInt)

			newCard := CardOutput{
				Translation:  card.Translation,
				Difficulty:   strconv.Itoa(difficultyInt),
				ReviewPeriod: getReviewPeriod(difficultyInt),
				LastReviewed: now,
				NextReview:   addReviewPeriod(getReviewPeriod(difficultyInt), now),
			}
			output.Cards = append(output.Cards, newCard)
		}
	}

	// Save to output.yaml
	outputData, err := yaml.Marshal(&output)
	if err != nil {
		fmt.Println("Error marshalling to yaml:", err)
		return
	}

	// Correct the output to ensure the values are enclosed in double quotes
	node := &yaml.Node{}
	yaml.Unmarshal(outputData, node)
	quoteNode(node)
	correctedOutput, err := yaml.Marshal(node)
	if err != nil {
		fmt.Printf("Error to marshall output: %v", err)
	}

	err = os.WriteFile("output.yaml", correctedOutput, 0644)
	if err != nil {
		fmt.Printf("Error writing to %s: %v", "output.yaml", err)
	}

	if allCardsAfterNextReview {
		fmt.Println("You've done a great job! Now it's time to rest.")
		return
	}
}
