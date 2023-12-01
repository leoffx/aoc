package main

import (
	"fmt"
	"io/ioutil"
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

func main() {
	inputFileName := "input.txt"
	content, err := ioutil.ReadFile(inputFileName)
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
			panic("couldn't convert combined digits to int")
		}
		sum += finalDigit
	}
	fmt.Println("Answer: ", sum)
}
