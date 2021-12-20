package structures_test

import (
	"testing"

	"adventOfCode.com/m/v2/structures"
)

func Test_AddEdge(t *testing.T) {
	g := structures.Graph{}
	n1 := structures.Node{Value: "a"}
	n2 := structures.Node{Value: "b"}
	g.AddEdge(n1, n2)

	if len(g.Edges) != 2 {
		t.Error("expected 2 edges in an undirected graph with 2 connected nodes")
	}
}

func Test_ContainsNode(t *testing.T) {
	g := structures.Graph{}
	n1 := structures.Node{Value: "a"}
	n2 := structures.Node{Value: "b"}
	g.AddEdge(n1, n2)

	if !g.Contains(n1) || !g.Contains(n2) {
		t.Error("expected graph to contain node after adding an edge connected to it")
	}
}

func Test_GetNode(t *testing.T) {
	g := structures.Graph{}
	n1 := structures.Node{Value: "a"}
	n2 := structures.Node{Value: "b"}
	g.AddEdge(n1, n2)

	gottenNode, ok := g.GetNode("a")
	if !ok {
		t.Error("GetNode not ok")
	}
	if n1 != gottenNode {
		t.Error("Node retrieved from GetNode != Node inserted")
	}
}
