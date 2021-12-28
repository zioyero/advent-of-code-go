package structures

import (
	"fmt"
	"math"
	"os"
	"strings"
)

type ComplexPolymer struct {
	value                string
	rules                map[string]string
	ruleList             []Rule
	precalculated        map[string]string
	longestPrecalculated []string
	stats                Stats
	generation           int
	memo                 map[string]map[string]int
}

type Stats struct {
	precalculated     int
	length            int
	hits              int
	miss              int
	insert            int
	maxSequenceLength int
}

func NewComplexPolymer(template string, rules []Rule) ComplexPolymer {
	r := make(map[string]string, len(rules))
	precalculated := make(map[string]string, 0)
	longest := make([]string, 0)
	for _, rule := range rules {
		r[rule.Pair] = rule.Insertion
	}
	return ComplexPolymer{
		value:                template,
		rules:                r,
		ruleList:             rules,
		memo:                 make(map[string]map[string]int),
		precalculated:        precalculated,
		longestPrecalculated: longest}
}

func AsSimplePolymer(complex ComplexPolymer) SimplePolymer {
	return SimplePolymer{value: complex.value, rules: complex.rules, ruleList: complex.ruleList}
}

func (p *ComplexPolymer) Stats() Stats {
	return Stats{
		hits:              p.stats.hits,
		length:            len(p.value),
		precalculated:     len(p.precalculated),
		miss:              p.stats.miss,
		insert:            p.stats.insert,
		maxSequenceLength: len(p.longestPrecalculated[len(p.longestPrecalculated)-1]),
	}
}

func (p *ComplexPolymer) Process() {
	// break up template into sections
	p.stats = Stats{}
	p.generation++
	fmt.Printf("\n\n======================================= GEN %v\n", p.generation)
	sections := p.sections(p.value)
	results := make(map[string]string)
	builder := strings.Builder{}
	for _, section := range sections {
		processedSection := p.processSection(section)
		builder.WriteString(processedSection)
		results[section] = processedSection
	}
	builder.WriteByte(p.value[len(p.value)-1])
	// simple := AsSimplePolymer(*p)
	// simple.Process()
	p.value = builder.String()
	// if simple.value != p.value {
	// 	fmt.Printf("Simple disagrees!\n")
	// 	os.Exit(1)
	// }
}

func (p *ComplexPolymer) Grow(sequence string, stepsLeft int) (addlCounts map[string]int) {
	addlCounts = map[string]int{}

	if stepsLeft == 0 {
		return addlCounts
	}

	key := fmt.Sprint(sequence, stepsLeft)
	if res, ok := p.memo[key]; ok {
		return res
	}

	for i := 0; i < len(sequence)-1; i++ {
		pair := sequence[i : i+2]
		between := p.rules[pair]
		addlCounts[between]++

		// get the additional characters for recursing just this three character section
		// calling grow will memoize that result to eliminate duplicate work
		recurse := p.Grow(pair[:1]+between+pair[1:], stepsLeft-1)

		// merge those additional characters into this (parent function call) addlCounts map
		for k, v := range recurse {
			addlCounts[k] += v
		}
	}

	// store it, return it
	p.memo[key] = addlCounts
	return addlCounts
}

func (p *ComplexPolymer) processSection(section string) string {
	if len(strings.TrimSpace(section)) == 0 {
		fmt.Printf("Error: empty section!")
		os.Exit(1)
	} else if len(section) == 1 {
		return ""
	}

	memoized, ok := p.precalculated[section]
	if ok {
		p.stats.hits++
		return memoized
	}
	p.stats.miss++
	if len(section) == 2 {
		processedPair := p.processPair(section)
		p.memoize(section, processedPair)
		return processedPair
	}
	subsections := p.sections(section)
	builder := strings.Builder{}
	for _, subsection := range subsections {
		// process the subsection
		processed := p.processSection(subsection)
		// memoize the processed subsection
		p.memoize(subsection, processed)
		// add the processed subsection to the result
		builder.WriteString(processed)
	}
	// compute the result
	processed := builder.String()
	// memoize the whole section as well
	p.memoize(section, processed)
	return processed
}

func (p *ComplexPolymer) sections(value string) []string {
	// fmt.Printf("Breaking up Section %v\n", value)
	c := int(math.Pow(3, math.Sqrt(float64(p.generation))))
	l := len(value)
	remainder := l % c
	s := l / c
	if s == 0 {
		return pairs(value)
	}
	chunks := make([]string, 0)
	for x := 0; x < c; x++ {
		chunk := value[x*s : (x+1)*s]
		if x > 0 {
			joiner := joiner(chunks[len(chunks)-1], chunk)
			chunks = append(chunks, joiner)
		}
		chunks = append(chunks, chunk)
	}
	if remainder > 0 {
		idx := c * s
		lastChunk := value[idx : idx+remainder]
		joiner := joiner(chunks[len(chunks)-1], lastChunk)
		chunks = append(chunks, joiner)
		chunks = append(chunks, lastChunk)
	}
	// fmt.Printf("Section %v => Subsections: %v\n", value, chunks)
	return chunks
}

func joiner(a string, b string) string {
	// fmt.Printf("joining %v and %v\n", a, b)
	return string(a[len(a)-1]) + string(b[0])
}

func pairs(value string) []string {
	pairs := make([]string, 0, len(value))
	for i, v := range value {
		if i+1 >= len(value) {
			continue
		}
		pair := string(v) + string(value[i+1])
		pairs = append(pairs, pair)
	}
	return pairs
}

func (p *ComplexPolymer) processPair(pair string) string {
	insertion, ok := p.rules[pair]
	if !ok {
		return pair[:1]
	}
	return pair[:1] + insertion
}

func (p *ComplexPolymer) memoize(sequence string, processed string) {
	p.stats.insert++
	p.precalculated[sequence] = processed
	var longestPrecalculatedSequence string
	if len(p.longestPrecalculated) > 0 {
		longestPrecalculatedSequence = p.longestPrecalculated[len(p.longestPrecalculated)-1]
	}
	if len(sequence) >= len(longestPrecalculatedSequence) {
		p.longestPrecalculated = append(p.longestPrecalculated, sequence)
	}
}

func (p *ComplexPolymer) countElements() map[string]int {
	elements := make(map[string]int)
	for _, element := range p.value {
		elements[string(element)]++
	}
	return elements
}

func (p *ComplexPolymer) MostCommonElement() string {
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

func (p *ComplexPolymer) LeastCommonElement() string {
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

func (p *ComplexPolymer) ElementCount(element string) int {
	return p.countElements()[element]
}

func (p *ComplexPolymer) Value() string {
	return p.value
}

func (p *ComplexPolymer) Rules() []Rule {
	return p.ruleList
}

func (s Stats) String() string {
	return fmt.Sprintf("{hits: %v, misses: %v, inserts: %v, precalculated: %v, polymerLength: %v, maxSequence: %v}", s.hits, s.miss, s.insert, s.precalculated, s.length, s.maxSequenceLength)
}
