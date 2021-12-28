package days

import (
	"fmt"
	"math"
	"strings"

	"adventOfCode.com/m/v2/structures"
)

type Fourteen struct{}

func (f *Fourteen) PartOne(lines []string) int {
	polymer := parsePolymer(lines, true)
	fmt.Printf("Processing Polymer %v\nRules:%v\n", polymer.Value(), polymer.Rules())
	for x := 0; x < 10; x++ {
		polymer.Process()
	}
	most := polymer.MostCommonElement()
	least := polymer.LeastCommonElement()
	mostCount := polymer.ElementCount(most)
	leastCount := polymer.ElementCount(least)
	fmt.Printf("Most Common: %v => %v | Least Common: %v => %v\n", most, mostCount, least, leastCount)
	return mostCount - leastCount
}

func (f *Fourteen) PartTwo(lines []string) int {
	polymer := parseComplexPolymer(lines)
	fmt.Printf("Processing Polymer %v\nRules:%v\n", polymer.Value(), polymer.Rules())

	initialCount := map[string]int{}
	for _, r := range polymer.Value() {
		initialCount[string(r)]++
	}
	growthCount := polymer.Grow(polymer.Value(), 40)

	for k, v := range initialCount {
		growthCount[k] += v
	}

	// most common element minus least common element
	most, least := 0, math.MaxInt64
	for _, v := range growthCount {
		if v > most {
			most = v
		}
		if v < least {
			least = v
		}
	}
	return most - least
}

func parsePolymer(lines []string, simple bool) structures.Polymer {
	template := lines[0]
	rules := make([]structures.Rule, 0, len(lines)-2)
	for _, line := range lines {
		tokens := strings.Split(line, " -> ")
		if len(tokens) != 2 {
			continue
		}
		pair := tokens[0]
		insertion := tokens[1]
		rule := structures.NewRule(pair, insertion)
		rules = append(rules, rule)
	}
	if simple {
		poly := structures.NewSimplePolymer(template, rules)
		return &poly
	} else {
		poly := structures.NewComplexPolymer(template, rules)
		return &poly
	}
}

func parseComplexPolymer(lines []string) structures.ComplexPolymer {
	template := lines[0]
	rules := make([]structures.Rule, 0, len(lines)-2)
	for _, line := range lines {
		tokens := strings.Split(line, " -> ")
		if len(tokens) != 2 {
			continue
		}
		pair := tokens[0]
		insertion := tokens[1]
		rule := structures.NewRule(pair, insertion)
		rules = append(rules, rule)
	}

	poly := structures.NewComplexPolymer(template, rules)
	return poly
}
