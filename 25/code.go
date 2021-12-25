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

func moveEast(input [][]string) bool {
	changed := false
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if input[i][j] == "." || input[i][j] == "v" {
				continue
			}

			if input[i][j] == ">" {
				if j+1 < len(input[i]) {
					if input[i][j+1] == "." {
						input[i][j] = "."
						input[i][j+1] = ">"
						changed = true
						j++
						continue
					}
				} else {
					if input[i][0] == "." {
						input[i][j] = "."
						input[i][0] = ">"
						changed = true
					}
				}
			}
		}
	}

	return changed
}

type point struct {
	x, y int
}

func moveSouth(input [][]string) bool {
	changed := false
	toSkip := make(map[point]bool)
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if toSkip[point{i, j}] {
				continue
			}

			if input[i][j] == "." || input[i][j] == ">" {
				continue
			}

			if input[i][j] == "v" {
				if i+1 < len(input) {
					if input[i+1][j] == "." {
						input[i][j] = "."
						input[i+1][j] = "v"
						toSkip[point{i + 1, j}] = true
						changed = true
						continue
					}
				} else {
					if input[0][j] == "." {
						input[i][j] = "."
						input[0][j] = "v"
						changed = true
					}
				}
			}
		}
	}

	return changed
}

func print(input [][]string) {
	for _, row := range input {
		fmt.Println(row)
	}
	fmt.Println()
}

func part1(input [][]string) int {
	count := 0

	for i := 0; i < 10; i++ {
		fmt.Println(count)
		print(input)
		changedEast := moveEast(input)
		changedSouth := moveSouth(input)
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
