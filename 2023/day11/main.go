package main

import (
	"fmt"
	"os"
	"strings"
)

type Coord struct{ x, y int }

func expandUniverse(universe [][]string) [][]string {
	emptyCols := make(map[int]bool)
	for i := range universe[0] {
		emptyCols[i] = true
		for _, row := range universe {
			if row[i] != "." {
				emptyCols[i] = false
				break
			}
		}
	}

	var newUniverse [][]string
	for _, row := range universe {
		newRow := make([]string, 0, len(row)*2)
		for i, cell := range row {
			newRow = append(newRow, cell)
			if emptyCols[i] {
				newRow = append(newRow, ".")
			}
		}
		newUniverse = append(newUniverse, newRow)
		if isEmptyLine(row) {
			newUniverse = append(newUniverse, newRow)
		}
	}
	return newUniverse
}

func expandUniverse2(universe [][]string) [][]string {
	emptyCols := make(map[int]bool)
	for i := range universe[0] {
		emptyCols[i] = true
		for _, row := range universe {
			if row[i] != "." {
				emptyCols[i] = false
				break
			}
		}
	}
	expansionRate := 3
	var newUniverse [][]string
	for _, row := range universe {
		newRow := make([]string, 0, len(row)*2)
		for i, cell := range row {
			newRow = append(newRow, cell)
			if emptyCols[i] {
				for i := 0; i < expansionRate; i++ {
					newRow = append(newRow, ".")
				}
			}
		}
		newUniverse = append(newUniverse, newRow)
		if isEmptyLine(row) {
			for i := 0; i < expansionRate; i++ {
				newUniverse = append(newUniverse, newRow)
			}
		}
	}
	return newUniverse
}

func isEmptyLine(row []string) bool {
	for _, cell := range row {
		if cell != "." {
			return false
		}
	}
	return true
}

func getGalaxiesCoords(universe [][]string) map[int]Coord {
	galaxyToCoord := make(map[int]Coord)
	galaxyName := 1
	for i, row := range universe {
		for j, cell := range row {
			if cell == "#" {
				galaxyToCoord[galaxyName] = Coord{
					x: i,
					y: j,
				}
				galaxyName += 1
			}
		}
	}
	return galaxyToCoord
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func calculateDistanceBetweenTwoPoints(a, b Coord) int {
	return abs(a.x-b.x) + abs(a.y-b.y)
}

func firstPart() {
	inputFileName := "input.txt"
	content, err := os.ReadFile(inputFileName)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(content), "\n")
	oldUniverse := make([][]string, 0, len(lines))
	for _, line := range lines {
		oldUniverse = append(oldUniverse, strings.Split(line, ""))
	}
	universe := expandUniverse(oldUniverse)
	galaxyToCoord := getGalaxiesCoords(universe)
	ans := 0
	for galaxy1, coord1 := range galaxyToCoord {
		for galaxy2, coord2 := range galaxyToCoord {
			if galaxy1 == galaxy2 {
				continue
			}
			ans += calculateDistanceBetweenTwoPoints(coord1, coord2)
		}
	}
	ans /= 2
	fmt.Println("Part 1 Answer: ", ans)
}

func secondPart() {
	inputFileName := "input.txt"
	content, err := os.ReadFile(inputFileName)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(content), "\n")
	oldUniverse := make([][]string, 0, len(lines))
	for _, line := range lines {
		oldUniverse = append(oldUniverse, strings.Split(line, ""))
	}
	universe := expandUniverse(oldUniverse)
	galaxyToCoord := getGalaxiesCoords(universe)

	// since the expansion is linear, we can use the old galaxy distances to project the new universe growth
	oldGalaxyToCoord := getGalaxiesCoords(oldUniverse)

	galaxyPairToDistance := make(map[[2]int]int)
	for i := 1; i <= len(galaxyToCoord); i++ {
		for j := i + 1; j <= len(galaxyToCoord); j++ {
			coord1 := galaxyToCoord[i]
			coord2 := galaxyToCoord[j]
			galaxyPairToDistance[[2]int{i, j}] = calculateDistanceBetweenTwoPoints(coord1, coord2)
		}
	}

	oldGalaxyPairToDistance := make(map[[2]int]int)
	for i := 1; i <= len(galaxyToCoord); i++ {
		for j := i + 1; j <= len(galaxyToCoord); j++ {
			coord1 := oldGalaxyToCoord[i]
			coord2 := oldGalaxyToCoord[j]
			oldGalaxyPairToDistance[[2]int{i, j}] = calculateDistanceBetweenTwoPoints(coord1, coord2)
		}
	}

	ans := 0
	expansionRate := 1000000
	for key, distance := range galaxyPairToDistance {
		oldDistance := oldGalaxyPairToDistance[key]
		ans += oldDistance + (distance-oldDistance)*(expansionRate-1)
	}
	fmt.Println("Part 2 Answer: ", ans)
}

func main() {
	firstPart()
	secondPart()
}
