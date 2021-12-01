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

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide a file name as argument")
	}

	input := readInput(os.Args[1])
	fmt.Println("Part 1:", part1(input))
}
