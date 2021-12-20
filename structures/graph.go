package structures

import (
	"fmt"
	"sync"
)

type Graph struct {
	Nodes []Node
	Edges map[Node][]*Node
	lock  sync.Mutex
}

type Node struct {
	Value string
}

type Edge struct {
	A *Node
	B *Node
}

func (n *Node) String() string {
	return fmt.Sprintf("%v", n.Value)
}

// addNode adds a node to the graph
func (g *Graph) addNode(n Node) {
	fmt.Printf("attempting to add node %v\n", n)
	if !g.Contains(n) {
		fmt.Printf("node added!\n")
		g.Nodes = append(g.Nodes, n)
	}
}

func (g *Graph) GetNode(v string) (*Node, bool) {
	for _, node := range g.Nodes {
		if node.Value == v {
			return &node, true
		}
	}
	return &Node{}, false
}

func (g *Graph) GetAdjacentNodes(n Node) []*Node {
	return g.Edges[n]
}

func (g *Graph) Contains(n Node) bool {
	for _, node := range g.Nodes {
		if node.Value == n.Value {
			return true
		}
	}
	return false
}

// AddEdge adds an edge to the graph
func (g *Graph) AddEdge(n1, n2 Node) {
	g.lock.Lock()
	if g.Edges == nil {
		g.Edges = make(map[Node][]*Node)
	}
	n1p, ok := g.GetNode(n1.Value)
	if !ok {
		g.addNode(n1)
		n1p = &n1
	}
	n2p, ok := g.GetNode(n2.Value)
	if !ok {
		g.addNode(n2)
		n2p = &n2
	}
	g.Edges[n1] = append(g.Edges[n1], n2p)
	g.Edges[n2] = append(g.Edges[n2], n1p)
	g.lock.Unlock()
}

func (g *Graph) String() string {
	g.lock.Lock()
	s := ""
	for i := 0; i < len(g.Nodes); i++ {
		s += g.Nodes[i].String() + " -> "
		near := g.Edges[g.Nodes[i]]
		for j := 0; j < len(near); j++ {
			s += near[j].String() + " "
		}
		s += "\n"
	}
	g.lock.Unlock()
	return s
}
