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

func costPart1(steps int, value int) int {
	return steps * value
}

func costPart2(steps int, value int) int {
	cost := 0

	for i := 0; i < steps; i++ {
		cost += (i + 1) * value
	}

	return cost
}

func calculate(crabs map[int]int, orderedCrabs []int, cost func(int, int) int) int {
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

			route += cost(steps, value)
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
	fmt.Println("Part1: ", calculate(crabs, orderedCrabs, costPart1))
	fmt.Println("Part2: ", calculate(crabs, orderedCrabs, costPart2))
}
