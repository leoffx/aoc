package main

import (
	"fmt"
	"os"
	"strings"
)

func parseNodes(nodes *[]string) (map[string][]string, error) {
	nodesMap := make(map[string][]string)
	for _, node := range *nodes {
		if node == "" {
			continue
		}
		nodeSplit := strings.Split(node, " = ")
		if len(nodeSplit) != 2 {
			return nil, fmt.Errorf("could not parse node %s", node)
		}
		elements := strings.Trim(nodeSplit[1], "()")
		elementsSplit := strings.Split(elements, ", ")
		if len(elementsSplit) != 2 {
			return nil, fmt.Errorf("could not parse node elements %s", elements)
		}
		nodesMap[nodeSplit[0]] = elementsSplit
	}
	return nodesMap, nil
}

func firstPart() {
	inputFileName := "input.txt"
	content, err := os.ReadFile(inputFileName)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(content), "\n")
	directions := lines[0]
	nodesStr := lines[1:]
	nodesMap, err := parseNodes(&nodesStr)
	if err != nil {
		panic(err)
	}
	currNodeKey := "AAA"
	steps := 0
outer:
	for {
		for _, direction := range directions {
			if currNodeKey == "ZZZ" {
				break outer
			}
			steps += 1
			if direction == 'R' {
				currNodeKey = nodesMap[currNodeKey][1]
			} else if direction == 'L' {
				currNodeKey = nodesMap[currNodeKey][0]
			} else {
				panic(fmt.Sprintf("unreachable direction %q", direction))
			}
		}
	}
	fmt.Println("Part 1 Answer: ", steps)
}

// func secondPart() {
// 	inputFileName := "input.txt"
// 	content, err := os.ReadFile(inputFileName)
// 	if err != nil {
// 		panic(err)
// 	}
// 	lines := strings.Split(string(content), "\n")
// 	fmt.Println("Part 2 Answer: ", lines)
// }

func main() {
	firstPart()
	// secondPart()
}
