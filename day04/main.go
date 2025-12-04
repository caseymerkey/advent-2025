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
	executionTime := time.Since(startTime).Microseconds()
	fmt.Printf("Completed Part 1 in %d µs\n\n", executionTime)
}

func part1(input []string) int {
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

	movableTotal := 0
	for roll := range rolls {
		if neighborCounts[roll.Y][roll.X] < 4 {
			movableTotal++
		}
	}
	return movableTotal
}
