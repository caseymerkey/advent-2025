package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

type Point3D struct {
	X, Y, Z int
}

func (p Point3D) String() string {
	return fmt.Sprintf("%d,%d,%d", p.X, p.Y, p.Z)
}

type Circuit struct {
	Id     int
	Points map[Point3D]bool
}

type QueueItem struct {
	Value    [2]Point3D
	Priority float64
	index    int
}

type PriorityQueue []*QueueItem

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// For a min-heap, use pq[i].priority < pq[j].priority
	// For a max-heap, use pq[i].priority > pq[j].priority
	return pq[i].Priority < pq[j].Priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*QueueItem)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
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
	// executionTime = time.Since(startTime).Microseconds()
	// fmt.Printf("Completed Part 2 in %d µs\n\n", executionTime)
}

func part1(input []string) int {
	allPoints := make([]Point3D, 0)
	for _, line := range input {
		valStrs := strings.Split(line, ",")
		x, _ := strconv.Atoi(valStrs[0])
		y, _ := strconv.Atoi(valStrs[1])
		z, _ := strconv.Atoi(valStrs[2])
		point := Point3D{X: x, Y: y, Z: z}
		allPoints = append(allPoints, point)
	}

	comboMap := make(map[string]bool)
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	for _, a := range allPoints {
		for _, b := range allPoints {
			if a != b {
				var key string
				if a.String() < b.String() {
					key = fmt.Sprintf("%v-%v", a, b)
				} else {
					key = fmt.Sprintf("%v-%v", b, a)
				}
				if comboMap[key] {
					continue
				}
				comboMap[key] = true
				item := &QueueItem{Value: [2]Point3D{a, b}, Priority: distanceBetween(a, b)}
				heap.Push(&pq, item)
			}
		}
	}
	cycles := 1000
	if len(input) < 1000 {
		cycles = 10
	}

	allCircuits := make([]*Circuit, 0)
	pointCircuitMap := make(map[Point3D]*Circuit)

	for k := range cycles {
		item := heap.Pop(&pq).(*QueueItem)
		ptA := item.Value[0]
		ptB := item.Value[1]

		circA, foundA := pointCircuitMap[ptA]
		circB, foundB := pointCircuitMap[ptB]

		if foundA && foundB {
			// both these points are in circuits already. merge the circuits
			// first make sure they're not already in the same circuit
			if circA != circB {
				// copy all of B's points to circA and update the reference map
				for pt, _ := range circB.Points {
					circA.Points[pt] = true
					pointCircuitMap[pt] = circA
				}
				// clear out circB
				circB.Points = map[Point3D]bool{}
			}
		} else if foundA {
			// A is in a circuit already but B is not. Add B to A's circuit
			circA.Points[ptB] = true
			pointCircuitMap[ptB] = circA
		} else if foundB {
			// B is in a circuit already but A is not. Add A to B's circuit
			circB.Points[ptA] = true
			pointCircuitMap[ptA] = circB
		} else {
			// This is a brand new circuit
			newCircuit := Circuit{Id: k, Points: make(map[Point3D]bool)}
			newCircuit.Points[ptA] = true
			newCircuit.Points[ptB] = true
			allCircuits = append(allCircuits, &newCircuit)
			pointCircuitMap[ptA] = &newCircuit
			pointCircuitMap[ptB] = &newCircuit
		}
	}

	slices.SortFunc(allCircuits, func(a, b *Circuit) int {
		return len(b.Points) - len(a.Points)
	})
	product := 1
	for k := range 3 {
		product *= len(allCircuits[k].Points)
	}

	// 5814 is too low
	return product
}

func distanceBetween(a, b Point3D) float64 {
	// sqrt( (a.x-b.x)^2 + (a.y-b.y)^2 + (a.z-b.z)^2 )
	return math.Sqrt(math.Pow(float64(a.X-b.X), 2) + math.Pow(float64(a.Y-b.Y), 2) + math.Pow(float64(a.Z-b.Z), 2))
}
