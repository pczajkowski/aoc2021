package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func readInput(file string) []int {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	fishStrings := strings.Split(string(content), ",")
	var input []int
	for _, fishString := range fishStrings {
		if fishString == "" {
			continue
		}

		if number, err := strconv.Atoi(fishString); err == nil {
			input = append(input, number)
		} else {
			log.Fatal(err)
		}
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
