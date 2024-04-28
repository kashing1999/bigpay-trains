package trains

import (
	"github.com/kashing1999/bigpay-trains/internal/data/node"
)

type Status int

const (
	Stationary Status = iota // train starts stationary
	Moving
)

type Train struct {
	Name     string
	Capacity int
	Location node.Node

	Status            Status
	DistanceRemaining int
	Parcels           []Parcel
	Destination       node.Node
}

func (t *Train) Move() {
	switch t.Status {
	case Stationary:
	case Moving:
		t.DistanceRemaining -= 1
		if t.DistanceRemaining == 0 {
			t.Status = Stationary
			t.Location = t.Destination
		}
	}
}

func (t *Train) StartJourney(dest node.Node, distance int) {
	t.Destination = dest
	t.DistanceRemaining = distance
	t.Status = Moving
}

func (t *Train) IsStationary() bool {
	return t.Status == Stationary
}

func (t *Train) RemainingCapacity() int {
	currentLoad := 0
	for _, p := range t.Parcels {
		currentLoad += p.Weight
	}
	return t.Capacity - currentLoad
}

func (t *Train) Offload(location node.Node) []Parcel {
	offload := make([]Parcel, 0)
	remaining := make([]Parcel, 0)

	if t.Status == Moving {
		return offload
	}

	for _, p := range t.Parcels {
		if p.Destination == location {
			offload = append(offload, p)
		} else {
			remaining = append(remaining, p)
		}
	}
	t.Parcels = remaining
	return offload
}

func (t *Train) CanPickup(parcels []Parcel) bool {
	remainingCapacity := t.RemainingCapacity()
	for _, parcel := range parcels {
		if parcel.Weight <= remainingCapacity {
			return true
		}
	}
	return false
}

func (t *Train) Pickup(parcels []Parcel) ([]Parcel, []Parcel) {
	pickedUpPackages := make([]Parcel, 0)
	remainingPackages := make([]Parcel, 0)
	remainingCapacity := t.RemainingCapacity()
	// simple algorithm to loop through the parcels, and check which one can be picked up
	// this is a knapsack problem, and can be improved with a heuristic
	for _, p := range parcels {
		if p.Destination != t.Location && p.Weight <= remainingCapacity {
			t.Parcels = append(t.Parcels, p)
			remainingCapacity -= p.Weight
			pickedUpPackages = append(pickedUpPackages, p)
		} else {
			remainingPackages = append(remainingPackages, p)
		}
	}
	return pickedUpPackages, remainingPackages
}

func AllStationary(trains []Train) bool {
	for _, t := range trains {
		if !t.IsStationary() {
			return false
		}
	}
	return true
}
