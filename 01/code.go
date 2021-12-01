package main

import (
	"fmt"
	"io/ioutil"
	"log"
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

func main() {
	input := readInput("./testinput1")
	fmt.Println(input)
}
