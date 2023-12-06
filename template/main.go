package main

import (
	"fmt"
	"os"
	"strings"
)

func firstPart() {
	inputFileName := "input.txt"
	content, err := os.ReadFile(inputFileName)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(content), "\n")
	fmt.Println("Part 1 Answer: ", lines)
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
	secondPart()
}
