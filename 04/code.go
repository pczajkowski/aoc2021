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

func allTrue(slice []bool) bool {
	for _, value := range slice {
		if !value {
			return false
		}
	}

	return true
}

func checkWinner(board [][]Number, row int, col int) bool {
	rowCheck := make([]bool, 5)
	for i := 0; i < 5; i++ {
		rowCheck[i] = board[row][i].Marked
	}

	if allTrue(rowCheck) {
		return true
	}

	colCheck := make([]bool, 5)
	for i := 0; i < 5; i++ {
		colCheck[i] = board[i][col].Marked
	}

	return allTrue(colCheck)
}

func mark(boards [][][]Number, number int) *[][]Number {
	var winner *[][]Number
	for _, board := range boards {
		for i, row := range board {
			for j, _ := range row {
				if row[j].Val == number {
					row[j].Marked = true
					if checkWinner(board, i, j) {
						winner = &board
					}
				}
			}
		}
	}

	return winner
}

func calculateBoard(board *[][]Number) int {
	sum := 0
	for _, row := range *board {
		for _, number := range row {
			if !number.Marked {
				sum += number.Val
			}
		}
	}

	return sum
}

func part1(boards [][][]Number, numbers []int) int {
	lastNumber := 0
	sumOfUnmarkedNumbers := 0
	for _, number := range numbers {
		lastNumber = number
		winner := mark(boards, number)
		if winner != nil {
			sumOfUnmarkedNumbers = calculateBoard(winner)
			break
		}
	}

	return lastNumber * sumOfUnmarkedNumbers
}

func isWinner(boardIndex int, winners []int) bool {
	for _, winner := range winners {
		if boardIndex == winner {
			return true
		}
	}

	return false
}

func mark2(boards [][][]Number, number int, winners []int) (bool, []int) {
	newWinner := false
	for x, board := range boards {
		if isWinner(x, winners) {
			continue
		}

		for i, row := range board {
			for j, _ := range row {
				if row[j].Val == number {
					row[j].Marked = true
					if checkWinner(board, i, j) {
						winners = append(winners, x)
						newWinner = true
					}
				}
			}
		}
	}

	return newWinner, winners
}

func part2(boards [][][]Number, numbers []int) int {
	lastNumber := 0
	var winners []int
	newWinner := false
	for _, number := range numbers {
		newWinner, winners = mark2(boards, number, winners)
		if newWinner {
			lastNumber = number
		}
	}

	return lastNumber * calculateBoard(&boards[winners[len(winners)-1]])
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
	fmt.Println("Part 1: ", part1(boards, numbers))

	_, err = file.Seek(0, 0)
	if err != nil {
		log.Fatal("Seek: ", err)
	}

	boards, numbers = readInput(file)
	fmt.Println("Part 2: ", part2(boards, numbers))
}
