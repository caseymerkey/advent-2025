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

type PressesAndJoltages struct {
	ButtonPresses int
	Joltages      []int
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

	startTime = time.Now()
	result = part2(input)
	fmt.Printf("Part 2: %d\n", result)
	executionTime = time.Since(startTime).Milliseconds()
	fmt.Printf("Completed Part 2 in %d ms\n\n", executionTime)
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

// Borrowed heavily from https://github.com/JoanaBLate/advent-of-code-js/blob/main/2025/day10-solve2.js, which, in
// turn leveraged https://www.reddit.com/r/adventofcode/comments/1pk87hl/2025_day_10_part_2_bifurcate_your_way_to_victory/
func part2(input []string) int {
	var result float64
	for _, line := range input {
		joltages := readJoltages(line)
		buttons := readButtons(line)
		cache := make(map[string]float64)

		combosByPattern := combosByPattern(buttons, len(joltages))

		var countPresses func(target []int, level int) float64
		countPresses = func(target []int, level int) float64 {
			var sb strings.Builder
			for i, n := range target {
				if i > 0 {
					sb.WriteString(",")
				}
				sb.WriteString(strconv.Itoa(n))
			}
			key := sb.String()
			val, found := cache[key]
			if found {
				return float64(val)
			}

			onlyZeroes := true
			for _, jolt := range target {
				switch {
				case jolt < 0:
					return math.Inf(1)
				case jolt > 0:
					onlyZeroes = false
				}
			}
			if onlyZeroes {
				return 0
			}

			parityPattern := parityPattern(target)

			total := math.Inf(1)

			combos, found := combosByPattern[parityPattern]
			if !found {
				cache[key] = total
				return total
			}

			for _, combo := range combos {
				halfTarget := make([]int, len(target))
				for i, j := range combo.Joltages {
					halfTarget[i] = (target[i] - j) / 2
				}

				presses := float64(combo.ButtonPresses) + 2*countPresses(halfTarget, level+1)

				if presses < total {
					total = presses
				}
			}
			cache[key] = total
			return total
		}
		result += countPresses(joltages, 1)
	}
	return int(result)
}

func readJoltages(line string) []int {
	joltages := make([]int, 0)
	segments := strings.Split(line, " ")
	segment := segments[len(segments)-1]
	segment = segment[1 : len(segment)-1]
	for nStr := range strings.SplitSeq(segment, ",") {
		n, _ := strconv.Atoi(nStr)
		joltages = append(joltages, n)
	}
	return joltages
}

func readButtons(line string) [][]int {

	segments := strings.Split(line, " ")
	segments = segments[1 : len(segments)-1]

	buttons := make([][]int, 0)
	for _, segment := range segments {
		button := make([]int, 0)
		numStrs := segment[1 : len(segment)-1]
		for nStr := range strings.SplitSeq(numStrs, ",") {
			n, _ := strconv.Atoi(nStr)
			button = append(button, n)
		}
		buttons = append(buttons, button)
	}
	return buttons
}

func combosByPattern(buttons [][]int, bitcount int) map[string][]PressesAndJoltages {
	//combos is a map of the even-odd parity pattern to
	// a secondary map of button presses that generate that pattern,
	// and their corresponding joltages produced
	combos := make(map[string][]PressesAndJoltages)

	maxComboCount := int(math.Pow(2, float64(len(buttons))))
	for i := range maxComboCount {
		joltages := make([]int, bitcount)
		buttonsPressed := make([][]int, 0)
		for k := range len(buttons) {
			n := int(math.Pow(2, float64(k)))
			if (i & n) == n {
				buttonsPressed = append(buttonsPressed, buttons[k])
			}
		}

		for _, b := range buttonsPressed {
			joltages = pressButton(b, joltages)
		}
		joltageParityPattern := parityPattern(joltages)
		psAndJs, found := combos[joltageParityPattern]
		if !found {
			psAndJs = make([]PressesAndJoltages, 0)
		}
		psAndJs = append(psAndJs, PressesAndJoltages{ButtonPresses: len(buttonsPressed), Joltages: joltages})
		combos[joltageParityPattern] = psAndJs
	}

	return combos
}

func pressButton(button []int, joltages []int) []int {
	for _, n := range button {
		joltages[n] = joltages[n] + 1
	}
	return joltages
}

func parityPattern(joltages []int) string {
	var b strings.Builder
	for _, val := range joltages {
		b.WriteString(strconv.Itoa(val % 2))
	}
	return b.String()
}
