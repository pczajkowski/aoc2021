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

		parts := strings.Split(line, " -> ")
		if len(parts) != 2 {
			log.Fatal("Invalid line: ", line)
		}

		input[parts[0]] = []string{fmt.Sprintf("%c%s", parts[0][0], parts[1]), fmt.Sprintf("%s%c", parts[1], parts[0][1])}
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

func printCounts(counts map[byte]int) {
	for k, v := range counts {
		fmt.Printf("%c: %d;", k, v)
	}
	fmt.Println()
}

func part1(template map[string]int, count map[byte]int, input map[string][]string) (map[string]int, int) {
	fmt.Println(template)
	printCounts(count)

	for i := 0; i < 10; i++ {
		newTemplate := make(map[string]int)
		for k, v1 := range template {
			for _, v2 := range input[k] {
				newTemplate[v2] += v1
			}

			count[input[k][0][1]] += v1
		}

		template = newTemplate
		printCounts(count)
	}

	smallest, largest := countElements(count)
	fmt.Println(smallest, largest)
	return template, largest - smallest
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide a file name as argument")
	}

	template, count, input := readInput(os.Args[1])
	var diff int
	_, diff = part1(template, count, input)
	fmt.Println("Part1:", diff)
	//fmt.Println("Part1:", part2(template, input))
}
