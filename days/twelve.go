package days

import (
	"fmt"
	"strings"

	. "adventOfCode.com/m/v2/structures"
)

type Twelve struct{}

func (t Twelve) PartOne(lines []string) int {
	graph := makeGraph(lines)
	fmt.Printf("Graph:\n%s\n", graph.String())
	fmt.Printf("Nodes: %v\n", graph.Nodes)
	return 0
}

func (t Twelve) PartTwo(lines []string) int {
	return 0
}

func makeGraph(lines []string) *Graph {
	g := &Graph{}
	for _, line := range lines {
		n1, n2 := parseEdge(line)
		fmt.Printf("Adding edge %v -> %v\n", n1, n2)
		g.AddEdge(n1, n2)
	}
	return g
}

func parseEdge(line string) (Node, Node) {
	tokens := strings.Split(line, "-")
	a := Node{Value: tokens[0]}
	b := Node{Value: tokens[1]}
	return a, b
}
