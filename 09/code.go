package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readInput(file string) [][]int {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")
	var input [][]int
	for _, line := range lines {
		if line == "" {
			continue
		}

		var numbers []int
		for _, letter := range line {
			if number, err := strconv.Atoi(string(letter)); err == nil {
				numbers = append(numbers, number)
			} else {
				log.Fatal(err)
			}
		}

		input = append(input, numbers)
	}

	return input
}

func hasSmallerNeighbors(input [][]int, x, y int) bool {
	if x-1 >= 0 && input[x-1][y] <= input[x][y] {
		return true
	}

	if x+1 < len(input) && input[x+1][y] <= input[x][y] {
		return true
	}

	if y-1 >= 0 && input[x][y-1] <= input[x][y] {
		return true
	}

	if y+1 < len(input[0]) && input[x][y+1] <= input[x][y] {
		return true
	}

	return false
}

func part1(input [][]int) (int, [][]int) {
	var sum int
	var lowPoints [][]int
	for x, row := range input {
		for y, value := range row {
			if !hasSmallerNeighbors(input, x, y) {
				sum += value + 1
				lowPoints = append(lowPoints, []int{x, y})
			}
		}
	}

	return sum, lowPoints
}

type point struct {
	x, y int
}

func getNeighbors(input [][]int, x, y int) []point {
	var neighbors []point
	if x-1 >= 0 && input[x-1][y] != 9 {
		neighbors = append(neighbors, point{x - 1, y})
	}

	if x+1 < len(input) && input[x+1][y] != 9 {
		neighbors = append(neighbors, point{x + 1, y})
	}

	if y-1 >= 0 && input[x][y-1] != 9 {
		neighbors = append(neighbors, point{x, y - 1})
	}

	if y+1 < len(input[0]) && input[x][y+1] != 9 {
		neighbors = append(neighbors, point{x, y + 1})
	}

	return neighbors
}

func getBasin(input [][]int, x int, y int) []point {
	var basin []point
	neighbors := getNeighbors(input, x, y)

	checked := make(map[point]bool)
	for {
		if len(neighbors) == 0 {
			break
		}

		var next []point
		for _, neighbor := range neighbors {
			if checked[neighbor] {
				continue
			}

			checked[neighbor] = true
			basin = append(basin, neighbor)
			next = append(next, getNeighbors(input, neighbor.x, neighbor.y)...)
		}

		neighbors = next
	}

	return basin
}

func getBasins(input [][]int, lowPoints [][]int) [][]point {
	var basins [][]point

	for _, lowPoint := range lowPoints {
		basin := getBasin(input, lowPoint[0], lowPoint[1])
		basins = append(basins, basin)
	}

	return basins
}

func part2(input [][]int, lowPoints [][]int) int {
	basins := getBasins(input, lowPoints)

	sort.Slice(basins, func(i, j int) bool {
		return len(basins[i]) > len(basins[j])
	})

	result := 1
	for i := 0; i < 3; i++ {
		result *= len(basins[i])
	}

	return result
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide a file name as argument")
	}

	input := readInput(os.Args[1])
	sum, lowPoints := part1(input)
	fmt.Println("Part 1:", sum)
	fmt.Println("Part 2:", part2(input, lowPoints))
}
