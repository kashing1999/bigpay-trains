package main

import (
	"fmt"
	"github.com/kashing1999/bigpay-trains/internal/data/edge"
	"github.com/kashing1999/bigpay-trains/internal/data/node"
	"github.com/kashing1999/bigpay-trains/internal/logic/trains"
)

func main() {
	stations := []trains.Station{
		{
			Location: node.Node{Key: "A"},
			Parcels: []trains.Parcel{
				{
					Name:        "K1",
					Weight:      3,
					Destination: node.Node{Key: "C"},
				},
				{
					Name:        "K3",
					Weight:      3,
					Destination: node.Node{Key: "C"},
				},
			},
		},
		{
			Location: node.Node{Key: "B"},
			Parcels: []trains.Parcel{
				{
					Name:        "K2",
					Weight:      5,
					Destination: node.Node{Key: "A"},
				},
			},
		},
		{
			Location: node.Node{Key: "C"},
			Parcels:  []trains.Parcel{},
		},
	}

	edges := []edge.Edge{
		{
			Source: node.Node{Key: "A"},
			Dest:   node.Node{Key: "B"},
			Cost:   30,
		},
		{
			Source: node.Node{Key: "B"},
			Dest:   node.Node{Key: "C"},
			Cost:   10,
		},
	}

	ts := []trains.Train{
		{
			Name:     "Q1",
			Capacity: 6,
			Location: node.Node{Key: "B"},
		},
		{
			Name:     "Q2",
			Capacity: 3,
			Location: node.Node{Key: "A"},
		},
	}

	out, _ := trains.Deliver(stations, edges, ts)
	for _, o := range out {
		fmt.Printf("W=%d, T=%s, N1=%s, P1=%s, N2=%s, P2=%s\n", o.W, o.T.Name, o.N1.Key, o.P1, o.N2.Key, o.P2)
	}
}
