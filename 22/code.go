package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type action struct {
	on bool
	x  []int
	y  []int
	z  []int
}

func readInput(file string) []action {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")
	var input []action
	for _, line := range lines {
		if line == "" {
			continue
		}

		var a action
		if strings.HasPrefix(line, "on") {
			a.on = true
		}

		trash := strings.Split(line, " ")
		if len(trash) != 2 {
			log.Fatal("Invalid input")
		}

		parts := strings.Split(trash[1], ",")
		if len(parts) != 3 {
			log.Fatal("Invalid input")
		}

		for i, part := range parts {
			another := strings.Split(part, "=")
			if len(another) != 2 {
				log.Fatal("Invalid input")
			}

			numbers := strings.Split(another[1], "..")
			if len(numbers) != 2 {
				log.Fatal("Invalid input")
			}

			var first, second int
			if first, err = strconv.Atoi(numbers[0]); err != nil {
				log.Fatal(err)
			}

			if second, err = strconv.Atoi(numbers[1]); err != nil {
				log.Fatal(err)
			}

			switch i {
			case 0:
				a.x = []int{first, second}
			case 1:
				a.y = []int{first, second}
			case 2:
				a.z = []int{first, second}
			}
		}

		input = append(input, a)
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
