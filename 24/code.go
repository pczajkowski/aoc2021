package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type operation struct {
	op        string
	a         byte
	isBNumber bool
	b         int
	bC        byte
}

func readInput(file string) [][]operation {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")
	var input [][]operation
	var operations []operation
	for _, line := range lines {
		if line == "" {
			continue
		}

		parts := strings.Split(line, " ")
		if len(parts) < 2 {
			log.Fatal("Invalid line: ", line)
		}

		if parts[0] == "inp" && len(operations) > 0 {
			input = append(input, operations)
			operations = []operation{}
		}

		current := operation{op: parts[0], a: parts[1][0]}
		if len(parts) > 2 {
			if unicode.IsDigit(rune(parts[2][0])) {
				current.b, err = strconv.Atoi(parts[2])
				if err != nil {
					log.Fatal(err)
				}

				current.isBNumber = true
			} else {
				current.bC = parts[2][0]
			}
		}

		operations = append(operations, current)
	}

	if len(operations) > 0 {
		input = append(input, operations)
	}

	return input
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide a file name as argument")
	}

	input := readInput(os.Args[1])
	fmt.Println(input)
}
