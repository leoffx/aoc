package main

import (
	"fmt"
	"os"
	"strings"
)

type NodesToDirections = map[string][]string

func getStartingNodes(nodesMap *NodesToDirections) []string {
	var startingNodes []string
	for node := range *nodesMap {
		if strings.HasSuffix(node, "A") {
			startingNodes = append(startingNodes, node)
		}
	}
	return startingNodes
}

func parseNodes(nodes *[]string) (NodesToDirections, error) {
	nodesMap := make(NodesToDirections)
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
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return a / gcd(a, b) * b
}

func lcmOfArray(arr []int) int {
	result := arr[0]
	for i := 1; i < len(arr); i++ {
		result = lcm(result, arr[i])
	}
	return result
}

func secondPart() {
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
	currNodesKeys := getStartingNodes(&nodesMap)
	steps := 0
	canEnd := false
	hitZIdx := make([]int, len(currNodesKeys))
outer:
	for {
		for _, direction := range directions {
			canEnd = true
			for _, i := range hitZIdx {
				if i == 0 {
					canEnd = false
				}
			}
			if canEnd {
				break outer
			}
			canEnd = true
			steps += 1
			for idx, currNodeKey := range currNodesKeys {
				if direction == 'R' {
					currNodesKeys[idx] = nodesMap[currNodeKey][1]
				} else if direction == 'L' {
					currNodesKeys[idx] = nodesMap[currNodeKey][0]
				} else {
					panic(fmt.Sprintf("unreachable direction %q", direction))
				}
				if !strings.HasSuffix(currNodesKeys[idx], "Z") {
					canEnd = false
				} else if hitZIdx[idx] == 0 {
					hitZIdx[idx] = steps
				}
			}
		}
	}
	ans := lcmOfArray(hitZIdx)
	fmt.Println("Part 2 Answer: ", ans)
}

func main() {
	firstPart()
	secondPart()
}
