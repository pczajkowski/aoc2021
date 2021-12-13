package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x, y int
}

func readInput(file string) ([]point, int, int) {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")
	var input []point
	var foldX, foldY int
	readingPoints := true
	for _, line := range lines {
		if line == "" {
			readingPoints = false
			continue
		}
		if readingPoints {
			parts := strings.Split(line, ",")
			if len(parts) != 2 {
				log.Fatal("Invalid input")
			}

			x, err := strconv.Atoi(parts[0])
			if err != nil {
				log.Fatal(err)
			}

			y, err := strconv.Atoi(parts[1])
			if err != nil {
				log.Fatal(err)
			}

			input = append(input, point{x, y})
		} else {
			parts := strings.Split(line, "=")
			if len(parts) != 2 {
				log.Fatal("Invalid input")
			}

			if parts[0] == "fold along x" {
				foldX, err = strconv.Atoi(parts[1])
				if err != nil {
					log.Fatal(err)
				}
			} else if parts[0] == "fold along y" {
				foldY, err = strconv.Atoi(parts[1])
				if err != nil {
					log.Fatal(err)
				}
			} else {
				log.Fatal("Invalid input")
			}
		}
	}

	return input, foldX, foldY
}

func foldByY(input []point, foldY int) {
	for i, p := range input {
		if p.y < foldY {
			continue
		}

		newY := p.y % foldY
		if newY > 0 {
			newY = foldY - newY
		}

		input[i].y = newY
	}
}

func countPoints(input []point) int {
	counted := make(map[point]bool)
	count := 0
	for _, p := range input {
		if counted[p] {
			continue
		}

		counted[p] = true
		count++
	}

	return count
}

func part1(input []point, foldY int) int {
	foldByY(input, foldY)
	return countPoints(input)
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide a file name as argument")
	}

	input, _, foldY := readInput(os.Args[1])
	fmt.Println("Part1:", part1(input, foldY))
}
