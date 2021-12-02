package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type move struct {
	direction string
	steps     int
}

func readInput(file string) []move {
	var moves []move

	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}

		var currentMove move
		n, err := fmt.Sscanf(line, "%s %d", &currentMove.direction, &currentMove.steps)
		if err != nil {
			log.Fatal(err)
		}

		if n != 2 {
			continue
		}

		moves = append(moves, currentMove)
	}

	return moves
}

var actions map[string]int

func init() {
	actions = make(map[string]int)
	actions["forward"] = 1
	actions["down"] = 1
	actions["up"] = -1
}

func part1(moves []move) int {
	position := 0
	depth := 0

	for _, move := range moves {
		if move.direction == "forward" {
			position += move.steps * actions[move.direction]
		} else {
			depth += move.steps * actions[move.direction]
		}
	}

	return position * depth
}

func part2(input []move) int {
	position := 0
	depth := 0
	aim := 0

	for _, move := range input {
		switch move.direction {
		case "forward":
			position += move.steps
			depth += aim * move.steps
		case "down":
			aim += move.steps
		case "up":
			aim -= move.steps
		}
	}

	return position * depth
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("No input file specified")
	}

	moves := readInput(os.Args[1])
	fmt.Println("Part 1:", part1(moves))
	fmt.Println("Part 2:", part2(moves))
}
