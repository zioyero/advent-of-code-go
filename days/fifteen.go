package days

import "strconv"

type Fifteen struct {
}

func (f Fifteen) PartOne(lines []string) int {
	riskMap := NewRiskMap(lines)
	path := riskMap.leastRiskyPath()
	return riskMap.riskLevel(path)
}

func (f Fifteen) PartTwo(lines []string) int {
	return 0
}

type RiskMap struct {
	riskMap map[RiskPosition]int
}

type RiskPosition struct {
	x int
	y int
}

func NewRiskMap(lines []string) RiskMap {
	riskMap := make(map[RiskPosition]int)
	for y, line := range lines {
		for x, riskVal := range line {
			pos := RiskPosition{x, y}
			level, err := strconv.Atoi(string(riskVal))
			if err != nil {
				return RiskMap{}
			}
			riskMap[pos] = level
		}
	}
	return RiskMap{riskMap: riskMap}
}

type RiskMapPath struct {
	positions []RiskPosition
}

func (r RiskMap) leastRiskyPath() RiskMapPath {
	// depth-first search through the map graph
	// eliminating routes that have greater risk
	// than the current minimum risk level
	return RiskMapPath{positions: make([]RiskPosition, 0)}
}

func (r RiskMap) riskLevel(path RiskMapPath) int {
	return 0
}
