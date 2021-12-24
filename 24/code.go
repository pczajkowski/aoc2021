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

		if parts[0] == "inp" {
			if len(operations) > 0 {
				input = append(input, operations)
				operations = []operation{}
			}

			continue
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

func do(action operation, variables map[byte]int) bool {
	switch action.op {
	case "add":
		if action.isBNumber {
			variables[action.a] += action.b
		} else {
			variables[action.a] += variables[action.bC]
		}
		return true
	case "mul":
		if action.isBNumber {
			variables[action.a] *= action.b
		} else {
			variables[action.a] *= variables[action.bC]
		}
		return true
	case "div":
		if action.isBNumber {
			if action.b == 0 {
				return false
			}

			variables[action.a] /= action.b
		} else {
			if variables[action.bC] == 0 {
				return false
			}

			variables[action.a] /= variables[action.bC]
		}
		return true
	case "mod":
		if variables[action.a] < 0 {
			return false
		}

		if action.isBNumber {
			if action.b <= 0 {
				return false
			}

			variables[action.a] %= action.b
		} else {
			if variables[action.bC] <= 0 {
				return false
			}

			variables[action.a] %= variables[action.bC]
		}
		return true
	case "eql":
		if action.isBNumber {
			if variables[action.a] == action.b {
				variables[action.a] = 1
			} else {
				variables[action.a] = 0
			}
		} else {
			if variables[action.a] == variables[action.bC] {
				variables[action.a] = 1
			} else {
				variables[action.a] = 0
			}
		}
		return true
	}

	return false
}

func printVariables(variables map[byte]int) string {
	var result string
	for k, v := range variables {
		result += fmt.Sprintf("%c: %d ", k, v)
	}

	return result
}

func doSequence(sequence []operation, variables map[byte]int) bool {
	for _, action := range sequence {
		if !do(action, variables) {
			return false
		}
		fmt.Println(action, printVariables(variables))
	}

	return true
}

func part1(input [][]operation) []int {
	var number []int
	for i := 0; i < 1; i++ {
		for j := 9; j >= 1; j-- {
			variables := map[byte]int{}
			variables['w'] = j
			if !doSequence(input[i], variables) {
				fmt.Println("Failed for ", j)
			}
			fmt.Println(printVariables(variables))

			if variables['z'] == 0 {
				number = append(number, j)
				break
			}
		}
	}

	return number
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide a file name as argument")
	}

	input := readInput(os.Args[1])
	fmt.Println("Part1:", part1(input))
}
