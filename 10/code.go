package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
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
	var lastOpened []rune

	for _, char := range line {
		if char == '(' || char == '[' || char == '{' || char == '<' {
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

func part1(input []string) (int, []string) {
	var total int
	var incomplete []string
	for _, line := range input {
		illegal := parseLine(line)
		if illegal != ' ' {
			total += points[illegal]
			continue
		}

		incomplete = append(incomplete, line)
	}

	return total, incomplete
}

func fixLine(line string) []rune {
	var lastOpened []rune

	for _, char := range line {
		if char == '(' || char == '[' || char == '{' || char == '<' {
			lastOpened = append(lastOpened, char)
			continue
		}

		if len(lastOpened) == 0 {
			return lastOpened
		}

		switch char {
		case ')':
			if lastOpened[len(lastOpened)-1] == '(' {
				lastOpened = lastOpened[:len(lastOpened)-1]
			}
		case ']':
			if lastOpened[len(lastOpened)-1] == '[' {
				lastOpened = lastOpened[:len(lastOpened)-1]
			}
		case '}':
			if lastOpened[len(lastOpened)-1] == '{' {
				lastOpened = lastOpened[:len(lastOpened)-1]
			}
		case '>':
			if lastOpened[len(lastOpened)-1] == '<' {
				lastOpened = lastOpened[:len(lastOpened)-1]
			}
		}
	}

	return lastOpened
}

func calculateScore(line []rune) int {
	var score int

	max := len(line)
	for i := max - 1; i >= 0; i-- {
		score *= 5

		switch line[i] {
		case '(':
			score += 1
		case '[':
			score += 2
		case '{':
			score += 3
		case '<':
			score += 4
		}
	}

	return score
}

func part2(input []string) int {
	var scores []int
	for _, line := range input {
		toFix := fixLine(line)
		scores = append(scores, calculateScore(toFix))
	}

	sort.Ints(scores)
	return scores[len(scores)/2]
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide a file name as argument")
	}

	input := readInput(os.Args[1])
	total, incomplete := part1(input)
	fmt.Println("Part 1:", total)
	fmt.Println("Part 2:", part2(incomplete))
}
