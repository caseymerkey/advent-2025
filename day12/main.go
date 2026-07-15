package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Present struct {
	Pattern []string
}

func (p *Present) Tiles() int {
	tiles := 0
	for _, row := range p.Pattern {
		for _, ch := range row {
			if ch == '#' {
				tiles++
			}
		}
	}
	return tiles
}

type Region struct {
	Width, Height int
	Quantities    []int
}

func (r *Region) String() string {
	return fmt.Sprintf("%dx%d: %v", r.Width, r.Height, r.Quantities)
}
func (r *Region) Area() int {
	return r.Height * r.Width
}
func (r *Region) ChunkedArea(chunkSize int) int {
	return (r.Width / chunkSize) * (r.Height / chunkSize)
}
func (r *Region) TotalPresentCount() int {
	sum := 0
	for _, n := range r.Quantities {
		sum += n
	}
	return sum
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

	var startTime time.Time
	var result int
	var executionTime int64
	startTime = time.Now()
	result = part1(input)
	fmt.Printf("Part 1: %d\n", result)
	executionTime = time.Since(startTime).Milliseconds()
	fmt.Printf("Completed Part 1 in %d ms\n\n", executionTime)

	// startTime = time.Now()
	// result = part2(input)
	// fmt.Printf("Part 2: %d\n", result)
	// executionTime = time.Since(startTime).Milliseconds()
	// fmt.Printf("Completed Part 2 in %d ms\n\n", executionTime)
}

func readInput(input []string) ([]Present, []Region) {

	presents := make([]Present, 0)
	regions := make([]Region, 0)

	indexRE := regexp.MustCompile(`^\d+:$`)
	regionRE := regexp.MustCompile(`^(\d+)x(\d+):([\d ]+)+$`)

	var present Present

	for _, line := range input {

		if indexRE.MatchString(line) {
			// new Present
			present = Present{}

		} else if line == "" {
			presents = append(presents, present)

		} else if matches := regionRE.FindStringSubmatch(line); matches != nil {
			width, _ := strconv.Atoi(matches[1])
			height, _ := strconv.Atoi(matches[2])
			quantityStrings := strings.Split(strings.Trim(matches[3], " "), " ")
			quantities := make([]int, len(quantityStrings))
			for i, s := range quantityStrings {
				q, _ := strconv.Atoi(s)
				quantities[i] = q
			}

			region := Region{Width: width, Height: height, Quantities: quantities}
			regions = append(regions, region)

		} else {
			// building Present
			present.Pattern = append(present.Pattern, line)
		}
	}

	return presents, regions

}

func part1(input []string) int {
	presents, regions := readInput(input)

	regionsThatFit := 0
	for _, r := range regions {
		fmt.Printf("%v...  ", r)

		// 1) See if it just fits every present in its own 3x3 square
		if r.TotalPresentCount() <= r.ChunkedArea(3) {
			regionsThatFit++
			fmt.Println("true!")
			continue
		}

		// 2) Exclude any that don't have space for all present tiles
		totalPresentTiles := 0
		for i, count := range r.Quantities {
			totalPresentTiles += presents[i].Tiles() * count
		}
		if totalPresentTiles > r.Area() {
			fmt.Println("false!")
			continue
		}

		// 3) If we're here, we need to do some more work
		fmt.Println("unknown. More packing required")

	}

	return regionsThatFit
}
