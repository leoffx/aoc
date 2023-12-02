package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func getFirstDigit(line string) (string, bool) {
	for _, val := range line {
		if unicode.IsDigit(val) {
			return string(val), true
		}
	}
	return "", false
}

func getLastDigit(line string) (string, bool) {
	for i := len(line) - 1; i >= 0; i-- {
		if unicode.IsDigit(rune(line[i])) {
			return string(line[i]), true
		}
	}
	return "", false
}

func convertSpelledOutNumbers(input *string) {
	spelledOutNumbers := [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	for _, number := range spelledOutNumbers {
		if strings.Contains(*input, number) {
			*input = strings.ReplaceAll(*input, number, fmt.Sprintf("%s%s%s", number, spelledOutNumbertoNumeral(number), number))
		}
	}
	for _, number := range spelledOutNumbers {
		if strings.Contains(*input, number) {
			*input = strings.ReplaceAll(*input, number, "")
		}
	}
}

func spelledOutNumbertoNumeral(number string) string {
	translation := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
	return translation[number]
}

func firstPart() {
	inputFileName := "input.txt"
	content, err := os.ReadFile(inputFileName)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(content), "\n")
	sum := 0
	for _, line := range lines {
		firstDigit, ok := getFirstDigit(line)
		if !ok {
			panic("no number found")
		}
		lastDigit, ok := getLastDigit(line)
		if !ok {
			panic("no number found")
		}
		finalDigit, err := strconv.Atoi(firstDigit + lastDigit)
		if err != nil {
			panic("could not convert combined digits to int")
		}
		sum += finalDigit
	}
	fmt.Println("Part 1 Answer: ", sum)
}

func secondPart() {
	inputFileName := "input.txt"
	content, err := os.ReadFile(inputFileName)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(content), "\n")
	sum := 0
	for _, line := range lines {
		convertSpelledOutNumbers((&line))
		firstDigit, ok := getFirstDigit(line)
		if !ok {
			panic("no number found")
		}
		lastDigit, ok := getLastDigit(line)
		if !ok {
			panic("no number found")
		}
		finalDigit, err := strconv.Atoi(firstDigit + lastDigit)
		if err != nil {
			panic("could not convert combined digits to int")
		}
		sum += finalDigit
	}
	fmt.Println("Part 2 Answer: ", sum)
}

func main() {
	firstPart()
	secondPart()
}
