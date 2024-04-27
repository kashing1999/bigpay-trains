package main

import (
	"fmt"

	"github.com/kashing1999/bigpay-trains/internal/algorithms"
	"github.com/kashing1999/bigpay-trains/internal/data"
)

func main() {
	fmt.Println("Hello bigpay")
	graph := data.NewSimpleGraph()

	graph.AddNode(data.Node{Key: "A"})
	graph.AddNode(data.Node{Key: "B"})
	graph.AddNode(data.Node{Key: "C"})
	graph.AddNode(data.Node{Key: "D"})

	graph.AddEdge(data.Edge{
		Source: data.Node{Key: "A"},
		Dest:   data.Node{Key: "B"},
		Cost:   30,
	})

	graph.AddEdge(data.Edge{
		Source: data.Node{Key: "B"},
		Dest:   data.Node{Key: "C"},
		Cost:   10,
	})

	graph.AddEdge(data.Edge{
		Source: data.Node{Key: "A"},
		Dest:   data.Node{Key: "C"},
		Cost:   5,
	})

	graph.AddEdge(data.Edge{
		Source: data.Node{Key: "A"},
		Dest:   data.Node{Key: "D"},
		Cost:   100,
	})

	result := algorithms.Djikstra(graph, data.Node{Key: "A"})
	fmt.Printf("%+v", result)
}
