package structures

import (
	"math"
)

type Polymer interface {
	Value() string
	Rules() []Rule
	Process()
	MostCommonElement() string
	LeastCommonElement() string
	ElementCount(element string) int
	Stats() Stats
}

type SimplePolymer struct {
	value    string
	rules    map[string]string
	ruleList []Rule
}

type Rule struct {
	Pair      string
	Insertion string
}

func NewSimplePolymer(template string, rules []Rule) SimplePolymer {
	r := make(map[string]string, len(rules))
	for _, rule := range rules {
		r[rule.Pair] = rule.Insertion
	}
	return SimplePolymer{value: template, rules: r, ruleList: rules}
}

func NewRule(pair string, insertion string) Rule {
	return Rule{Pair: pair, Insertion: insertion}
}

func (p *SimplePolymer) Stats() Stats {
	return Stats{}
}

func (p *SimplePolymer) Process() {
	// break up template into pairs
	pairs := p.pairs()
	// fmt.Printf("Processing Polymer %v...\nPairs: %v\n", p.value, pairs)
	processedValue := ""
	for _, pair := range pairs {
		processedPair := p.processPair(pair)
		processedValue += processedPair
	}
	processedValue += string(p.value[len(p.value)-1])
	p.value = processedValue
}

func (p *SimplePolymer) processPair(pair string) string {
	insertion, ok := p.rules[pair]
	if !ok {
		return pair[:1]
	}
	return pair[:1] + insertion
}

func (p *SimplePolymer) pairs() []string {
	pairs := make([]string, 0, len(p.value))
	for i, v := range p.value {
		if i+1 >= len(p.value) {
			continue
		}
		pair := string(v) + string(p.value[i+1])
		pairs = append(pairs, pair)
	}
	return pairs
}

func (p *SimplePolymer) countElements() map[string]int {
	elements := make(map[string]int)
	for _, element := range p.value {
		elements[string(element)]++
	}
	return elements
}

func (p *SimplePolymer) MostCommonElement() string {
	mostCommon := ""
	mostCommonCount := 0
	for key, value := range p.countElements() {
		if value > mostCommonCount {
			mostCommon = key
			mostCommonCount = value
		}
	}
	return mostCommon
}

func (p *SimplePolymer) LeastCommonElement() string {
	leastCommon := ""
	leastCommonCount := math.MaxInt
	for key, value := range p.countElements() {
		if value < leastCommonCount {
			leastCommon = key
			leastCommonCount = value
		}
	}
	return leastCommon
}

func (p *SimplePolymer) ElementCount(element string) int {
	return p.countElements()[element]
}

func (p *SimplePolymer) Value() string {
	return p.value
}

func (p *SimplePolymer) Rules() []Rule {
	return p.ruleList
}
