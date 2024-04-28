package trains

import (
	"github.com/kashing1999/bigpay-trains/internal/data/node"
)

type Station struct {
	Location node.Node
	Parcels  []Parcel
}
