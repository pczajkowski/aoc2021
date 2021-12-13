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

type fold struct {
	val int
	cat string
}

func readInput(file string) ([]point, []fold) {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")
	var input []point
	var folds []fold
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

			val, err := strconv.Atoi(parts[1])
			if err != nil {
				log.Fatal(err)
			}
			if parts[0] == "fold along x" {
				folds = append(folds, fold{val, "x"})
			} else if parts[0] == "fold along y" {
				folds = append(folds, fold{val, "y"})
			} else {
				log.Fatal("Invalid input")
			}
		}
	}

	return input, folds
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

func foldByX(input []point, foldX int) {
	for i, p := range input {
		if p.x > foldX {
			input[i].x = p.x%foldX - 1
			continue
		}

		input[i].x = foldX - p.x - 1
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

func part1(input []point, folds []fold) int {
	firstFold := folds[0]

	if firstFold.cat == "x" {
		foldByX(input, firstFold.val)
	} else {
		foldByY(input, firstFold.val)
	}

	return countPoints(input)
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide a file name as argument")
	}

	input, folds := readInput(os.Args[1])
	fmt.Println("Part1:", part1(input, folds))
}
