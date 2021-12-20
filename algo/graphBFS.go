package algo

import (
	"fmt"
	"strings"

	. "adventOfCode.com/m/v2/structures"
)

type IDepthFirstSearch interface {
	PathsBetween(n1 Node, n2 Node) [][]Node
}

type DFS struct {
	Graph    *Graph
	explored []Edge
}

type Path struct {
	nodes []Node
	edges []Edge
}

func (dfs *DFS) Clear() {
	dfs.explored = []Edge{}
}

// DFS

func (dfs *DFS) PathsBetween(n1 Node, n2 Node) []Path {
	return dfs.pathsBetween(n1, n2, Path{nodes: make([]Node, 0)})
}

func (dfs *DFS) pathsBetween(n1 Node, n2 Node, p Path) []Path {
	// base case
	if len(p.nodes) == 0 {
		p.nodes = append(p.nodes, n1)
	}
	neighbors := dfs.Graph.Edges[n1]
	// exit condition - we've reached the end
	if n1 == n2 {
		return []Path{p}
	}

	paths := make([]Path, 0)

	fmt.Printf("\n-------\nPath: %v: Options: %v\n\n", p.nodes, neighbors)

	for _, n := range neighbors {
		edge := Edge{A: n1, B: n}
		if p.RestrictedNextNode(n) {
			continue
		}
		// make a new copy of this path to continue down
		pathNodes := make([]Node, len(p.nodes), len(p.nodes)+1)
		pathEdges := make([]Edge, len(p.edges), len(p.edges)+1)
		copy(pathNodes, p.nodes)
		copy(pathEdges, p.edges)
		// add this neighbor node to the path
		pathNodes = append(pathNodes, n)
		pathEdges = append(pathEdges, edge)
		// continue exploring
		updatedPath := Path{nodes: pathNodes, edges: pathEdges}
		subpaths := dfs.pathsBetween(n, n2, updatedPath)

		paths = append(paths, subpaths...)
	}
	return paths
}

// Path

func (p Path) RestrictedNextNode(n Node) bool {
	fmt.Printf("Adding node %v to ", n)
	if len(p.nodes) > 0 && n.Value == "start" {
		fmt.Printf("Path %v is restricted because it visits 'start' more than once\n", p.Nodes())
		return true
	}
	if n.Value != strings.ToLower(n.Value) {
		fmt.Println("current path allowed because it is uppercase")
		return false
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
			fmt.Printf("Path %v is restricted because it visits %v more than twice\n", p.Nodes(), pathNode)
			return true
		}
	}

	visitedCaveTwice := false
	for key, value := range smallCaves {
		if visitedCaveTwice && value == 2 {
			fmt.Printf("Path %v is restricted because it visits %v more than once\n", p.nodes, key)
			return true
		}
		visitedCaveTwice = value == 2 || visitedCaveTwice
	}

	fmt.Printf("current path %v allowed\n", p.nodes)
	return false
}

func (p Path) Length() int {
	return len(p.nodes)
}

func (p Path) Nodes() []Node {
	return p.nodes
}
