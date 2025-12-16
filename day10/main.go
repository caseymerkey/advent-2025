package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

// Queue is a type alias for a slice of integers.
type Queue []int

// Enqueue adds an element to the rear of the queue.
func (q *Queue) Enqueue(value int) {
	*q = append(*q, value)
}

// Dequeue removes and returns an element from the front of the queue.
func (q *Queue) Dequeue() (int, error) {
	if q.IsEmpty() {
		return 0, fmt.Errorf("empty queue")
	}
	value := (*q)[0]

	(*q)[0] = 0
	*q = (*q)[1:]
	return value, nil
}

// IsEmpty checks if the queue is empty.
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
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

	// startTime = time.Now()
	// result = part2(input)
	// fmt.Printf("Part 2: %d\n", result)
	// executionTime = time.Since(startTime).Milliseconds()
	// fmt.Printf("Completed Part 2 in %d ms\n\n", executionTime)
}

func part1(input []string) int {
	total := 0
	for _, line := range input {
		segments := strings.Split(line, " ")

		// treat each goal as a binary representation of an int
		// Also, if we start at the end goal, we can work backwards to zero
		// for the same result.
		goalState := readGoalState(segments[0])
		stateBitCount := len(segments[0]) - 2

		// each button can also be an int that can bitwise xor the
		// machine's state
		buttons := make([]int, 0)
		for idx := 1; idx < len(segments)-1; idx++ {
			var button float64
			numStrs := segments[idx][1 : len(segments[idx])-1]
			for nStr := range strings.SplitSeq(numStrs, ",") {
				n, _ := strconv.Atoi(nStr)
				button += math.Pow(2, float64(stateBitCount-n-1))
			}
			buttons = append(buttons, int(button))
		}

		state := 0
		visited := make(map[int]bool)
		queue := &Queue{state}
		counter := 0
		found := false
		tempQueue := &Queue{}

		for {
			for !queue.IsEmpty() {
				state, _ = queue.Dequeue()
				visited[state] = true

				for _, btn := range buttons {
					newState := state ^ btn

					if visited[newState] {
						continue
					}
					if newState == goalState {
						found = true
						break
					} else {
						tempQueue.Enqueue(newState)
					}
				}
				if found {
					break
				}
			}
			if found {
				total += counter + 1
				break
			} else {
				queue = tempQueue
				tempQueue = &Queue{}
				counter++
			}
		}
	}
	return total
}

func readGoalState(goalString string) int {
	goalString = goalString[1 : len(goalString)-1]
	var state float64
	bitCount := len(goalString)
	for i, ch := range goalString {
		if ch == '#' {
			state += math.Pow(2, float64(bitCount-i-1))
		}
	}
	return int(state)
}

// func button(buttonString string) int {

// }
