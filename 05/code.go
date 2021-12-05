package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func readInput(file string) [][][]int {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")
	var points [][][]int
	for _, line := range lines {
		if line == "" {
			continue
		}

		pointsString := strings.Split(line, " -> ")
		if len(pointsString) != 2 {
			log.Fatal("Invalid line: ", line)
		}

		start := strings.Split(pointsString[0], ",")
		if len(start) != 2 {
			log.Fatal("Invalid start point: ", start)
		}

		startPoint := make([]int, 2)
		startPoint[0], err = strconv.Atoi(start[0])
		if err != nil {
			log.Fatal(err)
		}
		startPoint[1], err = strconv.Atoi(start[1])
		if err != nil {
			log.Fatal(err)
		}

		end := strings.Split(pointsString[1], ",")
		if len(end) != 2 {
			log.Fatal("Invalid end point: ", end)
		}
		endPoint := make([]int, 2)
		endPoint[0], err = strconv.Atoi(end[0])
		if err != nil {
			log.Fatal(err)
		}
		endPoint[1], err = strconv.Atoi(end[1])
		if err != nil {
			log.Fatal(err)
		}

		points = append(points, [][]int{startPoint, endPoint})
	}

	return points
}

func largest(points [][][]int, index int) int {
	largest := 0
	for _, point := range points {
		if point[0][index] > largest {
			largest = point[0][index]
		} else if point[1][index] > largest {
			largest = point[1][index]
		}
	}

	return largest
}

func makeDiagram(points [][][]int) [][]int {
	largestX := largest(points, 0)
	largestY := largest(points, 1)

	diagram := make([][]int, largestY+1)
	for i, _ := range diagram {
		diagram[i] = make([]int, largestX+1)
	}
	return diagram
}

func fillDiagram(points [][][]int, diagram [][]int) {
	for _, point := range points {
		if point[0][0] != point[1][0] && point[0][1] != point[1][1] {
			continue
		}

		if point[0][0] == point[1][0] {
			start := point[0][1]
			end := point[1][1]
			if start > end {
				start, end = end, start
			}

			for i := start; i <= end; i++ {
				diagram[i][point[0][0]]++
			}
		} else {
			start := point[0][0]
			end := point[1][0]
			if start > end {
				start, end = end, start
			}

			for i := start; i <= end; i++ {
				diagram[point[0][1]][i]++
			}
		}
	}
}

func part1(diagram [][]int) int {
	var count int
	for _, row := range diagram {
		for _, value := range row {
			if value >= 2 {
				count++
			}
		}
	}

	return count
}

func fillDiagramDiagonal(points [][][]int, diagram [][]int) {
	for _, point := range points {
		if point[0][0] == point[1][0] {
			start := point[0][1]
			end := point[1][1]
			if start > end {
				start, end = end, start
			}

			for i := start; i <= end; i++ {
				diagram[i][point[0][0]]++
			}
		} else if point[0][1] == point[1][1] {
			start := point[0][0]
			end := point[1][0]
			if start > end {
				start, end = end, start
			}

			for i := start; i <= end; i++ {
				diagram[point[0][1]][i]++
			}
		} else {
			start := point[0]
			end := point[1]
			if start[0] > end[0] {
				start, end = end, start
			}

			goUp := true
			if start[1] < end[1] {
				goUp = false
			}

			x := start[0]
			y := start[1]
			for {
				if x > end[0] {
					break
				}

				if goUp && y < end[1] {
					break
				} else if !goUp && y > end[1] {
					break
				}

				diagram[y][x]++
				x++

				if goUp {
					y--
				} else {
					y++
				}
			}
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide a file name as argument")
	}

	input := readInput(os.Args[1])
	diagram := makeDiagram(input)
	fillDiagram(input, diagram)
	fmt.Println("Part1: ", part1(diagram))

	diagram = makeDiagram(input)
	fillDiagramDiagonal(input, diagram)
	fmt.Println("Part2: ", part1(diagram))
}
