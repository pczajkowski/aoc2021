package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func readInput(file string) [][][]int {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")
	var points [][][]int
	for _, line := range lines {
		if line == "" {
			continue
		}

		pointsString := strings.Split(line, " -> ")
		if len(pointsString) != 2 {
			log.Fatal("Invalid line: ", line)
		}

		start := strings.Split(pointsString[0], ",")
		if len(start) != 2 {
			log.Fatal("Invalid start point: ", start)
		}

		startPoint := make([]int, 2)
		startPoint[0], err = strconv.Atoi(start[0])
		if err != nil {
			log.Fatal(err)
		}
		startPoint[1], err = strconv.Atoi(start[1])
		if err != nil {
			log.Fatal(err)
		}

		end := strings.Split(pointsString[1], ",")
		if len(end) != 2 {
			log.Fatal("Invalid end point: ", end)
		}
		endPoint := make([]int, 2)
		endPoint[0], err = strconv.Atoi(end[0])
		if err != nil {
			log.Fatal(err)
		}
		endPoint[1], err = strconv.Atoi(end[1])
		if err != nil {
			log.Fatal(err)
		}

		points = append(points, [][]int{startPoint, endPoint})
	}

	return points
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide a file name as argument")
	}

	input := readInput(os.Args[1])
	fmt.Println(input)
}
