package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type point struct {
	min, max int
}

func readInput(file string) []point {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	var points []point
	line := string(content)
	initialSplit := strings.Split(line, ", ")
	if len(initialSplit) != 2 {
		log.Fatal("Invalid input")
	}

	firstPartSplit := strings.Split(initialSplit[0], "=")
	if len(firstPartSplit) != 2 {
		log.Fatal("Invalid input")
	}

	firstPartNumbers := strings.Split(firstPartSplit[1], "..")
	if len(firstPartNumbers) != 2 {
		log.Fatal("Invalid input")
	}

	minX, err := strconv.Atoi(firstPartNumbers[0])
	if err != nil {
		log.Fatal(err)
	}

	maxX, err := strconv.Atoi(firstPartNumbers[1])
	if err != nil {
		log.Fatal(err)
	}

	points = append(points, point{minX, maxX})

	secondPartNumbers := strings.Split(initialSplit[1], "..")
	if len(secondPartNumbers) != 2 {
		log.Fatal("Invalid input")
	}

	minY, err := strconv.Atoi(strings.TrimLeft(secondPartNumbers[0], "y="))
	if err != nil {
		log.Fatal(err)
	}

	maxY, err := strconv.Atoi(secondPartNumbers[1])
	if err != nil {
		log.Fatal(err)
	}

	points = append(points, point{minY, maxY})
	return points
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide a file name as argument")
	}

	input := readInput(os.Args[1])
	fmt.Println(input)
}
