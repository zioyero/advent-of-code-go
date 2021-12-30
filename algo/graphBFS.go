package algo

import (
	. "adventOfCode.com/m/v2/structures"
)

type CanAddNodeToPath func(p Path, n Node) bool
type IDepthFirstSearch interface {
	PathsBetween(n1 Node, n2 Node) [][]Node
}

type BFS struct {
	Graph    *Graph
	explored []Edge
	CanAdd   CanAddNodeToPath
}

type Path struct {
	nodes []Node
	edges []Edge
}

func (dfs *BFS) Clear() {
	dfs.explored = []Edge{}
}

// DFS

func (dfs *BFS) PathsBetween(n1 Node, n2 Node) []Path {
	return dfs.pathsBetween(n1, n2, Path{nodes: make([]Node, 0)})
}

func (dfs *BFS) pathsBetween(n1 Node, n2 Node, p Path) []Path {
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

	for _, n := range neighbors {
		edge := Edge{A: n1, B: n}
		if !dfs.CanAdd(p, n) {
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

func (p Path) Length() int {
	return len(p.nodes)
}

func (p Path) Nodes() []Node {
	return p.nodes
}
