package trains

import (
	"github.com/kashing1999/bigpay-trains/internal/data/node"
)

type Status int

const (
	Stationary Status = iota // train starts stationary
	Moving
)

// Train is a train in the train network
type Train struct {
	Name     string
	Capacity int
	Location node.Node

	Status            Status
	DistanceRemaining int
	Parcels           []Parcel
	Destination       node.Node
}

// Move moves the train towards its destination
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

// StartJourney starts the journey of a train towards its destination
func (t *Train) StartJourney(dest node.Node, distance int) {
	t.Destination = dest
	t.DistanceRemaining = distance

	if distance <= 0 {
		t.Status = Stationary
	} else {
		t.Status = Moving
	}
}

// IsStationary returns whether the train is stationary or not
func (t *Train) IsStationary() bool {
	return t.Status == Stationary
}

// RemainingCapacity returns the remaining capacity of a train, taking into account the current load of the train
func (t *Train) RemainingCapacity() int {
	currentLoad := 0
	for _, p := range t.Parcels {
		currentLoad += p.Weight
	}
	return t.Capacity - currentLoad
}

// Offload tries to offload a set of parcels at a location
// It returns the parcels that were offloaded
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

// CanPickup returns true if the train can pick up any parcel within the station
func (t *Train) CanPickup(station Station) bool {
	remainingCapacity := t.RemainingCapacity()
	for _, parcel := range station.Parcels {
		// can only pick up parcels that are not in the correct station
		// and parcels that is under the train's current capacity
		if !parcel.IsDestination(station.Location) && parcel.Weight <= remainingCapacity {
			return true
		}
	}
	return false
}

// Pickup will pick up a set of parcels for a train to deliver
// It returns a list of packages that were picked up, and the remaining packages
func (t *Train) Pickup(parcels []Parcel) ([]Parcel, []Parcel) {
	pickedUpPackages := make([]Parcel, 0)
	remainingPackages := make([]Parcel, 0)
	remainingCapacity := t.RemainingCapacity()
	// simple algorithm to loop through the parcels, and check which one can be picked up
	// this is a knapsack problem, and can be improved with a heuristic
	for _, p := range parcels {
		if !p.IsDestination(t.Location) && p.Weight <= remainingCapacity {
			t.Parcels = append(t.Parcels, p)
			remainingCapacity -= p.Weight
			pickedUpPackages = append(pickedUpPackages, p)
		} else {
			remainingPackages = append(remainingPackages, p)
		}
	}
	return pickedUpPackages, remainingPackages
}

// AllStationary returns true if the input trains are all stationary
func AllStationary(trains []Train) bool {
	for _, t := range trains {
		if !t.IsStationary() {
			return false
		}
	}
	return true
}
