package graph

import (
	"github.com/kashing1999/bigpay-trains/internal/data/edge"
	"github.com/kashing1999/bigpay-trains/internal/data/node"
)

// SimpleGraph is a simple implementation of a directed, weighted graph
// Could use open source library, but for purposes of this assessment, will use a simple implementation
type SimpleGraph struct {
	nodes map[node.Node][]edge.Edge
}

// New returns a new simple graph
func New() SimpleGraph {
	return SimpleGraph{
		nodes: make(map[node.Node][]edge.Edge),
	}
}

// AddNode adds a node to the graph
func (g *SimpleGraph) AddNode(node node.Node) {
	g.nodes[node] = []edge.Edge{}
}

// AddEdge adds a directed edge to the graph. If the source not does not exist, it will create the node.
func (g *SimpleGraph) AddEdge(source node.Node, destination node.Node, cost int) {
	g.nodes[source] = append(g.nodes[source], edge.Edge{
		Source: source,
		Dest:   destination,
		Cost:   cost,
	})
}

// AddTwoWayEdge adds a bidirectional edge to the graph. If the source not does not exist, it will create the node.
func (g *SimpleGraph) AddTwoWayEdge(a node.Node, b node.Node, cost int) {
	g.AddEdge(a, b, cost)
	g.AddEdge(b, a, cost)
}

// Neighbours returns all edges for a given node
func (g *SimpleGraph) Neighbours(node node.Node) []edge.Edge {
	return g.nodes[node]
}

// Nodes returns a list of all the nodes in this graph
func (g *SimpleGraph) Nodes() []node.Node {
	nodes := make([]node.Node, len(g.nodes))
	i := 0
	for n := range g.nodes {
		nodes[i] = n
		i++
	}
	return nodes
}
