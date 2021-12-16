package algo

import . "adventOfCode.com/m/v2/structures"

type IDepthFirstSearch interface {
	PathsBetween(n1 Node, n2 Node) [][]Node
}

type DFS struct {
	Graph    Graph
	explored []Edge
}

func (dfs *DFS) Clear() {
	dfs.explored = []Edge{}
}

func (dfs *DFS) PathsBetween(n1 Node, n2 Node) [][]Node {
	// adjacentNodes := dfs.Graph.GetAdjacentNodes(n1)
	return [][]Node{}
}
