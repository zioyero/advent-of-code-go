package days

import (
	"fmt"
	"sort"
	"strings"
)

type Day interface {
	PartOne(lines []string) int
	PartTwo(lines []string) int
}

type Ten struct {
}

func (t Ten) PartOne(lines []string) int {
	var sum int
	for _, line := range lines {
		symbols := strings.Split(line, "")
		errorValue, _ := syntaxCheck(symbols)
		sum += errorValue
	}
	return sum
}

func (t Ten) PartTwo(lines []string) int {
	scores := make([]int, 0, len(lines))
	for _, line := range lines {
		symbols := strings.Split(line, "")
		_, missingOperands := syntaxCheck(symbols)
		if len(missingOperands) > 0 {
			score := autocompleteScore(missingOperands)
			scores = append(scores, score)
			// fmt.Printf("autocompleting %s with %s => %d points\n", line, missingOperands, score)
		}
	}
	// get middle score
	sort.Ints(scores)
	fmt.Printf("Scores: %d\n", scores)
	return scores[len(scores)/2]
}

func syntaxCheck(symbols []string) (score int, missingOperands string) {
	// fmt.Printf("Syntax checking symbols %s\n", symbols)
	// start empty stack
	openOperands := make([]string, len(symbols))
	for _, operand := range symbols {
		if isOpenOperand(operand) {
			// if it's an open-type operand, we can store this in our openOperand stack
			openOperands = append(openOperands, operand)
		} else if matchesOpenOperand(operand, openOperands[len(openOperands)-1]) {
			// if it matches the current open operand, we can pop the last and move on
			openOperands = openOperands[:len(openOperands)-1]
		} else {
			// otherwise this is a syntax error
			score = illegalCharacterScore(operand)
			return score, missingOperands
		}
	}
	missingOperands = autocomplete(openOperands)
	return score, missingOperands
}

func autocompleteScore(insertedOperands string) int {
	score := 0
	for _, operand := range insertedOperands {
		score = score*5 + autocompleteCharacterScore(string(operand))
	}
	return score
}

func autocomplete(openOperands []string) string {
	var missing string
	for i := len(openOperands) - 1; i >= 0; i-- {
		missing += expectedCloseOperand(openOperands[i])
	}
	return missing
}

func isOpenOperand(operand string) bool {
	return operand == "{" || operand == "[" || operand == "<" || operand == "("
}

func matchesOpenOperand(operand string, openOperand string) bool {
	return operand == expectedCloseOperand(openOperand)
}

func expectedCloseOperand(openOperand string) string {
	switch openOperand {
	case "{":
		return "}"
	case "[":
		return "]"
	case "<":
		return ">"
	case "(":
		return ")"
	default:
		return ""
	}
}

func illegalCharacterScore(operand string) int {
	switch operand {
	case "}":
		return 1197
	case "]":
		return 57
	case ">":
		return 25137
	case ")":
		return 3
	default:
		return 0
	}
}

func autocompleteCharacterScore(operand string) int {
	switch operand {
	case "}":
		return 3
	case "]":
		return 2
	case ">":
		return 4
	case ")":
		return 1
	default:
		return 0
	}
}
