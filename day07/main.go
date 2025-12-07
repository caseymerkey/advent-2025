package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

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

	startTime = time.Now()
	result = part2(input)
	fmt.Printf("Part 2: %d\n", result)
	executionTime = time.Since(startTime).Microseconds()
	fmt.Printf("Completed Part 2 in %d µs\n\n", executionTime)
}

func part1(input []string) int {

	s := strings.Index(input[0], "S")
	var beams = map[int]bool{s: true}
	splitCount := 0

	for row := 2; row < len(input); row += 2 {
		splitted := make(map[int]bool)
		for col := range beams {
			switch input[row][col] {
			case '.':
				splitted[col] = true
			case '^':
				splitted[col-1] = true
				splitted[col+1] = true
				splitCount++
			}
		}
		beams = splitted
	}
	return splitCount
}

func part2(input []string) int {

	s := strings.Index(input[0], "S")
	countCache := make(map[string]int)

	var pathsFromHere func(level int, beamLoc int) int
	pathsFromHere = func(level, beamLoc int) int {
		if level == len(input)-1 {
			return 1
		}
		key := fmt.Sprintf("%dx%d", level, beamLoc)
		if count, found := countCache[key]; found {
			return count
		}
		if input[level][beamLoc] == '^' {
			leftCount := pathsFromHere(level+1, beamLoc-1)
			rightCount := pathsFromHere(level+1, beamLoc+1)
			val := leftCount + rightCount
			countCache[key] = val
			return val
		} else {
			val := pathsFromHere(level+1, beamLoc)
			countCache[key] = val
			return val
		}
	}

	return pathsFromHere(1, s)
}
