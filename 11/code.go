package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type octo struct {
	value   int
	flashed bool
}

func readInput(file string) [][]octo {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")
	var input [][]octo
	for _, line := range lines {
		if line == "" {
			continue
		}

		var row []octo
		for _, char := range line {
			if number, err := strconv.Atoi(string(char)); err == nil {
				row = append(row, octo{number, false})
			} else {
				log.Fatal(err)
			}
		}

		input = append(input, row)
	}

	return input
}

func doNeighbors(input [][]octo, x, y int) int {
	var neighbors int
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}

			if x+i < 0 || x+i >= len(input[y]) {
				continue
			}

			if y+j < 0 || y+j >= len(input) {
				continue
			}

			if !input[y+j][x+i].flashed {
				input[y+j][x+i].value++

				if input[y+j][x+i].value > 9 {
					input[y+j][x+i].value = 0
					input[y+j][x+i].flashed = true
					neighbors++
					neighbors += doNeighbors(input, x+i, y+j)
				}
			}
		}
	}

	return neighbors
}

func reset(input [][]octo) {
	for y, row := range input {
		for x, _ := range row {
			if input[y][x].flashed {
				input[y][x].value = 0
				input[y][x].flashed = false
			}
		}
	}
}

func part1(input [][]octo) int {
	var flashed int
	for i := 0; i < 100; i++ {
		for y, row := range input {
			for x, _ := range row {
				if input[y][x].flashed {
					continue
				}

				input[y][x].value++
				if input[y][x].value > 9 {
					input[y][x].value = 0
					input[y][x].flashed = true
					flashed++
					flashed += doNeighbors(input, x, y)
				}
			}
		}

		reset(input)
	}

	return flashed
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide a file name as argument")
	}

	input := readInput(os.Args[1])
	fmt.Println(part1(input))
}
