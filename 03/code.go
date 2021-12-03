package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
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

func part1(lines []string) int64 {
	var gamma []string
	var epsilon []string

	for i := 0; i < len(lines[0]); i++ {
		counts := []int{0, 0}
		for _, line := range lines {
			if line[i] == '0' {
				counts[0]++
			} else {
				counts[1]++
			}
		}

		if counts[0] > counts[1] {
			gamma = append(gamma, "0")
			epsilon = append(epsilon, "1")
		} else {
			gamma = append(gamma, "1")
			epsilon = append(epsilon, "0")
		}
	}

	gammaNumber, err := strconv.ParseInt(strings.Join(gamma, ""), 2, 64)
	if err != nil {
		log.Fatal(err)
	}

	epsilonNumber, err := strconv.ParseInt(strings.Join(epsilon, ""), 2, 64)
	if err != nil {
		log.Fatal(err)
	}

	return gammaNumber * epsilonNumber
}

func keepWith(slice []string, index int, char byte) []string {
	var newSlice []string
	for _, line := range slice {
		if line[index] == char {
			newSlice = append(newSlice, line)
		}
	}

	return newSlice
}

func most(lines []string, values []byte, counts []int, index int) []string {
	if counts[0] > counts[1] {
		return keepWith(lines, index, values[0])
	}

	return keepWith(lines, index, values[1])
}

func least(lines []string, values []byte, counts []int, index int) []string {
	if counts[0] < counts[1] {
		return keepWith(lines, index, values[0])
	}

	return keepWith(lines, index, values[1])
}

func filter(lines []string, values []byte, decision func([]string, []byte, []int, int) []string) string {
	for i := 0; i < len(lines[0]); i++ {
		if len(lines) == 1 {
			break
		}

		counts := []int{0, 0}
		for _, line := range lines {
			if line[i] == values[0] {
				counts[0]++
			} else {
				counts[1]++
			}
		}

		lines = decision(lines, values, counts, i)
	}

	if len(lines) == 1 {
		return lines[0]
	}

	return ""
}

func part2(lines []string) int64 {
	generator := make([]string, len(lines))
	copy(generator, lines)
	scrubber := make([]string, len(lines))
	copy(scrubber, lines)

	generatorNumber, err := strconv.ParseInt(filter(generator, []byte{'0', '1'}, most), 2, 64)
	if err != nil {
		log.Fatal(err)
	}

	scrubberNumber, err := strconv.ParseInt(filter(scrubber, []byte{'1', '0'}, least), 2, 64)
	if err != nil {
		log.Fatal(err)
	}

	return generatorNumber * scrubberNumber
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("You need to provide path to input file!")
	}

	lines := readInput(os.Args[1])
	fmt.Println("Part1: ", part1(lines))
	fmt.Println("Part2: ", part2(lines))
}
