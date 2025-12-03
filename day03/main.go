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

	banks := make([][]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		bank := make([]int, 0)
		for _, s := range strings.Split(line, "") {
			n, _ := strconv.Atoi(s)
			bank = append(bank, n)
		}
		banks = append(banks, bank)
	}

	var startTime = time.Now()
	result := part1(banks)
	fmt.Printf("Part 1: %d\n", result)
	executionTime := time.Since(startTime).Microseconds()
	fmt.Printf("Completed Part 1 in %d µs\n\n", executionTime)

	startTime = time.Now()
	result = part2(banks)
	fmt.Printf("Part 2: %d\n", result)
	executionTime = time.Since(startTime).Microseconds()
	fmt.Printf("Completed Part 2 in %d µs\n", executionTime)
}

func part1(banks [][]int) int {
	total := 0
	for _, bank := range banks {
		highestTen := 0
		highestTenIx := 0
		highestOne := 0

		for i := range bank {
			if bank[i] > highestTen && i < len(bank)-1 {
				highestTen = bank[i]
				highestTenIx = i
			}
		}
		for i := highestTenIx + 1; i < len(bank); i++ {
			if bank[i] > highestOne {
				highestOne = bank[i]
			}
		}
		total += ((10 * highestTen) + highestOne)
	}
	return total
}

func part2(banks [][]int) int {
	total := 0

	for _, bank := range banks {
		highestDigits := make([]int, 12)
		highestDigitIndexes := make([]int, 12)

		for power := 11; power >= 0; power-- {
			rangeStart := 0
			if power < 11 {
				rangeStart = highestDigitIndexes[power+1] + 1
			}
			reservedIdx := len(bank) - power

			for i := rangeStart; i < reservedIdx; i++ {
				digit := bank[i]
				if digit > highestDigits[power] {
					highestDigits[power] = digit
					highestDigitIndexes[power] = i
				}
			}

		}

		var joltage int
		for power := 11; power >= 0; power-- {
			joltage += highestDigits[power] * int(math.Pow(10, float64(power)))
		}
		total += joltage

	}
	return total
}
