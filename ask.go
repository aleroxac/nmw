package main

import (
	"fmt"
)

func ask(sourceLang string, destineLang string, minDifficulty string, maxDifficulty string, translation map[string]string) int {
	fmt.Printf("%s: %v", sourceLang, translation[sourceLang])
	fmt.Print("\nPress 'Enter' to see the translation...")

	var enter string
	fmt.Scanln(&enter)
	fmt.Printf("%s: %v", destineLang, translation[destineLang])

	var difficultyChoice int
	fmt.Printf("\nDifficulty (%s-%s): ", minDifficulty, maxDifficulty)
	fmt.Scan(&difficultyChoice)
	fmt.Print("---------------------------------------------------------------------------------------\n")

	return difficultyChoice
}
