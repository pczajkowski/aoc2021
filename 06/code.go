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

	fishStrings := strings.Split(string(content), ",")
	var input []int
	for _, fishString := range fishStrings {
		if fishString == "" {
			continue
		}

		if number, err := strconv.Atoi(fishString); err == nil {
			input = append(input, number)
		} else {
			log.Fatal(err)
		}
	}

	return input
}

func part1(fish []int) int {
	for d := 0; d < 80; d++ {
		max := len(fish)
		for i := 0; i < max; i++ {
			fish[i]--
			if fish[i] < 0 {
				fish[i] = 6
				fish = append(fish, 8)
			}
		}
	}

	return len(fish)
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide a file name as argument")
	}

	input := readInput(os.Args[1])
	fmt.Println("Part 1:", part1(input))
}
