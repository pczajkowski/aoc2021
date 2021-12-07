package main

import "testing"

func TestPart1(t *testing.T) {
	crabs, orderedCrabs := readInput("./testinput")
	if len(crabs) == 0 || len(orderedCrabs) == 0 {
		t.Error("Could not read input")
	}

	if calculate(crabs, orderedCrabs, costPart1) != 37 {
		t.Error("Part 1 failed")
	}
}

func TestPart2(t *testing.T) {
	crabs, orderedCrabs := readInput("./testinput")
	if len(crabs) == 0 || len(orderedCrabs) == 0 {
		t.Error("Could not read input")
	}

	if calculate(crabs, orderedCrabs, costPart2) != 168 {
		t.Error("Part 2 failed")
	}
}
