package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func readInput(file string) [][]int {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")
	var input [][]int
	for _, line := range lines {
		if line == "" {
			continue
		}

		var numbers []int
		for _, letter := range line {
			if number, err := strconv.Atoi(string(letter)); err == nil {
				numbers = append(numbers, number)
			} else {
				log.Fatal(err)
			}
		}

		input = append(input, numbers)
	}

	return input
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide a file name as argument")
	}

	input := readInput(os.Args[1])
	fmt.Println(input)
}
