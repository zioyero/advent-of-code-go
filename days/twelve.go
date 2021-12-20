package days

import (
	"fmt"
	"strings"

	"adventOfCode.com/m/v2/algo"
	. "adventOfCode.com/m/v2/structures"
)

type Twelve struct{}

func (t Twelve) PartOne(lines []string) int {
	graph := makeGraph(lines)
	fmt.Printf("Graph:\n%s\n", graph.String())
	fmt.Printf("Nodes: %v\n", graph.Nodes)

	search := algo.DFS{Graph: graph, CanAdd: CanAddNodeP1}
	fmt.Println("Finding efficient routes...")
	start := Node{Value: "start"}
	end := Node{Value: "end"}
	paths := search.PathsBetween(start, end)
	fmt.Printf("Paths: %v\n", len(paths))
	fmt.Println("Routes Calculated!")

	return len(paths)
}

func CanAddNodeP1(p algo.Path, n Node) bool {
	if n.Value != strings.ToLower(n.Value) {
		return true
	}

	for _, pn := range p.Nodes() {
		if pn == n {
			return false
		}
	}

	return true
}

func (t Twelve) PartTwo(lines []string) int {
	graph := makeGraph(lines)
	fmt.Printf("Graph:\n%s\n", graph.String())
	fmt.Printf("Nodes: %v\n", graph.Nodes)

	search := algo.DFS{Graph: graph, CanAdd: CanAddNodeP2}
	fmt.Println("Finding efficient routes...")
	start := Node{Value: "start"}
	end := Node{Value: "end"}
	paths := search.PathsBetween(start, end)
	fmt.Printf("Paths: %v\n", len(paths))
	fmt.Println("Routes Calculated!")

	return len(paths)
}

func CanAddNodeP2(p algo.Path, n Node) bool {
	if len(p.Nodes()) > 0 && n.Value == "start" {
		return false
	}
	if n.Value != strings.ToLower(n.Value) {
		return true
	}

	// check for restrictions present in this path - only a single small cave may be visited twice
	smallCaves := make(map[string]int, len(p.Nodes()))
	smallCaves[n.Value]++
	for _, pathNode := range p.Nodes() {
		if pathNode.Value != strings.ToLower(pathNode.Value) {
			continue
		}
		smallCaves[pathNode.Value]++
		if smallCaves[pathNode.Value] > 2 {
			return false
		}
	}

	visitedCaveTwice := false
	for _, value := range smallCaves {
		if visitedCaveTwice && value == 2 {
			return false
		}
		visitedCaveTwice = value == 2 || visitedCaveTwice
	}

	return true
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
