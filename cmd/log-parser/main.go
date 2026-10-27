package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	filePath := flag.String("file", "", "Path to the log stream file")
	flag.Parse()

	input, parseErr := parseInput(*filePath)
	if parseErr != nil {
		fmt.Fprintln(os.Stderr, parseErr)
		os.Exit(1)
	}

	// if no filePath flag/arg read from stdin
	scanner := bufio.NewScanner(input)
	lineNo := 0
	summary := NewSummary()
	for scanner.Scan() {
		lineNo++
		line := scanner.Bytes()
		logRecord := LogRecord{}
		if err := json.Unmarshal(line, &logRecord); err != nil {
			if addErr := summary.AddMalformedRecord(lineNo, err); addErr != nil {
				fmt.Fprintln(os.Stderr, "adding malformed record failed:", addErr)
				os.Exit(1)
			}
		} else if err := logRecord.Validate(); err != nil {
			if addErr := summary.AddMalformedRecord(lineNo, err); addErr != nil {
				fmt.Fprintln(os.Stderr, "adding malformed record failed:", addErr)
				os.Exit(1)
			}
		} else {
			summary.AddRecord(logRecord)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
		os.Exit(1)
	}

	if err := summary.PrintReport(os.Stdout, os.Stderr); err != nil {
		fmt.Fprintln(os.Stderr, "adding malformed record failed:", err)
		os.Exit(1)
	}

	if summary.TotalMalformedRecords() == 0 {
		os.Exit(0)
	}
	os.Exit(1)
}

func parseInput(filePath string) (input io.Reader, err error) {
	input = os.Stdin

	if filePath != "" {
		file, openErr := os.Open(filePath)
		if openErr != nil {
			return nil, fmt.Errorf("opening file: %w", openErr)
		}

		defer func() {
			if closeErr := file.Close(); err == nil {
				err = fmt.Errorf("closing file: %w", closeErr)
			}
		}()

		input = file
	}

	return input, nil
}
