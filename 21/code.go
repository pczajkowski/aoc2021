package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func readInput(file string) (int, int) {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")
	if len(lines) != 2 {
		log.Fatal("expected 2 lines")
	}

	partsFirstLine := strings.Split(lines[0], ": ")
	if len(partsFirstLine) != 2 {
		log.Fatal("expected 2 parts")
	}

	player1, err := strconv.Atoi(partsFirstLine[1])
	if err != nil {
		log.Fatal(err)
	}

	partsSecondLine := strings.Split(lines[1], ": ")
	if len(partsSecondLine) != 2 {
		log.Fatal("expected 2 parts")
	}

	player2, err := strconv.Atoi(partsSecondLine[1])
	if err != nil {
		log.Fatal(err)
	}

	return player1, player2
}

var dice int

func roll() int {
	dice++
	return dice
}

func part1(player1, player2 int) int {
	var score1, score2 int
	first := true
	countDice := 0

	for {
		if score1 >= 1000 || score2 >= 1000 {
			break
		}

		if first {
			player1 = (roll() + roll() + roll() + player1) % 10
			if player1 == 0 {
				player1 = 10
			}

			score1 += player1
			first = false
			countDice += 3
		} else {
			player2 = (roll() + roll() + roll() + player2) % 10
			if player2 == 0 {
				player2 = 10
			}

			score2 += player2
			first = true
			countDice += 3
		}
	}

	if score1 > score2 {
		return score2 * countDice
	}

	return score1 * countDice
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide a file name as argument")
	}

	player1, player2 := readInput(os.Args[1])
	dice = 0
	fmt.Println("Part1:", part1(player1, player2))
}
