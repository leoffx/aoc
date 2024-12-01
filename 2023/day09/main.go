package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func stringsToInts(stringSlice []string) ([]int, error) {
	ints := make([]int, len(stringSlice))
	for i, str := range stringSlice {
		int, err := strconv.Atoi(strings.TrimSpace(str))
		if err != nil {
			return nil, err
		}
		ints[i] = int
	}
	return ints, nil
}

// calculateNextSequenceValue calculates the next term for a arithmetic progression of higher order
func calculateNextSequenceValue(sequence *[]int) int {

	diffs := [][]int{*sequence}
	lastSequence := diffs[len(diffs)-1]
	for len(lastSequence) > 1 {
		newDiff := make([]int, len(lastSequence)-1)
		for i := 0; i < len(lastSequence)-1; i++ {
			newDiff[i] = lastSequence[i+1] - lastSequence[i]
		}
		diffs = append(diffs, newDiff)
		lastSequence = diffs[len(diffs)-1]
	}

	nextTerm := 0
	for _, diff := range diffs {
		nextTerm += diff[len(diff)-1]
	}
	return nextTerm
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
		history := strings.Split(line, " ")
		historyValues, err := stringsToInts(history)
		if err != nil {
			panic(err)
		}
		extrapolation := calculateNextSequenceValue(&historyValues)
		if err != nil {
			panic(err)
		}
		ans += extrapolation
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
	ans := 0
	for _, line := range lines {
		history := strings.Split(line, " ")
		historyValues, err := stringsToInts(history)
		if err != nil {
			panic(err)
		}
		// to calculate the previous value, simply calculate the next value of the reverse sequence
		slices.Reverse(historyValues)
		extrapolation := calculateNextSequenceValue(&historyValues)
		if err != nil {
			panic(err)
		}
		ans += extrapolation
	}
	fmt.Println("Part 2 Answer: ", ans)
}

func main() {
	// firstPart()
	secondPart()
}
