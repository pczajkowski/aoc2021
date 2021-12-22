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

type cube struct {
	x, y, z int
}

func countOn(cubes map[cube]bool) int {
	count := 0
	for _, on := range cubes {
		if on {
			count++
		}
	}

	return count
}

func part1(input []action) int {
	cubes := make(map[cube]bool)
	for _, action := range input {
		startX := action.x[0]
		if startX < -50 {
			startX = -50
		}

		endX := action.x[1]
		if endX > 50 {
			endX = 50
		}

		for x := startX; x <= endX; x++ {
			startY := action.y[0]
			if startY < -50 {
				startY = -50
			}

			endY := action.y[1]
			if endY > 50 {
				endY = 50
			}

			for y := startY; y <= endY; y++ {
				startZ := action.z[0]
				if startZ < -50 {
					startZ = -50
				}

				endZ := action.z[1]
				if endZ > 50 {
					endZ = 50
				}

				for z := startZ; z <= endZ; z++ {
					cubes[cube{x, y, z}] = action.on
				}
			}
		}
	}

	return countOn(cubes)
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide a file name as argument")
	}

	input := readInput(os.Args[1])
	fmt.Println("Part1:", part1(input))
}
