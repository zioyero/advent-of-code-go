package algo_test

import (
	"testing"

	"adventOfCode.com/m/v2/algo"
	"adventOfCode.com/m/v2/structures"
)

func Test_PathsBetween(t *testing.T) {
	g := structures.Graph{}
	a := structures.Node{Value: "a"}
	b := structures.Node{Value: "b"}
	c := structures.Node{Value: "c"}
	d := structures.Node{Value: "d"}
	e := structures.Node{Value: "e"}
	g.AddEdge(a, b)

	dfs := algo.DFS{Graph: &g}
	/*
	 * a--b
	 */
	paths := dfs.PathsBetween(a, b)

	if len(paths) != 1 {
		t.Error("there should be a single path between a -- b, actual path count: ", len(paths))
	}

	dfs.Clear()
	g.AddEdge(b, c)
	/*
	 * a--b--c
	 */
	paths = dfs.PathsBetween(a, c)

	if len(paths) != 1 {
		t.Fatal("there should be a singe path between a -- c, actual path count: ", len(paths))
	}
	path := paths[0]
	if path.Length() != 3 {
		t.Error("there should be three nodes on the path: a-b-c")
	}

	dfs.Clear()
	g.AddEdge(b, d)
	g.AddEdge(c, d)
	/*
	 * a--b--c--d
	 *    \____/
	 */
	paths = dfs.PathsBetween(a, d)

	if len(paths) != 2 {
		t.Error("there should be two paths between a -- d, actual path cound: ", len(paths))
	}

	dfs.Clear()
	g.AddEdge(b, e)
	g.AddEdge(e, d)
	/*
	 *     ___e_
	 *    |     \
	 * a--b--c--d
	 *    \____/
	 */
	paths = dfs.PathsBetween(a, d)

	if len(paths) != 3 {
		t.Error("there should be three paths between a--d, actual path count: ", len(paths))
	}
}
