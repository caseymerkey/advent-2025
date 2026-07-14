package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

const DAC = 1
const FFT = 2

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

	var startTime time.Time
	var result int
	var executionTime int64
	startTime = time.Now()
	result = part1(input)
	fmt.Printf("Part 1: %d\n", result)
	executionTime = time.Since(startTime).Milliseconds()
	fmt.Printf("Completed Part 1 in %d ms\n\n", executionTime)

	startTime = time.Now()
	result = part2(input)
	fmt.Printf("Part 2: %d\n", result)
	executionTime = time.Since(startTime).Milliseconds()
	fmt.Printf("Completed Part 2 in %d ms\n\n", executionTime)
}

func part1(input []string) int {
	allNodes := parseInput(input)

	var pathsToEndFrom func(start string) int
	pathsToEndFrom = func(start string) int {
		goodPaths := 0
		for _, next := range allNodes[start] {
			if next == "out" {
				goodPaths += 1
			} else {
				goodPaths += pathsToEndFrom(next)
			}
		}
		return goodPaths
	}

	result := pathsToEndFrom("you")

	return result
}

func part2(input []string) int {
	allNodes := parseInput(input)

	cache := make(map[string][4]int)

	var dfs func(string) [4]int
	dfs = func(start string) [4]int {
		// search cache
		pathsToEnd, found := cache[start]
		if found {
			return pathsToEnd
		}

		mask := 0
		switch start {
		case "dac":
			mask = DAC
		case "fft":
			mask = FFT
		}

		for _, next := range allNodes[start] {
			if next == "out" {
				pathsToEnd[mask] += 1
			} else {
				downstream := dfs(next)
				if mask == 0 {
					for i, v := range downstream {
						pathsToEnd[i] += v
					}
				} else {
					for i, v := range downstream {
						newMask := i | mask
						pathsToEnd[newMask] += v
					}
				}
			}
		}
		cache[start] = pathsToEnd
		return pathsToEnd
	}

	result := dfs("svr")
	return result[3]
}

func parseInput(input []string) map[string][]string {
	allNodes := make(map[string][]string)

	for _, line := range input {
		tokens := strings.Split(line, " ")
		source := tokens[0][:len(tokens[0])-1]
		allNodes[source] = append(allNodes[source], tokens[1:]...)
	}

	return allNodes
}
