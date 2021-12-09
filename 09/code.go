package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
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

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide a file name as argument")
	}

	input := readInput(os.Args[1])
	sum, lowPoints := part1(input)
	fmt.Println("Part 1:", sum)
	fmt.Println(lowPoints)
}
