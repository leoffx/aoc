package main

import (
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
}

func secondPart() {
	inputFileName := "input.txt"
	content, err := os.ReadFile(inputFileName)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(content), "\n")
}

func main() {
	firstPart()
	secondPart()
}
