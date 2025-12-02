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
}

func part1(ranges []idRange) int {
	total := 0
	for _, rng := range ranges {
		results := invalidIdsInRange(rng)
		fmt.Printf("%s-%s:  %v\n", rng.min, rng.max, results)
		for _, id := range results {
			total += id
		}

	}
	return total
}

func invalidIdsInRange(rng idRange) []int {

	ids := make([]int, 0)
	// ids must be an even number of characters long
	n := len(rng.min)
	x := len(rng.max)
	d := x - n

	// if both min & max are the same, odd number of digits, there are no
	// invalid ids in range
	if d == 0 && (n%2 != 0) {
		return []int{}
	}

	// after the above test condition, one or both of min & max are even
	// number of digits
	min := rng.min
	max := rng.max
	// if the min is an odd number of digits, bump it up to the lowest n+1 digit value
	if n%2 != 0 {
		min = fmt.Sprintf("1%s", strings.Repeat("0", n))
		n++
	}
	// if the max is an odd number of digits, knock it down the highest x-1 digit value
	if x%2 != 0 {
		max = strings.Repeat("9", x-1)
		x--
	}

	maxVal, _ := strconv.Atoi(max)
	minVal, _ := strconv.Atoi(min)
	testVal, _ := strconv.Atoi(min[:n/2])

	// Split the id in half and begin incrementing that number
	for {
		valStr := strconv.Itoa(testVal)
		repeatVal, _ := strconv.Atoi(fmt.Sprintf("%s%s", valStr, valStr))
		//. if the number created by repeating the half number is in range, add it to the list
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
