package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

type ValueRange struct {
	min, max int
}

func (top ValueRange) InRange(n int) bool {
	return n >= top.min && n <= top.max
}
func (top ValueRange) Intersects(bottom ValueRange) bool {
	return !(top.max < bottom.min || top.min > bottom.max)
}
func (vr ValueRange) Union(vr2 ValueRange) ValueRange {
	union := ValueRange{min: min(vr.min, vr2.min), max: max(vr.max, vr2.max)}
	return union
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

	startTime = time.Now()
	result = part2(input)
	fmt.Printf("Part 2: %d\n", result)
	executionTime = time.Since(startTime).Microseconds()
	fmt.Printf("Completed Part 2 in %d µs\n\n", executionTime)
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

func part2(input []string) int {

	allRanges := make([]ValueRange, 0)
	for _, line := range input {
		if len(line) == 0 {
			break
		}
		str := strings.Split(line, "-")
		vr := ValueRange{}
		vr.min, _ = strconv.Atoi(str[0])
		vr.max, _ = strconv.Atoi(str[1])
		allRanges = append(allRanges, vr)
	}
	prevCount := 0
	newCount := len(allRanges)

	for prevCount != newCount {
		slices.SortFunc(allRanges, func(a, b ValueRange) int {
			return a.min - b.min
		})
		index := len(allRanges) - 1
		for index > 0 {
			r1 := allRanges[index]
			r2 := allRanges[index-1]
			if r1.Intersects(r2) {
				allRanges[index-1] = r1.Union(r2)
				allRanges = slices.Delete(allRanges, index, index+1)
			}
			index--
		}
		prevCount = newCount
		newCount = len(allRanges)
	}
	total := 0
	for _, rng := range allRanges {
		total += rng.max - rng.min + 1
	}
	return total
}
