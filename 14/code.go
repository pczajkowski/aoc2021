package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func readInput(file string) (map[string]int, map[byte]int, map[string][]string) {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")
	template := make(map[string]int)
	input := make(map[string][]string)
	count := make(map[byte]int)
	readingTemplate := true
	for _, line := range lines {
		if line == "" {
			readingTemplate = false
			continue
		}

		if readingTemplate {
			for i := 0; i < len(line); i++ {
				count[line[i]]++
			}

			for i := 0; i < len(line)-1; i++ {
				template[string(line[i:i+2])] += 1
			}
			continue
		}

		key, value, found := strings.Cut(line, " -> ")
		if !found {
			log.Fatal("Invalid line: ", line)
		}

		input[key] = []string{fmt.Sprintf("%c%s", key[0], value), fmt.Sprintf("%s%c", value, key[1])}
	}

	return template, count, input
}

func countElements(count map[byte]int) (int, int) {
	smallest := count['N']
	largest := count['N']
	for _, v := range count {
		if v < smallest {
			smallest = v
		}

		if v > largest {
			largest = v
		}
	}

	return smallest, largest
}

func process(template map[string]int, count map[byte]int, input map[string][]string, rounds int) (map[string]int, map[byte]int, int) {
	for i := 0; i < rounds; i++ {
		newTemplate := make(map[string]int)
		for k, v1 := range template {
			for _, v2 := range input[k] {
				newTemplate[v2] += v1
			}

			count[input[k][0][1]] += v1
		}

		template = newTemplate
	}

	smallest, largest := countElements(count)
	return template, count, largest - smallest
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide a file name as argument")
	}

	template, count, input := readInput(os.Args[1])
	var diff int
	template, count, diff = process(template, count, input, 10)
	fmt.Println("Part1:", diff)

	_, _, diff = process(template, count, input, 30)
	fmt.Println("Part2:", diff)
}
