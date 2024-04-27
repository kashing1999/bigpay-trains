package algorithms

import (
	"math"
	"sort"

	"github.com/kashing1999/bigpay-trains/internal/data"
)

func Djikstra(graph data.SimpleGraph, start data.Node) map[data.Node]int {
	distances := make(map[data.Node]int)
	for _, node := range graph.Nodes() {
		distances[node] = math.MaxInt32
	}
	distances[start] = 0

	var nodes []*data.Node
	for _, node := range graph.Nodes() {
		nodes = append(nodes, &node)
	}

	for len(nodes) != 0 {
		sort.SliceStable(nodes, func(i, j int) bool {
			return distances[*nodes[i]] < distances[*nodes[j]]
		})

		node := *nodes[0]
		nodes = nodes[1:]

		for _, edge := range graph.Neighbours(node) {
			alt := distances[node] + edge.Cost
			if alt < distances[edge.Dest] {
				distances[edge.Dest] = alt
			}
		}
	}

	return distances
}
