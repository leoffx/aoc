package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func convertAtoiSlice(slice *[]string) ([]int, error) {
	response := make([]int, 0)
	for _, char := range *slice {
		if char == "" {
			continue
		}
		num, err := strconv.Atoi(strings.TrimSpace(char))
		if err != nil {
			return nil, fmt.Errorf("char is not a number: %s", err)
		}
		response = append(response, num)
	}
	return response, nil
}

func countWinningNumbers(chars, winningChars *[]string) (int, error) {
	numCount := 0
	nums, err := convertAtoiSlice(chars)
	if err != nil {
		return 0, err
	}
	winningNums, err := convertAtoiSlice(winningChars)
	if err != nil {
		return 0, err
	}
	for _, num := range nums {
		for _, winningNum := range winningNums {
			if num == winningNum {
				numCount += 1
				break
			}
		}

	}
	return numCount, nil
}

func firstPart() {
	inputFileName := "input.txt"
	content, err := os.ReadFile(inputFileName)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(content), "\n")
	score := 0
	for _, line := range lines {
		cardSplit := strings.Split(line, ": ")
		if len(cardSplit) != 2 {
			panic("malformed game line " + line)
		}
		gameNumbers := cardSplit[1]
		gameNumberSplit := strings.Split(gameNumbers, " | ")
		if len(gameNumberSplit) != 2 {
			panic("malformed game numbers " + line)
		}
		winningNumbers := strings.Split(gameNumberSplit[0], " ")
		myNumbers := strings.Split(gameNumberSplit[1], " ")
		numCount, err := countWinningNumbers(&myNumbers, &winningNumbers)
		if err != nil {
			panic(err)
		}
		score += int(math.Pow(2, float64(numCount-1)))

	}
	fmt.Println("Part 1 Answer: ", score)
}

func secondPart() {
	inputFileName := "input.txt"
	content, err := os.ReadFile(inputFileName)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(content), "\n")
	count := 0
	linesLen := len(lines)
	cache := make(map[int]int)
	for {
		if count >= len(lines) {
			break
		}
		line := lines[count]
		count++
		cardSplit := strings.Split(line, ": ")
		if len(cardSplit) != 2 {
			panic("malformed game line " + line)
		}
		cardNumberSplit := strings.Split(cardSplit[0], " ")
		cardNumber, err := strconv.Atoi(cardNumberSplit[len(cardNumberSplit)-1])
		if err != nil {
			panic("could not convert card number to int " + cardNumberSplit[1])
		}
		gameNumbers := cardSplit[1]
		gameNumberSplit := strings.Split(gameNumbers, " | ")
		if len(gameNumberSplit) != 2 {
			panic("malformed game numbers " + gameNumbers)
		}
		winningNumbers := strings.Split(gameNumberSplit[0], " ")
		myNumbers := strings.Split(gameNumberSplit[1], " ")
		var numCount int
		numCount, found := cache[cardNumber]
		if !found {
			numCount, err = countWinningNumbers(&myNumbers, &winningNumbers)
			if err != nil {
				panic("couldn't count winning numbers")
			}
			cache[cardNumber] = numCount
		}
		for i := 0; i < numCount; i++ {
			idx := cardNumber + i
			if idx > linesLen {
				fmt.Println("adding non existent card?")
				break
			}
			lines = append(lines, lines[cardNumber+i])
		}
	}
	fmt.Println("Part 2 Answer: ", len(lines))
}

func main() {
	firstPart()
	secondPart()
}
