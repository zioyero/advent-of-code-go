package days

import (
	"fmt"
	"strings"

	"adventOfCode.com/m/v2/structures"
)

type Fourteen struct{}

func (f *Fourteen) PartOne(lines []string) int {
	polymer := parsePolymer(lines)
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
	polymer := parsePolymer(lines)
	fmt.Printf("Processing Polymer %v\nRules:%v\n", polymer.Value(), polymer.Rules())
	for x := 0; x < 40; x++ {
		polymer.Process()
	}
	most := polymer.MostCommonElement()
	least := polymer.LeastCommonElement()
	mostCount := polymer.ElementCount(most)
	leastCount := polymer.ElementCount(least)
	fmt.Printf("Most Common: %v => %v | Least Common: %v => %v\n", most, mostCount, least, leastCount)
	return mostCount - leastCount
}

func parsePolymer(lines []string) structures.SimplePolymer {
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
	return structures.NewPolymer(template, rules)
}
