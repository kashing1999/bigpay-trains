package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/kashing1999/bigpay-trains/internal/data/edge"
	"github.com/kashing1999/bigpay-trains/internal/data/node"
	"github.com/kashing1999/bigpay-trains/internal/logic/trains"
)

type Parser struct {
	reader *bufio.Reader
	isPipe bool
}

func (p *Parser) Parse(reader *bufio.Reader) ([]trains.Station, []edge.Edge, []trains.Train) {
	if reader == nil {
		handleError(errors.New("reader is nil"), "error in parser")
	}
	p.reader = reader

	if fi, err := os.Stdin.Stat(); err == nil {
		p.isPipe = (fi.Mode() & os.ModeCharDevice) == 0
	}

	stations := p.parseStations()

	edges := p.parseEdges()

	p.parseParcels(stations)

	trainSlice := p.parseTrains()

	if !p.isPipe {
		fmt.Println()
	}

	return stations, edges, trainSlice
}

func (p *Parser) parseStations() []trains.Station {
	if !p.isPipe {
		fmt.Println("Enter number of stations")
	}

	n, err := p.reader.ReadString('\n')
	handleError(err, "error reading input")

	n = strings.TrimSpace(n)

	numberOfStations, err := strconv.ParseInt(n, 10, 32)
	handleError(err, "error parsing number of stations")

	if !p.isPipe {
		fmt.Println("Enter stations:")
	}

	stations := make([]trains.Station, numberOfStations)
	for i := 0; i < int(numberOfStations); i++ {
		station, err := p.reader.ReadString('\n')
		handleError(err, "error parsing station")

		station = strings.TrimSpace(station)

		if station == "" {
			handleError(errors.New("station name cannot be empty"), "error parsing station")
		}

		stations[i] = trains.Station{
			Location: node.Node{Key: station},
			Parcels:  make([]trains.Parcel, 0),
		}
	}
	return stations
}

func (p *Parser) parseEdges() []edge.Edge {
	if !p.isPipe {
		fmt.Println()
		fmt.Println("Enter number of edges:")
	} else {
		processEmptyLine(p.reader)
	}

	n, err := p.reader.ReadString('\n')
	handleError(err, "error reading input")

	n = strings.TrimSpace(n)

	num, err := strconv.ParseInt(n, 10, 32)
	handleError(err, "error parsing number of edges")

	if !p.isPipe {
		fmt.Println("Enter edges:")
	}

	edges := make([]edge.Edge, num)
	for i := 0; i < int(num); i++ {
		input, err := p.reader.ReadString('\n')
		handleError(err, "error parsing edges")

		split := strings.Split(input, ",")
		if len(split) != 4 {
			handleError(fmt.Errorf("expected 4 parameters for an edge, got %d", len(split)), "error parsing edges")
		}

		split = trimSpaces(split)

		distance, err := strconv.ParseInt(split[3], 10, 32)
		handleError(err, "error parsing distance for edge")

		edges[i] = edge.Edge{
			Name:   split[0],
			Source: node.Node{Key: split[1]},
			Dest:   node.Node{Key: split[2]},
			Cost:   int(distance),
		}
	}
	return edges
}

func (p *Parser) parseParcels(stations []trains.Station) {
	if !p.isPipe {
		fmt.Println()
		fmt.Println("Enter number of parcels:")
	} else {
		processEmptyLine(p.reader)
	}

	n, err := p.reader.ReadString('\n')
	handleError(err, "error reading input")

	n = strings.TrimSpace(n)

	num, err := strconv.ParseInt(n, 10, 32)
	handleError(err, "error parsing number of parcels")

	if !p.isPipe {
		fmt.Printf("Enter %d parcels:\n", num)
	}

	for i := 0; i < int(num); i++ {
		input, err := p.reader.ReadString('\n')
		handleError(err, "error parsing parcels")

		split := strings.Split(input, ",")
		if len(split) != 4 {
			handleError(fmt.Errorf("expected 4 parameters for a parcel, got %d", len(split)), "error parsing parcel")
		}

		split = trimSpaces(split)

		weight, err := strconv.ParseInt(split[1], 10, 32)
		handleError(err, "error parsing weight for parcel")

		n := node.Node{Key: split[2]}
		for i := 0; i < len(stations); i++ {
			if stations[i].Location == n {
				stations[i].Parcels = append(stations[i].Parcels, trains.Parcel{
					Name:        split[0],
					Weight:      int(weight),
					Destination: node.Node{Key: split[3]},
				})
			}
		}
	}
}

func (p *Parser) parseTrains() []trains.Train {
	if !p.isPipe {
		fmt.Println()
		fmt.Println("Enter number of trains:")
	} else {
		processEmptyLine(p.reader)
	}

	n, err := p.reader.ReadString('\n')
	handleError(err, "error reading input")

	n = strings.TrimSpace(n)

	num, err := strconv.ParseInt(n, 10, 32)
	handleError(err, "error parsing number of trains")

	if !p.isPipe {
		fmt.Println("Enter trains:")
	}

	trainSlice := make([]trains.Train, num)
	for i := 0; i < int(num); i++ {
		input, err := p.reader.ReadString('\n')
		handleError(err, "error parsing trains")

		split := strings.Split(input, ",")
		if len(split) != 3 {
			handleError(fmt.Errorf("expected 3 parameters for a train, got %d", len(split)), "error parsing edges")
		}

		split = trimSpaces(split)

		capacity, err := strconv.ParseInt(split[1], 10, 32)
		handleError(err, "error parsing capacity for train")

		trainSlice[i] = trains.Train{
			Name:     split[0],
			Capacity: int(capacity),
			Location: node.Node{Key: split[2]},
		}
	}
	return trainSlice
}
