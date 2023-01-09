package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//Read input file
	input, _ := os.Open("input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	// A : X Rock
	// B : Y Paper
	// C : Z Scissors
	var scoreOne, scoreTwo int
	// Scores for part one
	scoresOne := map[string]int{"B X": 1, "C Y": 2, "A Z": 3, "A X": 4, "B Y": 5, "C Z": 6, "C X": 7, "A Y": 8, "B Z": 9}
	// Scores for part two
	scoresTwo := map[string]int{"B X": 1, "C X": 2, "A X": 3, "A Y": 4, "B Y": 5, "C Y": 6, "C Z": 7, "A Z": 8, "B Z": 9}

	for sc.Scan() {
		scoreOne += scoresOne[sc.Text()]
		scoreTwo += scoresTwo[sc.Text()]
	}

	fmt.Println("Part One:", scoreOne)
	fmt.Println("Part Two:", scoreTwo)
}
