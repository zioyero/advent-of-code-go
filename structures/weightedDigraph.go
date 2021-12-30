package structures

type WeightedDigraph struct {
	Nodes []PositionNode
	Edges map[PositionNode][]PositionNode
}

type PositionNode struct {
	Position
	Value int
}

type WeightedEdge struct {
	A PositionNode
	B PositionNode
	V int
}

func NewWeightedDigraph() WeightedDigraph {
	return WeightedDigraph{Nodes: make([]PositionNode, 0)}
}

func AppendNode(g WeightedDigraph, n PositionNode) WeightedDigraph {
	nodes := append(g.Nodes, n)
	return WeightedDigraph{Nodes: nodes}
}

func AppendEdge(g WeightedDigraph, e WeightedEdge) WeightedDigraph {
	adjacent := append(g.Edges[e.A], e.B)
	g.Edges[e.A] = adjacent
	return WeightedDigraph{Nodes: g.Nodes, Edges: g.Edges}
}

func (g WeightedDigraph) GetAdjacentNodes(n PositionNode) []PositionNode {
	return g.Edges[n]
}

type DepthFirstSearch struct {
	graph WeightedDigraph
	explored []
}

func LowestValuePath(g WeightedDigraph, explored []PositionNode, min int) ([]PositionNode, int) {

}
