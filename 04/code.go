package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readNumbers(line string) []int {
	var numbers []int
	numbersStrings := strings.Split(line, ",")
	for _, numberString := range numbersStrings {
		if number, err := strconv.Atoi(numberString); err == nil {
			numbers = append(numbers, number)
		} else {
			log.Fatal("Numbers: ", err)
		}
	}

	return numbers
}

func readRow(line string) []int {
	var numbers []int
	numbersStrings := strings.Split(line, " ")
	fmt.Println(numbersStrings)
	for _, numberString := range numbersStrings {
		if numberString == "" {
			continue
		}

		if number, err := strconv.Atoi(numberString); err == nil {
			numbers = append(numbers, number)
		} else {
			log.Fatal("Row: ", err, numberString)
		}
	}

	return numbers
}

type Number struct {
	Val    int
	Marked bool
}

func readInput(file *os.File) ([][][]Number, []int) {
	scanner := bufio.NewScanner(file)
	numbersRead := false
	var numbers []int
	var boards [][][]Number
	boardIndex := 0
	rowIndex := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		if !numbersRead {
			numbers = readNumbers(line)
			numbersRead = true
			continue
		}

		if rowIndex == 0 {
			boards = append(boards, make([][]Number, 5))
		}

		boards[boardIndex][rowIndex] = make([]Number, 5)
		fmt.Println(line)
		numbersInRow := readRow(line)
		for i, number := range numbersInRow {
			boards[boardIndex][rowIndex][i] = Number{number, false}
		}

		rowIndex++
		if rowIndex > 4 {
			rowIndex = 0
			boardIndex++
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Scanner error: %s", err)
	}

	return boards, numbers
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("You need to specify a file!")
	}

	filePath := os.Args[1]
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open %s!\n", filePath)

	}

	boards, numbers := readInput(file)
	fmt.Println(boards, numbers)
}
