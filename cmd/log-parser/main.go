package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

const maxCapacity = 1 << 20 // 1 MiB (1,048,576 bytes)

func main() {
	filePath := flag.String("file", "", "Path to the log stream file")
	flag.Parse()

	summary := &Summary{}
	err := withInput(*filePath, func(input io.Reader) error {
		return processInput(input, summary)
	})

	if err != nil {
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

func withInput(filePath string, process func(io.Reader) error) error {
	if filePath == "" {
		return process(os.Stdin)
	}
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("opening file failed: %w", err)
	}

	processErr := process(file)
	closeErr := file.Close()
	if closeErr != nil {
		closeErr = fmt.Errorf("closing file failed: %w", closeErr)
	}

	return errors.Join(processErr, closeErr)
}

func processInput(input io.Reader, summary *Summary) error {
	// if no filePath flag/arg read from stdin
	scanner := bufio.NewScanner(input)
	buf := make([]byte, maxCapacity)
	scanner.Split(bufio.ScanLines)
	scanner.Buffer(buf, maxCapacity)
	lineNo := 0
	for scanner.Scan() {
		lineNo++
		line := scanner.Bytes()
		logRecord := LogRecord{}
		if err := json.Unmarshal(line, &logRecord); err != nil {
			summary.AddMalformedRecord()
			fmt.Fprintf(os.Stderr, "line %d: malformed record\n", lineNo)
		} else if err := logRecord.Validate(); err != nil {
			summary.AddMalformedRecord()
			fmt.Fprintf(os.Stderr, "line %d: malformed record\n", lineNo)
		} else {
			summary.IncrementCounts(logRecord)
		}
	}

	return scanner.Err()
}
