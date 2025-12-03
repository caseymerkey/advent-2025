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
	executionTime := time.Since(startTime).Milliseconds()
	fmt.Printf("Completed Part 1 in %d ms\n\n", executionTime)
}

func part1(banks [][]int) any {
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
