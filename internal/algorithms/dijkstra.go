package algorithms

import (
	"math"
	"sort"

	"github.com/kashing1999/bigpay-trains/internal/data/graph"
	"github.com/kashing1999/bigpay-trains/internal/data/node"
)

type Costs map[node.Node]int
type Prev map[node.Node]*node.Node

// Dijkstra is an implementation of the Dijkstra algorithm
// It returns the costs to travel to every node in the graph
// And a previous map which allows the caller to construct a path from the node to a destination node
func Dijkstra(graph graph.SimpleGraph, start node.Node) (Costs, Prev) {
	costs := make(map[node.Node]int)
	prev := make(map[node.Node]*node.Node)
	for _, n := range graph.Nodes() {
		costs[n] = math.MaxInt32
		prev[n] = nil
	}
	costs[start] = 0

	var nodes []*node.Node
	for _, n := range graph.Nodes() {
		nodes = append(nodes, &n)
	}

	for len(nodes) != 0 {
		sort.SliceStable(nodes, func(i, j int) bool {
			return costs[*nodes[i]] < costs[*nodes[j]]
		})

		n := *nodes[0]
		nodes = nodes[1:]

		for _, edge := range graph.Neighbours(n) {
			alt := costs[n] + edge.Cost
			if alt < costs[edge.Dest] {
				costs[edge.Dest] = alt
				prev[edge.Dest] = &n
			}
		}
	}

	return costs, prev
}
