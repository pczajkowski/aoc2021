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

func main() {
	if len(os.Args) < 2 {
		log.Fatal("No input file specified")
	}

	moves := readInput(os.Args[1])
	fmt.Println(moves)
}
