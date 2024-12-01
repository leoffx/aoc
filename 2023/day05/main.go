package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var mapOrder = [7]string{
	"seed-to-soil",
	"soil-to-fertilizer",
	"fertilizer-to-water",
	"water-to-light",
	"light-to-temperature",
	"temperature-to-humidity",
	"humidity-to-location",
}

func stringsToInts(slice []string) ([]int, error) {
	response := make([]int, 0)
	for _, char := range slice {
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

func parseMaps(lines *[]string) (map[string][][]int, error) {
	response := make(map[string][][]int)
	currentMap := ""
	for _, line := range *lines {
		if strings.Contains(line, "map") {
			currentMap = strings.Split(line, " ")[0]
			response[currentMap] = make([][]int, 0)
		} else if currentMap != "" && len(line) > 1 {
			maps, err := stringsToInts(strings.Split(line, " "))
			if err != nil {
				return nil, err
			}
			response[currentMap] = append(response[currentMap], maps)
		}
	}
	return response, nil
}

func firstPart() {
	inputFileName := "input.txt"
	content, err := os.ReadFile(inputFileName)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(content), "\n")

	seedsSplit := strings.Split(lines[0], ": ")
	if len(seedsSplit) != 2 {
		panic("malformed seeds line")
	}
	seeds, err := stringsToInts(strings.Split(seedsSplit[1], " "))
	if err != nil {
		panic(err)
	}
	maps, err := parseMaps(&lines)
	if err != nil {
		panic(err)
	}
	ans := -1
	for _, seed := range seeds {
		currentKey := seed
		for _, mapType := range mapOrder {
			mapRanges := maps[mapType]
			for _, mapRange := range mapRanges {
				dstRange := mapRange[0]
				srcRange := mapRange[1]
				lenRange := mapRange[2]
				if currentKey >= srcRange && currentKey < srcRange+lenRange {
					currentKey = dstRange + (currentKey - srcRange)
					break
				}
			}
		}
		if currentKey < ans || ans == -1 {
			ans = currentKey
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

	seedsSplit := strings.Split(lines[0], ": ")
	if len(seedsSplit) != 2 {
		panic("malformed seeds line")
	}
	seeds, err := stringsToInts(strings.Split(seedsSplit[1], " "))
	if err != nil {
		panic(err)
	}
	maps, err := parseMaps(&lines)
	if err != nil {
		panic(err)
	}
	ans := -1
	for i := 0; i < len(seeds)/2; i += 2 {
		seedStart := seeds[i]
		seedLen := seeds[i+1]
		for i := 0; i < seedLen; i++ {
			seed := seedStart + i
			currentKey := seed
			for _, mapType := range mapOrder {
				mapRanges := maps[mapType]
				for _, mapRange := range mapRanges {
					dstRange := mapRange[0]
					srcRange := mapRange[1]
					lenRange := mapRange[2]
					if currentKey >= srcRange && currentKey < srcRange+lenRange {
						currentKey = dstRange + (currentKey - srcRange)
						break
					}
				}
			}
			if currentKey < ans || ans == -1 {
				ans = currentKey
			}
		}
	}
	fmt.Println("Part 2 Answer: ", ans)
}

func main() {
	firstPart()
	secondPart()
}
