package trains

import (
	"github.com/kashing1999/bigpay-trains/internal/data/node"
)

// Station is a representation of a train station in the train network
type Station struct {
	Location node.Node
	Parcels  []Parcel
}
