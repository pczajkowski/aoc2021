package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func readInput(file string) []int {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")
	var input []int
	for _, line := range lines {
		if line == "" {
			continue
		}

		if number, err := strconv.Atoi(line); err == nil {
			input = append(input, number)
		} else {
			log.Fatal(err)
		}
	}

	return input
}

func part1(input []int) int {
	increase := 0
	for i := 1; i < len(input); i++ {
		if input[i-1] < input[i] {
			increase++
		}
	}

	return increase
}

func part2(input []int) []int {
	var measurements []int
	index := 0
	for i, value := range input {
		if i+2 < len(input) {
			measurements = append(measurements, value)
			measurements[index] += input[i+1]
			measurements[index] += input[i+2]
			index++
		}
	}

	return measurements
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide a file name as argument")
	}

	input := readInput(os.Args[1])
	fmt.Println("Part 1:", part1(input))
	fmt.Println("Part 2:", part1(part2(input)))
}
