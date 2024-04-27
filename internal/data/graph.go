package data

type SimpleGraph struct {
	nodes map[Node][]Edge
}

func NewSimpleGraph() SimpleGraph {
	return SimpleGraph{
		nodes: make(map[Node][]Edge),
	}
}

func (g *SimpleGraph) AddNode(node Node) {
	g.nodes[node] = []Edge{}
}

func (g *SimpleGraph) AddEdge(edge Edge) {
	g.nodes[edge.Source] = append(g.nodes[edge.Source], edge)
}

func (g *SimpleGraph) Neighbours(node Node) []Edge {
	return g.nodes[node]
}

func (g *SimpleGraph) Nodes() []Node {
	nodes := make([]Node, len(g.nodes))
	i := 0
	for node := range g.nodes {
		nodes[i] = node
		i++
	}
	return nodes
}
