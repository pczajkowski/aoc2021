package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func readInput(file string) map[int]int {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	fishStrings := strings.Split(string(content), ",")
	input := make(map[int]int)
	for _, fishString := range fishStrings {
		if fishString == "" {
			continue
		}

		if number, err := strconv.Atoi(fishString); err == nil {
			input[number]++
		} else {
			log.Fatal(err)
		}
	}

	return input
}

func part1(fish map[int]int, days int) int {
	for d := 0; d < days; d++ {
		newFish := make(map[int]int)
		for k, v := range fish {
			if k == 0 {
				newFish[8] = v
				newFish[6] += v
				continue
			}

			newFish[k-1] += v
		}

		fish = newFish
	}

	sum := 0
	for _, v := range fish {
		sum += v
	}

	return sum
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide a file name as argument")
	}

	input := readInput(os.Args[1])
	fmt.Println("Part 1:", part1(input, 80))
	fmt.Println("Part 2:", part1(input, 256))
}
