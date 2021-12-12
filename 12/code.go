package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"unicode"
)

func readInput(file string) map[string][]string {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")
	input := make(map[string][]string)
	for _, line := range lines {
		if line == "" {
			continue
		}

		points := strings.Split(line, "-")
		if len(points) != 2 {
			log.Fatal("Invalid input")
		}

		input[points[0]] = append(input[points[0]], points[1])
		input[points[1]] = append(input[points[1]], points[0])
	}

	return input
}

var paths [][]string

func findAllPaths(start string, end string, visited map[string]bool, input map[string][]string, localPath []string) {
	if start == end {
		paths = append(paths, localPath)
		return
	}

	if unicode.IsLower(rune(start[0])) {
		visited[start] = true
	}

	for _, next := range input[start] {
		if !visited[next] {
			localPath = append(localPath, next)
			findAllPaths(next, end, visited, input, localPath)
			localPath = localPath[:len(localPath)-1]
		}
	}

	if visited[start] {
		visited[start] = false
	}
}

func part1(input map[string][]string) int {
	visited := make(map[string]bool)
	localPath := []string{"start"}
	findAllPaths("start", "end", visited, input, localPath)

	return len(paths)
}

func findAllPathsPart2(start string, end string, visited map[string]int, input map[string][]string, localPath []string) {
	if start == end {
		paths = append(paths, localPath)
		fmt.Println(localPath)
		return
	}

	if unicode.IsLower(rune(start[0])) {
		visited[start]++
	}

	for _, next := range input[start] {
		if next != "start" && visited[next] < 2 {
			localPath = append(localPath, next)
			findAllPathsPart2(next, end, visited, input, localPath)
			localPath = localPath[:len(localPath)-1]
		}
	}

	if visited[start] > 0 {
		visited[start]--
	}
}

func part2(input map[string][]string) int {
	visited := make(map[string]int)

	localPath := []string{"start"}
	paths = make([][]string, 0)
	findAllPathsPart2("start", "end", visited, input, localPath)

	return len(paths)
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide a file name as argument")
	}

	input := readInput(os.Args[1])
	fmt.Println("Part 1:", part1(input))
	fmt.Println("Part 2:", part2(input))
}
