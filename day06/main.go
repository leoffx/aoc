package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// S = V0*t
func calculateDistance(speed, time int) int {
	return speed * time
}

func parseLineValues(line *string) ([]int, error) {
	values := strings.Split(*line, ":")
	if len(values) != 2 {
		return nil, errors.New("malformed values string")
	}
	var nums []int
	for _, v := range strings.Split(values[1], " ") {
		if v == "" {
			continue
		}
		num, err := strconv.Atoi(strings.TrimSpace(v))
		if err != nil {
			return nil, fmt.Errorf("could not convert value to int %s", v)
		}
		nums = append(nums, num)
	}
	return nums, nil
}
func parsePartTwoLineValues(line *string) (int, error) {
	values := strings.Split(*line, ":")
	if len(values) != 2 {
		return 0, errors.New("malformed values string")
	}
	v, err := strconv.Atoi(strings.ReplaceAll(strings.TrimSpace(values[1]), " ", ""))
	if err != nil {
		return 0, fmt.Errorf("could not convert value to int %s", values[1])
	}
	return v, nil
}

func firstPart() {
	inputFileName := "input.txt"
	content, err := os.ReadFile(inputFileName)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(content), "\n")
	if len(lines) != 2 {
		panic("malformed input text")
	}
	times, err := parseLineValues(&lines[0])
	if err != nil {
		panic(err)
	}
	distances, err := parseLineValues(&lines[1])
	if err != nil {
		panic(err)
	}
	if len(times) != len(distances) {
		panic("times and distances length don't match")
	}
	ans := 1
	// iterate through each race
	for i := 0; i < len(times); i++ {
		time := times[i]
		recordDistance := distances[i]
		speed := 0
		score := 0
		// iterate through all possible times on this race
		for j := 1; j < time; j++ {
			speed += 1
			remainingTime := time - j
			if calculateDistance(speed, remainingTime) > recordDistance {
				score += 1
			}
		}
		if score != 0 {
			ans = ans * score
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
	if len(lines) != 2 {
		panic("malformed input text")
	}
	time, err := parsePartTwoLineValues(&lines[0])
	if err != nil {
		panic(err)
	}
	recordDistance, err := parsePartTwoLineValues(&lines[1])
	if err != nil {
		panic(err)
	}
	ans := 1
	speed := 0
	score := 0
	// could be done more efficiently with two binary searches
	for i := 1; i < time; i++ {
		speed += 1
		remainingTime := time - i
		if calculateDistance(speed, remainingTime) > recordDistance {
			score += 1
		}
	}
	if score != 0 {
		ans = ans * score
	}

	fmt.Println("Part 2 Answer: ", ans)
}

func main() {
	firstPart()
	secondPart()
}
