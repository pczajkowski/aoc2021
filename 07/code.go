package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readInput(file string) (map[int]int, []int) {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	numberStrings := strings.Split(string(content), ",")
	crabs := make(map[int]int)
	var orderedCrabs []int
	for _, numberString := range numberStrings {
		if numberString == "" {
			continue
		}

		if number, err := strconv.Atoi(numberString); err == nil {
			if _, ok := crabs[number]; !ok {
				orderedCrabs = append(orderedCrabs, number)
			}
			crabs[number]++
		} else {
			log.Fatal(err)
		}
	}

	sort.Ints(orderedCrabs)
	return crabs, orderedCrabs
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide a file name as argument")
	}

	crabs, orderedCrabs := readInput(os.Args[1])
	fmt.Println(crabs, orderedCrabs)
}
