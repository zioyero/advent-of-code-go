package algo

import (
	"fmt"

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
}

func (dfs *DFS) Clear() {
	dfs.explored = []Edge{}
}

func (dfs *DFS) didAlreadyExplore(e Edge) bool {
	for _, explored := range dfs.explored {
		if explored.A.Value == e.A.Value && explored.B.Value == e.B.Value {
			return true
		}
	}
	return false
}

func (dfs *DFS) PathsBetween(n1 Node, n2 Node) []Path {
	return dfs.pathsBetween(n1, n2, Path{nodes: make([]Node, 0)})
}

func (dfs *DFS) pathsBetween(n1 Node, n2 Node, p Path) []Path {
	fmt.Printf("\n")
	// base case
	if len(p.nodes) == 0 {
		p.nodes = append(p.nodes, n1)
	}
	fmt.Printf("Path: %v\n", p)
	fmt.Printf("Exploring paths from %v\n", n1)
	neighbors := dfs.Graph.Edges[n1]
	fmt.Printf("Neighbors: %v\n", neighbors)
	// exit condition - we've reached the end
	if n1.Value == "end" {
		fmt.Printf("At end! Returning path %v\n", p)
		return []Path{p}
	}

	paths := make([]Path, 0)
OUTER:
	for _, n := range neighbors {
		if dfs.didAlreadyExplore(Edge{A: &n1, B: n}) {
			fmt.Printf("ignoring explored path %v\n", Edge{A: &n1, B: n})
			continue
		}
		for _, exploredNode := range p.nodes {
			if n.Value == exploredNode.Value {
				fmt.Printf("Ignoring backtrack path returning to %v", n)
				continue OUTER
			}
		}
		// make a new copy of this path to continue down
		next := make([]Node, len(p.nodes), len(p.nodes)+1)
		copy(next, p.nodes)
		// add this neighbor node to the path
		next = append(next, *n)
		// mark this edge as explored
		edge := Edge{A: &n1, B: n}
		dfs.explored = append(dfs.explored, edge)
		// continue exploring
		subpaths := dfs.pathsBetween(*n, n2, Path{nodes: next})
		paths = append(paths, subpaths...)
	}
	return paths
}
