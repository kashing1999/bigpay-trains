package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/kashing1999/bigpay-trains/internal/logic/trains"
)

func main() {
	parser := Parser{}
	stations, edges, trainSlice := parser.Parse(bufio.NewReader(os.Stdin))

	results, W := trains.Deliver(stations, edges, trainSlice)

	for _, result := range results {
		fmt.Printf("W=%d, T=%s, N1=%s, P1=%s, N2=%s, P2=%s\n", result.W, result.T.Name, result.N1.Key, result.P1, result.N2.Key, result.P2)
	}
	fmt.Printf("Total wait time: %d\n", W)
}
