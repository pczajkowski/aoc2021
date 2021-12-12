package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func readInput(file string) [][]string {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")
	var input [][]string
	for _, line := range lines {
		if line == "" {
			continue
		}

		points := strings.Split(line, "-")
		if len(points) != 2 {
			log.Fatal("Invalid input")
		}

		input = append(input, []string{points[0], points[1]})
		input = append(input, []string{points[1], points[0]})
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
