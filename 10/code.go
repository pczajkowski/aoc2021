package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func readInput(file string) []string {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")
	return lines
}

var points map[rune]int

func init() {
	points = make(map[rune]int)
	points[')'] = 3
	points[']'] = 57
	points['}'] = 1197
	points['>'] = 25137
}

func parseLine(line string) rune {
	opened := make(map[rune]int)
	var lastOpened []rune

	for _, char := range line {
		if char == '(' || char == '[' || char == '{' || char == '<' {
			opened[char]++
			lastOpened = append(lastOpened, char)
			continue
		}

		if len(lastOpened) == 0 {
			return char
		}

		switch char {
		case ')':
			if lastOpened[len(lastOpened)-1] != '(' {
				return char
			}
		case ']':
			if lastOpened[len(lastOpened)-1] != '[' {
				return char
			}
		case '}':
			if lastOpened[len(lastOpened)-1] != '{' {
				return char
			}
		case '>':
			if lastOpened[len(lastOpened)-1] != '<' {
				return char
			}
		}

		lastOpened = lastOpened[:len(lastOpened)-1]
	}

	return ' '
}

func part1(input []string) int {
	var total int
	for _, line := range input {
		illegal := parseLine(line)
		if illegal != ' ' {
			total += points[illegal]
		}
	}

	return total
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide a file name as argument")
	}

	input := readInput(os.Args[1])
	fmt.Println("Part 1:", part1(input))
}
