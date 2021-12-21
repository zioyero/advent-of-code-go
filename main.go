package main

import (
	"fmt"

	"adventOfCode.com/m/v2/days"
	"adventOfCode.com/m/v2/input"
)

func main() {
	fmt.Println("Happy Advent of Code!")
	// run(days.Eleven{}, input.Parse("data/sample/eleven.txt"))
	run(&days.Thirteen{}, input.Parse("data/day13.txt"))
}

func run(d days.Day, lines []string) {
	one := d.PartOne(lines)
	two := d.PartTwo(lines)
	fmt.Printf("Part 1: %d | Part 2: %d\n", one, two)
}
