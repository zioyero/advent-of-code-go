package days

import (
	"fmt"
	"strconv"
	"strings"

	"adventOfCode.com/m/v2/structures"
)

type Thirteen struct {
}

func (t *Thirteen) PartOne(lines []string) int {
	p := structures.NewPaper()
	positions, err := parsePositions(lines)
	if err != nil {
		fmt.Printf("error parsing lines to dot positions: %v\n", err)
		return -1
	}
	p.AddDots(positions)
	// fmt.Printf("Parsed Into Original State:\n%s", p.String())
	folds, err := parseFolds(lines)
	if err != nil {
		fmt.Printf("error parsing lines to fold instructions, %v\n", err)
	}
	folded := p.Fold(folds[0])

	return folded.DotCount()
}

func (t *Thirteen) PartTwo(lines []string) int {
	p := structures.NewPaper()
	positions, err := parsePositions(lines)
	if err != nil {
		fmt.Printf("error parsing lines to dot positions: %v\n", err)
		return -1
	}
	p.AddDots(positions)
	folds, err := parseFolds(lines)
	if err != nil {
		fmt.Printf("error parsing lines to fold instructions, %v\n", err)
	}
	folded := p
	for _, fold := range folds {
		x, y := folded.Size()
		fmt.Printf("Folding %v ; Current Size: %v,%v\n", fold, x, y)
		folded = folded.Fold(fold)
		fmt.Printf("Dots after fold %v:\n%v\n", fold, folded.Dots())
	}
	fmt.Printf("After Folding:\n%v", folded.String())

	return folded.DotCount()
}

func parsePositions(lines []string) ([]structures.Position, error) {
	positions := make([]structures.Position, 0)
	for _, line := range lines {
		tokens := strings.Split(line, ",")
		if len(tokens) != 2 {
			// this is an empty line or a fold directive
			continue
		}
		x, err := strconv.Atoi(tokens[0])
		if err != nil {
			return nil, err
		}
		y, err := strconv.Atoi(tokens[1])
		if err != nil {
			return nil, err
		}
		positions = append(positions, structures.NewPosition(x, y))
	}
	return positions, nil
}

func parseFolds(lines []string) ([]structures.Fold, error) {
	folds := make([]structures.Fold, 0)
	for _, line := range lines {
		if !strings.HasPrefix(line, "fold") {
			continue
		}
		tokens := strings.Split(line, " ")
		val := tokens[2]
		valTokens := strings.Split(val, "=")
		isX := valTokens[0] == "x"
		isY := !isX
		n, err := strconv.Atoi(valTokens[1])
		if err != nil {
			return nil, err
		}
		fold := structures.Fold{X: isX, Y: isY, Val: n}
		folds = append(folds, fold)
	}
	return folds, nil
}
