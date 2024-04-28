package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func handleError(err error, msg string) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %s\n", msg, err.Error())
		os.Exit(1)
	}
	return
}

func trimSpaces(s []string) []string {
	for i := 0; i < len(s); i++ {
		s[i] = strings.TrimSpace(s[i])
	}
	return s
}

func processEmptyLine(reader *bufio.Reader) {
	if reader == nil {
		handleError(errors.New("reader is nil"), "error while parsing")
	}
	line, err := reader.ReadString('\n')
	handleError(err, "error while parsing")

	if strings.TrimSpace(line) != "" {
		handleError(fmt.Errorf("expected new line, got %s", strings.TrimSpace(line)), "error while parsing")
	}
}
