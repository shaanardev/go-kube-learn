package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

type UTCTimestamp struct {
	time.Time
}

func NewUTCTimestamp(value time.Time) UTCTimestamp {
	return UTCTimestamp{value}
}

func (t *UTCTimestamp) UnmarshalJSON(data []byte) error {
	var raw string
	if err := json.Unmarshal(data, &raw); err != nil {
		return fmt.Errorf("timestamp must be a string: %w", err)
	}
	// RFC3339 doesnt have the Z component that the contract specifies
	parsed, err := time.Parse(time.RFC3339, raw)
	if err != nil {
		return fmt.Errorf("invalid timestamp: %w", err)
	}

	if parsed.IsZero() {
		return errors.New("invalid zeroed timestamp")
	}

	if parsed.Location() != time.UTC || parsed.Nanosecond() != 0 || !strings.HasSuffix(raw, "Z") {
		return errors.New("timestamp must be UTC with whole seconds")
	}
	t.Time = parsed
	return nil
}

type LogRecord struct {
	Timestamp UTCTimestamp `json:"timestamp"`
	Level     LogLevel     `json:"level"`
	Message   string       `json:"message"`
}

func (r LogRecord) Validate() error {
	if r.Timestamp.IsZero() {
		return errors.New("invalid timestamp; zeroed out")
	}

	if r.Message == "" {
		return errors.New("message is required")
	}

	switch r.Level {
	case LogDebug, LogInfo, LogWarn, LogError:
		return nil
	default:
		return fmt.Errorf("invalid log level %q", r.Level)
	}
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
	totalRecords          int
	totalMalformedRecords int
	totalDebug            int
	totalInfo             int
	totalWarning          int
	totalError            int
}

func (s *Summary) IncrementCounts(record LogRecord) {
	s.totalRecords++

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

func (s *Summary) AddMalformedRecord() { s.totalMalformedRecords++ }

func (s Summary) TotalMalformedRecords() int { return s.totalMalformedRecords }

func (s Summary) PrintReport(stdout, stderr io.Writer) (err error) {
	if _, err = fmt.Fprintf(stdout,
		"records=%d\nmalformed=%d\nDEBUG=%d\nINFO=%d\nWARN=%d\nERROR=%d\n",
		s.totalRecords,
		s.totalMalformedRecords,
		s.totalDebug,
		s.totalInfo,
		s.totalWarning,
		s.totalError,
	); err != nil {
		return fmt.Errorf("PrintReport failed: %w", err)
	}

	return nil
}
