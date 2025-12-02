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

type idRange struct {
	min, max string
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

	ranges := make([]idRange, 0)
	for scanner.Scan() {
		line := scanner.Text()
		for _, rString := range strings.Split(line, ",") {
			if len(rString) > 0 {
				minMax := strings.Split(rString, "-")
				ranges = append(ranges, idRange{min: minMax[0], max: minMax[1]})
			}
		}
	}

	var startTime = time.Now()
	result := part1(ranges)
	fmt.Printf("Part 1: %d\n", result)
	executionTime := float32(time.Since(startTime).Milliseconds()) / float32(1000)
	fmt.Printf("Completed Part 1 in %f seconds\n\n", executionTime)

	startTime = time.Now()
	result = part2(ranges)
	fmt.Printf("Part 2: %d\n", result)
	executionTime = float32(time.Since(startTime).Milliseconds()) / float32(1000)
	fmt.Printf("Completed Part 2 in %f seconds\n", executionTime)
}

func part1(ranges []idRange) int {
	total := 0
	for _, rng := range ranges {
		results := invalidIdsInRangePart(rng, 2)
		// fmt.Printf("%s-%s:  %v\n", rng.min, rng.max, results)
		for _, id := range results {
			total += id
		}

	}
	return total
}

func part2(ranges []idRange) int {
	total := 0
	for _, rng := range ranges {
		tmpIds := make(map[int]bool)
		for k := 2; k <= len(rng.max); k++ {
			results := invalidIdsInRangePart(rng, k)
			for _, id := range results {
				if !tmpIds[id] {
					total += id
					tmpIds[id] = true
				}
			}
		}
		// fmt.Printf("%s-%s: %v\n", rng.min, rng.max, tmpIds)
	}
	return total
}

func invalidIdsInRangePart(rng idRange, chunks int) []int {

	ids := make([]int, 0)
	n := len(rng.min)
	x := len(rng.max)
	d := x - n

	// if both min & max are the same number of digits, but not divisible by # of chunks
	// there are no invalid ids in range
	if d == 0 && (n%chunks != 0) {
		return []int{}
	}

	// after the above test condition, one or both of min & max are an appropriate
	// number of digits
	min := rng.min
	max := rng.max
	// if the min is not evenly divisible by chunks, bump it up to the lowest n+1 digit value
	if n%chunks != 0 {
		min = fmt.Sprintf("1%s", strings.Repeat("0", n))
		n++
	}
	// if the max is not evenly divisible by chunks, knock it down the highest x-1 digit value
	if x%chunks != 0 {
		max = strings.Repeat("9", x-1)
		x--
	}

	maxVal, _ := strconv.Atoi(max)
	minVal, _ := strconv.Atoi(min)
	testVal, _ := strconv.Atoi(min[:n/chunks])

	// Split the id in chunks and begin incrementing that number
	for {
		valStr := strconv.Itoa(testVal)
		repeatVal, _ := strconv.Atoi(strings.Repeat(valStr, chunks))

		//. if the number created by repeating the chunk number is in range, add it to the list
		if repeatVal <= maxVal && repeatVal >= minVal {
			ids = append(ids, repeatVal)
		}
		// if we're out of range, then exit the loop
		if repeatVal > maxVal {
			break
		}
		testVal++
	}

	return ids
}
