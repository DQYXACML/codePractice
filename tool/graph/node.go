package graph

type Node struct {
	value int
	in    int
	out   int
	nexts []Node
	edges []Edge
}
