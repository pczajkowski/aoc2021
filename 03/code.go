package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func readInput(file string) []string {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")
	return lines
}

func part1(lines []string) int64 {
	var gamma []string
	var epsilon []string

	for i := 0; i < len(lines[0]); i++ {
		counts := []int{0, 0}
		for _, line := range lines {
			if line[i] == '0' {
				counts[0]++
			} else {
				counts[1]++
			}
		}

		if counts[0] > counts[1] {
			gamma = append(gamma, "0")
			epsilon = append(epsilon, "1")
		} else {
			gamma = append(gamma, "1")
			epsilon = append(epsilon, "0")
		}
	}

	gammaNumber, err := strconv.ParseInt(strings.Join(gamma, ""), 2, 64)
	if err != nil {
		log.Fatal(err)
	}

	epsilonNumber, err := strconv.ParseInt(strings.Join(epsilon, ""), 2, 64)
	if err != nil {
		log.Fatal(err)
	}

	return gammaNumber * epsilonNumber
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("You need to provide path to input file!")
	}

	lines := readInput(os.Args[1])
	fmt.Println(part1(lines))
}
