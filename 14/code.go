package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func readInput(file string) (string, map[string]string) {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")
	var template string
	input := make(map[string]string)
	readingTemplate := true
	for _, line := range lines {
		if line == "" {
			readingTemplate = false
			continue
		}

		if readingTemplate {
			template = line
			continue
		}

		parts := strings.Split(line, " -> ")
		if len(parts) != 2 {
			log.Fatal("Invalid line: ", line)
		}

		input[parts[0]] = parts[1]
	}

	return template, input
}

func process(template string, input map[string]string) string {
	var result []string
	for i := 0; i < len(template)-1; i++ {
		insert := input[template[i:i+2]]
		result = append(result, string(template[i]))
		result = append(result, insert)

		if i == len(template)-2 {
			result = append(result, string(template[i+1]))
		}
	}

	return strings.Join(result, "")
}

func countElements(template string) (int, int) {
	counts := make(map[rune]int)
	for _, c := range template {
		counts[c]++
	}

	smallest := counts['N']
	largest := counts['N']
	for _, c := range counts {
		if c < smallest {
			smallest = c
		}

		if c > largest {
			largest = c
		}
	}

	return smallest, largest
}

func part1(template string, input map[string]string) (string, int) {
	result := template
	for i := 0; i < 10; i++ {
		result = process(result, input)
	}

	smallest, largest := countElements(result)

	return result, largest - smallest
}

func part2(template string, input map[string]string) int {
	result := template
	for i := 0; i < 30; i++ {
		result = process(result, input)
	}

	smallest, largest := countElements(result)

	return largest - smallest
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide a file name as argument")
	}

	template, input := readInput(os.Args[1])
	var diff int
	template, diff = part1(template, input)
	fmt.Println("Part1:", diff)
	fmt.Println("Part1:", part2(template, input))
}
