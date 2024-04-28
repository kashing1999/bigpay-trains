package network

import (
	"github.com/kashing1999/bigpay-trains/internal/algorithms"
	"github.com/kashing1999/bigpay-trains/internal/data/edge"
	"github.com/kashing1999/bigpay-trains/internal/data/graph"
	"github.com/kashing1999/bigpay-trains/internal/data/node"
)

// Network is a simple network that allows the users to find the cost of traversing to any node in the network
// It also allows the user to find the next node to traverse to, to reach the shortest path
type Network struct {
	costMap map[node.Node]algorithms.Costs
	prevMap map[node.Node]algorithms.Prev
}

// Cost returns the cost between source and destination node
func (d *Network) Cost(source node.Node, destination node.Node) int {
	return d.costMap[source][destination]
}

// Next returns the next node in line to reach destination
// e.g. A -> B -> C
// If source is A and destination is C, the result will be B
// returns nil if there is path to destination from source
func (d *Network) Next(source node.Node, destination node.Node) *node.Node {
	current := destination
	for {
		prev := d.prevMap[source][current]
		if prev == nil {
			return nil
		}
		if *prev == source {
			return &current
		}
		current = *prev
	}
}

// Init initialises the network based on a set of nodes and edges
func (d *Network) Init(nodes []node.Node, edges []edge.Edge) {
	d.costMap = make(map[node.Node]algorithms.Costs)
	d.prevMap = make(map[node.Node]algorithms.Prev)

	g := graph.New()

	for _, n := range nodes {
		g.AddNode(n)
	}

	for _, e := range edges {
		g.AddTwoWayEdge(e.Source, e.Dest, e.Cost)
	}

	for _, n := range g.Nodes() {
		cost, prev := algorithms.Dijkstra(g, n)

		d.costMap[n] = cost
		d.prevMap[n] = prev
	}
}
