package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readInput(file string) (map[int]int, []int) {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	numberStrings := strings.Split(string(content), ",")
	crabs := make(map[int]int)
	var orderedCrabs []int
	for _, numberString := range numberStrings {
		if numberString == "" {
			continue
		}

		if number, err := strconv.Atoi(numberString); err == nil {
			if _, ok := crabs[number]; !ok {
				orderedCrabs = append(orderedCrabs, number)
			}
			crabs[number]++
		} else {
			log.Fatal(err)
		}
	}

	sort.Ints(orderedCrabs)
	return crabs, orderedCrabs
}

func part1(crabs map[int]int, orderedCrabs []int) int {
	shortest := 10000000

	for _, crab := range orderedCrabs {
		route := 0
		for key, value := range crabs {
			if key == crab {
				continue
			}

			if key < crab {
				route += (crab - key) * value
			} else {
				route += (key - crab) * value
			}
		}

		if route < shortest {
			shortest = route
		}
	}

	return shortest
}

func part2(crabs map[int]int, orderedCrabs []int) int {
	shortest := 1000000000

	max := orderedCrabs[len(orderedCrabs)-1]
	for crab := orderedCrabs[0]; crab <= max; crab++ {
		route := 0
		for key, value := range crabs {
			if key == crab {
				continue
			}

			steps := 0
			if key < crab {
				steps = crab - key
			} else {
				steps = key - crab
			}

			for i := 0; i < steps; i++ {
				route += (i + 1) * value
			}
		}

		if route < shortest {
			shortest = route
		}
	}

	return shortest
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide a file name as argument")
	}

	crabs, orderedCrabs := readInput(os.Args[1])
	fmt.Println("Part1: ", part1(crabs, orderedCrabs))
	fmt.Println("Part2: ", part2(crabs, orderedCrabs))
}
