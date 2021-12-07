package main

import "testing"

func TestPart1(t *testing.T) {
	fish := readInput("./testinput")
	if len(fish) == 0 {
		t.Error("Could not read input")
	}

	if breed(fish, 80) != 5934 {
		t.Error("Part 1 failed")
	}
}

func TestPart2(t *testing.T) {
	fish := readInput("./testinput")
	if len(fish) == 0 {
		t.Error("Could not read input")
	}

	if breed(fish, 256) != 26984457539 {
		t.Error("Part 2 failed")
	}
}
