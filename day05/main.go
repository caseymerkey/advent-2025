package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type ValueRange struct {
	min, max int
}

func (vr ValueRange) InRange(n int) bool {
	return n >= vr.min && n <= vr.max
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

}

func part1(input []string) int {
	freshCount := 0
	idSection := false
	ranges := make([]ValueRange, 0)
	for _, line := range input {
		if idSection {
			id, _ := strconv.Atoi(line)
			for _, vr := range ranges {
				if vr.InRange(id) {
					freshCount++
					break
				}
			}

		} else if len(line) == 0 {
			idSection = true
		} else {
			var rng ValueRange
			str := strings.Split(line, "-")
			rng.min, _ = strconv.Atoi(str[0])
			rng.max, _ = strconv.Atoi(str[1])
			ranges = append(ranges, rng)
		}
	}
	return freshCount
}
