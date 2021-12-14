package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func readInput(file string) (string, map[string]string) {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")
	var template string
	input := make(map[string]string)
	readingTemplate := true
	for _, line := range lines {
		if line == "" {
			readingTemplate = false
			continue
		}

		if readingTemplate {
			template = line
			continue
		}

		parts := strings.Split(line, " -> ")
		if len(parts) != 2 {
			log.Fatal("Invalid line: ", line)
		}

		input[parts[0]] = parts[1]
	}

	return template, input
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide a file name as argument")
	}

	template, input := readInput(os.Args[1])
	fmt.Println(template, input)
}