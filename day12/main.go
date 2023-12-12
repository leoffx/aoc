package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"
)

func generateAllPossibilities(springs []string) [][]string {
	if len(springs) == 0 {
		return [][]string{{}}
	}

	var helper func([]string, int) [][]string
	helper = func(springs []string, index int) [][]string {
		if index == len(springs) {
			return [][]string{append([]string{}, springs...)}
		}

		if springs[index] == "?" {
			withHash := append([]string{}, springs...)
			withHash[index] = "#"
			withDot := append([]string{}, springs...)
			withDot[index] = "."

			return append(helper(withHash, index+1), helper(withDot, index+1)...)
		}

		return helper(springs, index+1)
	}

	return helper(springs, 0)
}

func countDamages(row []string) []string {
	var damages []string
	count := 0
	for _, c := range row {
		if c == "#" {
			count += 1
		}
		if c == "." && count > 0 {
			damages = append(damages, fmt.Sprintf("%d", count))
			count = 0
		}
	}
	if count > 0 {
		damages = append(damages, fmt.Sprintf("%d", count))
	}
	return damages
}

func firstPart() {
	inputFileName := "input.txt"
	content, err := os.ReadFile(inputFileName)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(content), "\n")
	ans := 0
	for _, line := range lines {
		lineSplit := strings.Split(line, " ")
		if len(lineSplit) != 2 {
			panic("could not parse line")
		}
		springs := strings.Split(lineSplit[0], "")
		damages := strings.Split(lineSplit[1], ",")
		possibilities := generateAllPossibilities(springs)
		for _, possibility := range possibilities {
			c := countDamages(possibility)
			if reflect.DeepEqual(c, damages) {
				ans += 1
			}

		}
	}
	fmt.Println("Part 1 Answer: ", ans)
}

func secondPart() {
	inputFileName := "input.txt"
	content, err := os.ReadFile(inputFileName)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(content), "\n")
	fmt.Println("Part 2 Answer: ", lines)
}

func main() {
	firstPart()
	// secondPart()
}
