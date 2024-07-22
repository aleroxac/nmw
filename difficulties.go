package main

func getDifficulties(config Config) (string, string) {
	var minDifficulty string
	var maxDifficulty string
	for level := range config.ConfigDifficultyLevels.Level {
		if minDifficulty == "" || level < minDifficulty {
			minDifficulty = level
		}
		if maxDifficulty == "" || level > maxDifficulty {
			maxDifficulty = level
		}
	}

	// var minDifficultyInt int
	// if minDifficultyInt, err := strconv.Atoi(minDifficulty); err != nil {
	// 	fmt.Printf("Error convert minDifficulty(%d) from string to int: %v", minDifficultyInt, err)
	// }

	// var maxDifficultyInt int
	// if maxDifficultyInt, err := strconv.Atoi(minDifficulty); err != nil {
	// 	fmt.Printf("Error convert maxDifficultyInt(%d) from string to int: %v", maxDifficultyInt, err)
	// }

	return minDifficulty, maxDifficulty
}
