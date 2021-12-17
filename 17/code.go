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
	min, max int
}

func readInput(file string) []point {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	var points []point
	line := string(content)
	initialSplit := strings.Split(line, ", ")
	if len(initialSplit) != 2 {
		log.Fatal("Invalid input")
	}

	firstPartSplit := strings.Split(initialSplit[0], "=")
	if len(firstPartSplit) != 2 {
		log.Fatal("Invalid input")
	}

	firstPartNumbers := strings.Split(firstPartSplit[1], "..")
	if len(firstPartNumbers) != 2 {
		log.Fatal("Invalid input")
	}

	minX, err := strconv.Atoi(firstPartNumbers[0])
	if err != nil {
		log.Fatal(err)
	}

	maxX, err := strconv.Atoi(firstPartNumbers[1])
	if err != nil {
		log.Fatal(err)
	}

	points = append(points, point{minX, maxX})

	secondPartNumbers := strings.Split(initialSplit[1], "..")
	if len(secondPartNumbers) != 2 {
		log.Fatal("Invalid input")
	}

	minY, err := strconv.Atoi(strings.TrimLeft(secondPartNumbers[0], "y="))
	if err != nil {
		log.Fatal(err)
	}

	maxY, err := strconv.Atoi(secondPartNumbers[1])
	if err != nil {
		log.Fatal(err)
	}

	points = append(points, point{minY, maxY})
	return points
}

func findMinXVelocity(xPoint point) int {
	for i := 1; i < xPoint.min; i++ {
		v := i
		x := 0
		for {
			x += v
			if v > 0 {
				v--
			}

			if v == 0 {
				break
			}
		}

		if x >= xPoint.min {
			return i
		}
	}

	return 0
}

func simulateY(v int, yPoint point) int {
	y := 0
	largestY := y
	previousY := y
	for {
		previousY = y
		if y <= yPoint.max {
			break
		}

		y += v
		if y > largestY {
			largestY = y
		}
		v--
	}

	if previousY >= yPoint.min && previousY <= yPoint.max {
		return largestY
	}

	return 0
}

func part2(minXVelocity int, input []point) int {
	found := 0

	for initialXVelocity := minXVelocity; initialXVelocity <= input[0].max; initialXVelocity++ {
		for initialYVelocity := 0 - input[1].min - 1; initialYVelocity >= input[1].min; initialYVelocity-- {
			x := 0
			y := 0
			xV := initialXVelocity
			yV := initialYVelocity
			for {
				x += xV
				y += yV
				if xV > 0 {
					xV--
				}
				yV--

				if x > input[0].max || y < input[1].min {
					break
				}

				if (x >= input[0].min && x <= input[0].max) && (y >= input[1].min && y <= input[1].max) {
					found++
					break
				}
			}
		}
	}

	return found
}

func main() {
	input := readInput(os.Args[1])
	minXVelocity := findMinXVelocity(input[0])
	fmt.Println("Part1:", simulateY(0-input[1].min-1, input[1]))
	fmt.Println("Part2:", part2(minXVelocity, input))
}
