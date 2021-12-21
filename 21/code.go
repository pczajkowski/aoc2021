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

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide a file name as argument")
	}

	player1, player2 := readInput(os.Args[1])
	fmt.Println(player1, player2)
}
