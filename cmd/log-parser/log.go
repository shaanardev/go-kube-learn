package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"time"
)

type LogRecord struct {
	Timestamp time.Time `json:"timestamp"`
	Level     LogLevel  `json:"level"`
	Message   string    `json:"message"`
}

func (r LogRecord) Validate() error {
	if r.Timestamp.IsZero() {
		return errors.New("timestamp is required")
	}
	if r.Message == "" {
		return errors.New("message is required")
	}
	return nil
}

type LogLevel string

const (
	LogDebug LogLevel = "DEBUG"
	LogInfo  LogLevel = "INFO"
	LogWarn  LogLevel = "WARN"
	LogError LogLevel = "ERROR"
)

func (l *LogLevel) UnmarshalJSON(data []byte) (err error) {
	var value string
	if err = json.Unmarshal(data, &value); err != nil {
		return fmt.Errorf("log level must be a string: %w", err)
	}
	level := LogLevel(value)
	switch level {
	case LogDebug, LogInfo, LogWarn, LogError:
		*l = level
		return nil
	default:
		return fmt.Errorf("invalid log level %q", value)
	}
}

type Summary struct {
	malformedRecords map[int]error
	records          []LogRecord
	totalDebug       int
	totalInfo        int
	totalWarning     int
	totalError       int
}

func NewSummary() *Summary {
	return &Summary{
		malformedRecords: make(map[int]error),
	}
}

func (s *Summary) AddRecord(record LogRecord) {
	s.records = append(s.records, record)
	switch record.Level {
	case LogDebug:
		s.totalDebug++
	case LogInfo:
		s.totalInfo++
	case LogWarn:
		s.totalWarning++
	case LogError:
		s.totalError++
	default:
		fmt.Fprintln(os.Stderr, fmt.Errorf("unhandled log level %s", record.Level))
	}
}

func (s *Summary) AddMalformedRecord(lineNumber int, reason error) error {
	if _, exists := s.malformedRecords[lineNumber]; exists {
		return fmt.Errorf("malformed record already record for line %d", lineNumber)
	}
	s.malformedRecords[lineNumber] = reason
	return nil
}

func (s Summary) TotalMalformedRecords() int { return len(s.malformedRecords) }

func (s Summary) PrintReport(stdout, stderr io.Writer) (err error) {
	if len(s.malformedRecords) > 0 {
		for key, value := range s.malformedRecords {
			if _, err = fmt.Fprintf(stderr, "line %d: malformed record; %s\n", key, value); err != nil {
				return fmt.Errorf("PrintReport failed: %w", err)
			}
		}
	}

	if _, err = fmt.Fprintf(stdout,
		"records=%d\nmalformed=%d\nDEBUG=%d\nINFO=%d\nWARN=%d\nERROR=%d\n",
		len(s.records),
		len(s.malformedRecords),
		s.totalDebug,
		s.totalInfo,
		s.totalWarning,
		s.totalError,
	); err != nil {
		return fmt.Errorf("PrintReport failed: %w", err)
	}

	return nil
}
