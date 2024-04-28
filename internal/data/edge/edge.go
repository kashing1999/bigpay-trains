package edge

import "github.com/kashing1999/bigpay-trains/internal/data/node"

// Edge is a simple data structure describing an edge, which consists of a source, a destination and it's cost
type Edge struct {
	Source node.Node
	Dest   node.Node
	Cost   int
}
