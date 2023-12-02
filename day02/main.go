package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Color string

const (
	Red   Color = "red"
	Green Color = "green"
	Blue  Color = "blue"
)

var bagCapacity = map[Color]int{
	Red:   12,
	Green: 13,
	Blue:  14,
}

func checkBagAmountIsEnough(sets string) bool {
	setsSplit := strings.Split(sets, ";")
	for _, set := range setsSplit {
		cubesRevealed := strings.Split(set, ", ")
		for _, cube := range cubesRevealed {
			cubesSplit := strings.Split(strings.TrimSpace(cube), " ")
			if len(cubesSplit) != 2 {
				panic("malformed cube string: " + cube)
			}
			cubeAmount := cubesSplit[0]
			cubeAmountValue, err := strconv.Atoi(cubeAmount)
			if err != nil {
				panic("coudln't parse cube amount number")
			}
			cubeColor := Color(cubesSplit[1])
			if bagCapacity[cubeColor] < cubeAmountValue {
				return false
			}
		}
	}
	return true
}

// The power of a set of cubes is equal to the numbers of red, green, and blue cubes multiplied together
func getMinimumCubeSetPower(sets string) int {
	minimumSet := map[Color]int{
		Red:   0,
		Green: 0,
		Blue:  0,
	}
	setsSplit := strings.Split(sets, ";")
	for _, set := range setsSplit {
		cubesRevealed := strings.Split(set, ", ")
		for _, cube := range cubesRevealed {
			cubesSplit := strings.Split(strings.TrimSpace(cube), " ")
			if len(cubesSplit) != 2 {
				panic("malformed cube string: " + cube)
			}
			cubeAmount := cubesSplit[0]
			cubeAmountValue, err := strconv.Atoi(cubeAmount)
			if err != nil {
				panic("coudln't parse cube amount number")
			}
			cubeColor := Color(cubesSplit[1])
			if minimumSet[cubeColor] < cubeAmountValue {
				minimumSet[cubeColor] = cubeAmountValue
			}
		}
	}
	return minimumSet[Red] * minimumSet[Green] * minimumSet[Blue]
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
		gameSetsSplit := strings.Split(line, ":")
		if len(gameSetsSplit) != 2 {
			panic("malformed round string")
		}
		gameSplit := strings.Split(gameSetsSplit[0], " ")
		gameNum, err := strconv.Atoi(gameSplit[len(gameSplit)-1])
		if err != nil {
			panic("couldn't parse game number")
		}
		sets := gameSetsSplit[1]
		if checkBagAmountIsEnough(sets) {
			sum += gameNum
		}
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
		gameSetsSplit := strings.Split(line, ":")
		if len(gameSetsSplit) != 2 {
			panic("malformed round string")
		}
		sets := gameSetsSplit[1]
		sum += getMinimumCubeSetPower(sets)
	}
	fmt.Println("Part 2 Answer: ", sum)
}

func main() {
	firstPart()
	secondPart()
}
