package trains

import (
	"fmt"
	"github.com/kashing1999/bigpay-trains/internal/data/node"
	"strings"
)

// Parcel is a package that belongs to a station at any given time
// name it as parcel instead of package to prevent confusion with the "package" keyword
type Parcel struct {
	Name        string
	Weight      int
	Destination node.Node
}

type Parcels []Parcel

func (packages Parcels) String() string {
	names := make([]string, len(packages))
	for i, p := range packages {
		names[i] = p.Name
	}

	return fmt.Sprintf("[%v]", strings.Join(names, ", "))
}
