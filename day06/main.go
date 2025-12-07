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

var addition = func(values []int) int {
	total := 0
	for _, val := range values {
		total += val
	}
	return total
}

var multiplication = func(values []int) int {
	product := 1
	for _, val := range values {
		product *= val
	}
	return product
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
	total := 0

	fields := make([][]string, 0)
	for _, line := range input {
		fields = append(fields, strings.Fields(line))
	}

	for i := range len(fields[0]) {
		values := make([]int, 0)
		for j := range len(fields) - 1 {
			val, _ := strconv.Atoi(fields[j][i])
			values = append(values, val)
		}
		var operation func([]int) int
		if fields[len(fields)-1][i] == "+" {
			operation = addition
		} else {
			operation = multiplication
		}

		total += operation(values)
	}

	return total
}

func part2(input []string) int {
	total := 0

	lastRow := len(input) - 1
	col := 0
	var problem int
	var operation func([]int) int
	for {
		var numBuilder strings.Builder
		strikes := 0

		if len(input[lastRow]) > col {
			switch input[lastRow][col] {
			case '+':
				operation = addition
				problem = 0
			case '*':
				operation = multiplication
				problem = 1
			}
		} else {
			strikes++
		}

		for row := range lastRow {
			if len(input[row]) > col {

				n := input[row][col]
				if n != ' ' {
					numBuilder.WriteByte(n)
				}

				strikes = 0
			} else {
				strikes++
			}
		}
		numStr := numBuilder.String()
		if len(numStr) > 0 {
			num, _ := strconv.Atoi(numStr)
			problem = operation([]int{problem, num})
		} else {
			// time to start a new problem
			total += problem
			problem = 0
		}
		if strikes == len(input) {
			break
		}
		col++
	}

	return total
}
