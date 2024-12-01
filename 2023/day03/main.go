package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var directions = []struct{ x, y int }{{1, 0}, {0, 1}, {-1, 0}, {0, -1}, {1, 1}, {-1, 1}, {-1, -1}, {1, -1}}
var nonSymbols = [11]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "."}
var gear = string('*')

func isSymbol(char string) bool {
	for _, nonSymbol := range nonSymbols {
		if char == nonSymbol {
			return false
		}
	}
	return true
}

func hasAdjacentSymbol(schematic *[][]string, i, j int) bool {
	rows_len, cols_len := len(*schematic)-1, len((*schematic)[0])-1
	for _, tuple := range directions {
		x, y := i+tuple.x, j+tuple.y
		if 0 > x || x >= rows_len || 0 > y || y >= cols_len {
			continue
		}
		char := (*schematic)[x][y]
		if isSymbol(char) {
			return true
		}
	}
	return false
}

func parseFullNumber(schematic *[][]string, i, j int, visited *map[[2]int]bool) (int, error) {
	fullNumber := 0
	col := j
	for {
		coord := [2]int{i, col}
		if _, found := (*visited)[coord]; found {
			col--
			return 0, errors.New("already visited")
		}
		(*visited)[coord] = true
		char := (*schematic)[i][col]
		// backtrack until reaching the first digit of full number
		if _, err := strconv.Atoi(char); err == nil {
			col -= 1
			if col < 0 {
				break
			}
			continue
		}
		break
	}
	// now parse full number
	for {
		col++
		char := (*schematic)[i][col]
		num, err := strconv.Atoi(char)
		if err != nil {
			break
		}
		fullNumber = 10*fullNumber + num
	}
	return fullNumber, nil
}

func getAdjacentNumbers(schematic *[][]string, i, j int) []int {
	var visited = make(map[[2]int]bool)
	adjacentNumbers := make([]int, 0)
	rows_len, cols_len := len(*schematic), len((*schematic)[0])
	for _, tuple := range directions {
		x, y := i+tuple.x, j+tuple.y
		if 0 > x || x >= rows_len || 0 > y || y >= cols_len {
			continue
		}
		char := (*schematic)[x][y]
		if _, err := strconv.Atoi(char); err == nil {
			fullNumber, err := parseFullNumber(schematic, x, y, &visited)
			if err != nil {
				continue
			}
			adjacentNumbers = append(adjacentNumbers, fullNumber)
		}
	}
	return adjacentNumbers
}

func firstPart() {
	inputFileName := "input.txt"
	content, err := os.ReadFile(inputFileName)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(content), "\n")
	schematic := make([][]string, 0)
	for _, line := range lines {
		lineSplit := strings.Split(line, "")
		schematic = append(schematic, lineSplit)
	}
	sum := 0
	for i, line := range schematic {
		currNumber := 0
		isNumberValid := false
		for j, symbol := range line {
			char, err := strconv.Atoi(symbol)
			if err != nil {
				if currNumber != 0 && isNumberValid {
					sum += currNumber
				}
				currNumber = 0
				isNumberValid = false
			} else {
				currNumber = 10*currNumber + char
				if !isNumberValid && hasAdjacentSymbol(&schematic, i, j) {
					isNumberValid = true
				}
			}

		}
		if isNumberValid {
			sum += currNumber
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
	schematic := make([][]string, 0)
	for _, line := range lines {
		lineSplit := strings.Split(line, "")
		schematic = append(schematic, lineSplit)
	}
	sum := 0
	for i, line := range schematic {
		for j, symbol := range line {
			if symbol != gear {
				continue
			}
			adjacentNumbers := getAdjacentNumbers(&schematic, i, j)
			if len(adjacentNumbers) != 2 {
				continue
			}
			sum += adjacentNumbers[0] * adjacentNumbers[1]
		}
	}
	fmt.Println("Part 2 Answer: ", sum)
}

func main() {
	firstPart()
	secondPart()
}
