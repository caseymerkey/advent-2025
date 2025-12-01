package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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
	rotations := make([]string, 0)
	for scanner.Scan() {
		rotations = append(rotations, scanner.Text())
	}

	var startTime = time.Now()
	result := part1(rotations)
	fmt.Printf("Part 1: %d\n", result)
	executionTime := float32(time.Since(startTime).Milliseconds()) / float32(1000)
	fmt.Printf("Completed Part 1 in %f seconds\n\n", executionTime)

	startTime = time.Now()
	result = part2(rotations)
	fmt.Printf("Part 1: %d\n", result)
	executionTime = float32(time.Since(startTime).Milliseconds()) / float32(1000)
	fmt.Printf("Completed Part 2 in %f seconds\n\n", executionTime)
}

const DIAL_SIZE = 100

func part1(rotations []string) int {
	pointingAt := 50
	zeroCount := 0
	for _, r := range rotations {
		pointer, _ := spin(pointingAt, r)
		if pointer == 0 {
			zeroCount++
		}
		pointingAt = pointer
	}
	return zeroCount
}

func part2(rotations []string) int {
	zeroCount := 0
	pointingAt := 50
	for _, r := range rotations {
		newVal, pastZero := spin(pointingAt, r)
		pointingAt = newVal
		zeroCount += pastZero
	}
	return zeroCount
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func spin(pointingAt int, r string) (int, int) {

	direction := r[0]
	count, _ := strconv.Atoi(r[1:])
	if direction == 'L' {
		count = count * -1
	}

	delta := count % DIAL_SIZE
	newVal := pointingAt + delta
	if newVal < 0 {
		newVal = DIAL_SIZE + newVal
	}
	if newVal > DIAL_SIZE-1 {
		newVal = newVal - DIAL_SIZE
	}

	fullRotations := abs(count / DIAL_SIZE)
	timesPastZero := fullRotations
	if direction == 'R' && newVal < pointingAt {
		timesPastZero++
	}
	if direction == 'L' && pointingAt != 0 && (newVal > pointingAt || newVal == 0) {
		timesPastZero++
	}

	return newVal, timesPastZero
}
