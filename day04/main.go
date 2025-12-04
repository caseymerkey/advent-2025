package main

import (
	"bufio"
	"fmt"
	"image"
	"log"
	"os"
	"time"
)

var allDirections = []image.Point{
	{0, -1},
	{1, -1},
	{1, 0},
	{1, 1},
	{0, 1},
	{-1, 1},
	{-1, 0},
	{-1, -1},
}

func main() {

	inputFile := "sample.txt"
	if len(os.Args) > 1 && len(os.Args[1]) > 0 {
		inputFile = os.Args[1]
	}

	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	input := make([]string, 0)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	var startTime = time.Now()
	result := part1(input)
	fmt.Printf("Part 1: %d\n", result)
	executionTime := time.Since(startTime).Milliseconds()
	fmt.Printf("Completed Part 1 in %d ms\n\n", executionTime)

	startTime = time.Now()
	result = part2(input)
	fmt.Printf("Part 2: %d\n", result)
	executionTime = time.Since(startTime).Milliseconds()
	fmt.Printf("Completed Part 2 in %d ms\n\n", executionTime)
}

func part1(input []string) int {
	rolls, neighborCounts := initPuzzle(input)

	return removeRolls(rolls, neighborCounts)
}

func part2(input []string) int {

	rolls, neighborCounts := initPuzzle(input)

	removedTotal := 0
	for {
		removed := removeRolls(rolls, neighborCounts)
		removedTotal += removed
		if removed == 0 {
			break
		}
	}
	return removedTotal
}

func initPuzzle(input []string) (map[image.Point]bool, [][]int) {
	mapHeight := len(input)
	mapWidth := len(input[0])

	rolls := make(map[image.Point]bool)
	neighborCounts := make([][]int, mapHeight)
	for row := range mapHeight {
		neighborCounts[row] = make([]int, mapWidth)
	}

	for r, row := range input {
		for c, cell := range row {
			if cell == '@' {
				loc := image.Point{X: c, Y: r}
				rolls[loc] = true
				for _, dir := range allDirections {
					adj := loc.Add(dir)
					if adj.X >= 0 && adj.X < mapWidth && adj.Y >= 0 && adj.Y < mapHeight {
						neighborCounts[adj.Y][adj.X]++
					}
				}
			}
		}
	}
	return rolls, neighborCounts
}

func removeRolls(rollLocations map[image.Point]bool, neighborCounts [][]int) int {

	mapHeight := len(neighborCounts)
	mapWidth := len(neighborCounts[0])

	movedTotal := 0
	toDecrementMap := make(map[image.Point]int)
	for roll := range rollLocations {
		if neighborCounts[roll.Y][roll.X] < 4 {

			for _, dir := range allDirections {
				adj := roll.Add(dir)
				if adj.X >= 0 && adj.X < mapWidth && adj.Y >= 0 && adj.Y < mapHeight {
					toDecrementMap[adj] = toDecrementMap[adj] + 1
				}
			}
			delete(rollLocations, roll)
			movedTotal++
		}
	}
	for adj, val := range toDecrementMap {
		neighborCounts[adj.Y][adj.X] = neighborCounts[adj.Y][adj.X] - val
	}

	return movedTotal
}
