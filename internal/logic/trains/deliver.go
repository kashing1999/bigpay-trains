package trains

import (
	"github.com/kashing1999/bigpay-trains/internal/data/edge"
	"github.com/kashing1999/bigpay-trains/internal/data/node"
	"github.com/kashing1999/bigpay-trains/internal/logic/network"
)

type Result struct {
	W  int
	T  Train
	N1 node.Node
	P1 Parcels
	N2 node.Node
	P2 Parcels
}

// Deliver is a simple algorithm that reports the actions a set of trains must take to delivery all parcels to their appropriate location
// It uses a greedy algorithm, and each train is not aware of the other train
func Deliver(stations []Station, edges []edge.Edge, trains []Train) ([]Result, int) {
	// stationNodes is used to build the train network
	stationNodes := make([]node.Node, 0)
	// stationMap is used to quickly index available stations
	stationMap := make(map[node.Node]*Station)
	for _, station := range stations {
		stationMap[station.Location] = &station
		stationNodes = append(stationNodes, station.Location)
	}

	trainNetwork := network.Network{}
	trainNetwork.Init(stationNodes, edges)

	results := make([]Result, 0)

	W := -1 // -1 signifies that the delivery hasn't started
	for {
		if W >= 0 && AllStationary(trains) {
			break
		}
		W++

		// use a for loop here to signify time, since multiple trains can run concurrently at the same time
		for i := 0; i < len(trains); i++ {
			// take a pointer to modify the train in place, otherwise range will create a copy of Train
			train := &trains[i]

			train.Move()

			// train still moving to destination
			if !train.IsStationary() {
				continue
			}

			station := stationMap[train.Location]

			// offload parcels to this station
			offloaded := train.Offload(station.Location)

			// pickup parcels from this station
			pickedUp, remaining := train.Pickup(station.Parcels)
			station.Parcels = remaining

			// not needed by requirements, just simulating parcels movement to a station
			station.Parcels = append(station.Parcels, offloaded...)

			// decide where to go next
			if len(train.Parcels) > 0 { // train already has parcels
				nextParcel := train.Parcels[0]
				for _, parcel := range train.Parcels[1:] {
					// greedily chooses a parcel whose destination is the nearest to the train's current location
					if trainNetwork.Cost(train.Location, nextParcel.Destination) > trainNetwork.Cost(train.Location, parcel.Destination) {
						nextParcel = parcel
					}
				}
				// todo: add error checking here
				// if nextStation is nil, that means the graph does not have a path for the train to traverse to the desired package location
				if nextStation := trainNetwork.Next(train.Location, nextParcel.Destination); nextStation != nil {
					train.StartJourney(*nextStation, trainNetwork.Cost(train.Location, *nextStation))
				}
			} else { // train has no parcels, travel to a station to pick some up
				var nextStation *Station
				for _, station := range stationMap {
					if station.Location == train.Location {
						continue
					}
					if train.CanPickup(*station) {
						if nextStation != nil {
							if trainNetwork.Cost(train.Location, nextStation.Location) > trainNetwork.Cost(train.Location, station.Location) {
								// greedily chooses nearest station
								nextStation = station
							}
						} else {
							nextStation = station
						}
					}
				}
				// if nextStation == nil, the train has no more actions that it can take, so it stays stationary
				if nextStation != nil {
					if nextStation := trainNetwork.Next(train.Location, nextStation.Location); nextStation != nil {
						train.StartJourney(*nextStation, trainNetwork.Cost(train.Location, *nextStation))
					}
				}
			}

			// only add result if the train performed an action this loop
			if !train.IsStationary() || len(offloaded) > 0 {
				result := Result{
					W:  W,
					T:  *train,
					N1: train.Location,
					P1: pickedUp,
					N2: train.Destination,
					P2: offloaded,
				}
				results = append(results, result)
			}
		}
	}
	return results, W
}
