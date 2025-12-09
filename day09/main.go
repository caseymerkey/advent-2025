package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"image"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type QueueItem struct {
	Value    [2]image.Point
	Priority int
	index    int
}

type PriorityQueue []*QueueItem

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// For a min-heap, use pq[i].priority < pq[j].priority
	// For a max-heap, use pq[i].priority > pq[j].priority
	return pq[i].Priority > pq[j].Priority
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

var Up = image.Point{X: 0, Y: 1}
var Right = image.Point{X: 1, Y: 0}
var Down = image.Point{X: 0, Y: -1}
var Left = image.Point{X: -1, Y: 0}

var Directions = []image.Point{Up, Right, Down, Left}

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
	allPoints := make([]image.Point, 0)
	for _, line := range input {
		valStrs := strings.Split(line, ",")
		x, _ := strconv.Atoi(valStrs[0])
		y, _ := strconv.Atoi(valStrs[1])
		point := image.Point{X: x, Y: y}
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
				item := &QueueItem{Value: [2]image.Point{a, b}, Priority: rectangleArea(a, b)}
				heap.Push(&pq, item)
			}
		}
	}

	item := heap.Pop(&pq).(*QueueItem)

	return item.Priority
}

func part2(input []string) int {
	allPoints := make([]image.Point, 0)
	allPointsBoundary := make(map[image.Point]bool)

	for k, line := range input {
		valStrs := strings.Split(line, ",")
		x, _ := strconv.Atoi(valStrs[0])
		y, _ := strconv.Atoi(valStrs[1])
		point := image.Point{X: x, Y: y}
		allPoints = append(allPoints, point)

		if k > 0 {
			linePoints := allPointsBetween(allPoints[k-1], point)
			for _, p := range linePoints {
				allPointsBoundary[p] = true
			}
		}
	}
	linePoints := allPointsBetween(allPoints[len(allPoints)-1], allPoints[0])
	for _, p := range linePoints {
		allPointsBoundary[p] = true
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
				item := &QueueItem{Value: [2]image.Point{a, b}, Priority: rectangleArea(a, b)}
				heap.Push(&pq, item)
			}
		}
	}

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*QueueItem)
		if isInBounds(item.Value[0], item.Value[1], allPointsBoundary) {
			return item.Priority
		}
	}

	return 0
}

func rectFromCorners(a, b image.Point) image.Rectangle {
	minX := min(a.X, b.X)
	maxX := max(a.X, b.X)
	minY := min(a.Y, b.Y)
	maxY := max(a.Y, b.Y)
	return image.Rect(minX, minY, maxX, maxY)
}

func innerByOne(a, b image.Point) image.Rectangle {
	rect := rectFromCorners(a, b)
	return rect.Inset(1)
}

func isInBounds(a, b image.Point, boundaryMap map[image.Point]bool) bool {

	rect := innerByOne(a, b)
	corner1 := image.Point{X: rect.Min.X, Y: rect.Min.Y}
	corner2 := image.Point{X: rect.Max.X, Y: rect.Min.Y}
	corner3 := image.Point{X: rect.Max.X, Y: rect.Max.Y}
	corner4 := image.Point{X: rect.Min.X, Y: rect.Max.Y}

	line := allPointsBetween(corner2, corner1)
	for _, p := range line {
		if boundaryMap[p] {
			return false
		}
	}
	line = allPointsBetween(corner2, corner3)
	for _, p := range line {
		if boundaryMap[p] {
			return false
		}
	}
	line = allPointsBetween(corner4, corner3)
	for _, p := range line {
		if boundaryMap[p] {
			return false
		}
	}
	line = allPointsBetween(corner4, corner1)
	for _, p := range line {
		if boundaryMap[p] {
			return false
		}
	}

	return true
}

func allPointsBetween(a, b image.Point) []image.Point {
	var direction image.Point
	var steps int
	if a.X == b.X {
		if a.Y < b.Y {
			direction = Up
		} else {
			direction = Down
		}
		steps = abs(a.Y - b.Y)
	} else {
		if a.X < b.X {
			direction = Right
		} else {
			direction = Left
		}
		steps = abs(a.X - b.X)
	}
	allPoints := make([]image.Point, 0)
	p := a
	for i := 0; i < steps; i++ {
		newPoint := p.Add(direction)
		allPoints = append(allPoints, newPoint)
		p = newPoint
	}
	return allPoints
}

func rectangleArea(a, b image.Point) int {
	return (abs(a.X-b.X) + 1) * (abs(a.Y-b.Y) + 1)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
