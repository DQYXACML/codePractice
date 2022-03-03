package graph

type Graph struct {
	Nodes map[int]Node
	Edges []Edge
}

func NewGraph() *Graph {
	return &Graph{
		Nodes: make(map[int]Node),
		Edges: make([]Edge, 0),
	}
}