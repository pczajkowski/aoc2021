package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func readInput(file string) [][]string {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")
	var input [][]string
	for _, line := range lines {
		if line == "" {
			continue
		}

		row := strings.Split(line, "")
		input = append(input, row)
	}

	return input
}

func moveEast(input [][]string) (bool, [][]string) {
	changed := false
	newBoard := make([][]string, len(input))
	for i := 0; i < len(input); i++ {
		newBoard[i] = make([]string, len(input[i]))
	}

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if input[i][j] == ">" {
				if j+1 < len(input[i]) {
					if input[i][j+1] == "." {
						newBoard[i][j] = "."
						newBoard[i][j+1] = ">"
						changed = true
						continue
					}
				} else {
					if input[i][0] == "." {
						newBoard[i][j] = "."
						newBoard[i][0] = ">"
						changed = true
					}
				}
			}

			if newBoard[i][j] == "" {
				newBoard[i][j] = input[i][j]
			}
		}
	}

	return changed, newBoard
}

func moveSouth(input [][]string) (bool, [][]string) {
	changed := false
	newBoard := make([][]string, len(input))
	for i := 0; i < len(input); i++ {
		newBoard[i] = make([]string, len(input[i]))
	}

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if input[i][j] == "v" {
				if i+1 < len(input) {
					if input[i+1][j] == "." {
						newBoard[i][j] = "."
						newBoard[i+1][j] = "v"
						changed = true
						continue
					}
				} else {
					if input[0][j] == "." {
						newBoard[i][j] = "."
						newBoard[0][j] = "v"
						changed = true
						continue
					}
				}
			}

			if newBoard[i][j] == "" {
				newBoard[i][j] = input[i][j]
			}
		}
	}

	return changed, newBoard
}

func part1(input [][]string) int {
	count := 1
	var changedEast, changedSouth bool

	for {
		changedEast, input = moveEast(input)
		changedSouth, input = moveSouth(input)
		if !changedEast && !changedSouth {
			break
		}

		count++
	}

	return count
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide a file name as argument")
	}

	input := readInput(os.Args[1])
	fmt.Println("Part1:", part1(input))
}
