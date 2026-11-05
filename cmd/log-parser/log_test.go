package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"strings"
	"testing"
	"time"
)

func TestLogRecord_Validate(t *testing.T) {
	tests := []struct {
		name    string
		record  LogRecord
		wantErr bool
	}{
		{
			name:    "valid record",
			record:  LogRecord{Timestamp: NewUTCTimestamp(time.Date(2026, time.October, 27, 9, 0, 1, 0, time.UTC)), Level: LogInfo, Message: "mock message"},
			wantErr: false,
		},
		{
			name:    "invalid record; unsupported log level",
			record:  LogRecord{Timestamp: NewUTCTimestamp(time.Date(2026, time.October, 27, 9, 0, 1, 0, time.UTC)), Level: LogLevel("TRACE"), Message: "mock message"},
			wantErr: true,
		},
		{
			name:    "invalid record; empty message",
			record:  LogRecord{Timestamp: NewUTCTimestamp(time.Date(2026, time.October, 27, 9, 0, 1, 0, time.UTC)), Level: LogInfo, Message: ""},
			wantErr: true,
		},
		{
			name:    "invalid record; zeroed timestamp",
			record:  LogRecord{Timestamp: NewUTCTimestamp(time.Time{}), Level: LogInfo, Message: "mock message"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.record.Validate()
			if (err != nil) != tt.wantErr {
				t.Fatalf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUTCTimestamp_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr bool
	}{
		{name: "valid", input: []byte(`"2026-10-27T09:00:01Z"`), wantErr: false},
		{name: "invalid; missing local component", input: []byte(`"2026-10-27T09:00:01"`), wantErr: true},
		{name: "invalid; not a timestamp", input: []byte(`"not-a-timestamp"`), wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expected := &UTCTimestamp{}
			err := json.Unmarshal(tt.input, expected)
			if (err != nil) != tt.wantErr {
				t.Fatalf("Unmarshal() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLogLevel_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr bool
	}{
		{name: "valid debug log level", input: []byte(`"DEBUG"`), wantErr: false},
		{name: "valid info log level", input: []byte(`"INFO"`), wantErr: false},
		{name: "valid warn log level", input: []byte(`"WARN"`), wantErr: false},
		{name: "valid error log level", input: []byte(`"ERROR"`), wantErr: false},
		{name: "invalid empty input", input: []byte(`""`), wantErr: true},
		{name: "invalid nil input", input: []byte(nil), wantErr: true},
		{name: "invalid bad input", input: []byte(`"this is not a log level"`), wantErr: true},
		{name: "invalid mixed casing", input: []byte(`"deBuG"`), wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var expected LogLevel
			err := json.Unmarshal([]byte(tt.input), &expected)

			if (err != nil) != tt.wantErr {
				t.Fatalf("Unmarshal() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestProcessInput(t *testing.T) {
	tests := []struct {
		name                 string
		input                string
		wantRecords          int
		wantMalformedRecords int
	}{
		{
			name:                 "valid",
			input:                `{"timestamp":"2026-10-27T09:00:00Z","level":"INFO","message":"health check passed"}`,
			wantRecords:          1,
			wantMalformedRecords: 0,
		},
		{
			name:                 "invalid: bad json object",
			input:                `"{"timestamp":"2026-10-27T09:00:00Z","level":"INFO","message":"health check passed"}"`,
			wantRecords:          0,
			wantMalformedRecords: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := &Summary{}
			err := processInput(strings.NewReader(tt.input), actual)
			if err != nil {
				t.Fatalf("processInput() error = %v", err)
			}

			if actual.totalRecords != tt.wantRecords {
				t.Errorf("processInput() totalRecords = %d, want %d", actual.totalRecords, tt.wantRecords)
			}
			if actual.totalMalformedRecords != tt.wantMalformedRecords {
				t.Errorf("processInput() totalMalformedRecords = %d, want %d", actual.totalMalformedRecords, tt.wantMalformedRecords)
			}
		})
	}
}

type errorReader struct{}

func (errorReader) Read([]byte) (int, error) {
	return 0, errors.New("read failed")
}

func TestProcessInput_ReadError(t *testing.T) {
	summary := &Summary{}

	err := processInput(errorReader{}, summary)

	if err == nil {
		t.Fatal("processInput() error = nil, want read error")
	}
}

func TestSummary_PrintReport(t *testing.T) {
	tests := []struct {
		name       string
		summary    Summary
		wantStdout string
		wantStderr string
		wantErr    bool
	}{
		{
			name:       "valid",
			summary:    Summary{totalRecords: 10, totalMalformedRecords: 0, totalDebug: 2, totalInfo: 3, totalWarning: 4, totalError: 1},
			wantStdout: "records=10\nmalformed=0\nDEBUG=2\nINFO=3\nWARN=4\nERROR=1\n",
			wantStderr: "",
			wantErr:    false,
		},
		{
			name:       "malformed",
			summary:    Summary{totalRecords: 0, totalMalformedRecords: 1, totalDebug: 0, totalInfo: 0, totalWarning: 0, totalError: 0},
			wantStdout: "records=0\nmalformed=1\nDEBUG=0\nINFO=0\nWARN=0\nERROR=0\n",
			wantStderr: "",
			wantErr:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var mockStdout, mockStderr bytes.Buffer
			err := tt.summary.PrintReport(&mockStdout, &mockStderr)
			if (err != nil) != tt.wantErr {
				t.Errorf("PrintReport() error = %v, wantErr %v", err, tt.wantErr)
			}
			if got, want := mockStdout.String(), tt.wantStdout; got != want {
				t.Errorf("PrintReport stdout = %q, want %q", got, want)
			}
			if got, want := mockStderr.String(), tt.wantStderr; got != want {
				t.Errorf("PrintReport stderr = %q, want %q", got, want)
			}
		})
	}
}

type errorWriter struct{}

func (errorWriter) Write(p []byte) (n int, err error) {
	return 0, errors.New("write failed")
}

func TestSummary_PrintReportError(t *testing.T) {
	mockSummary := &Summary{}
	err := mockSummary.PrintReport(errorWriter{}, errorWriter{})

	if err == nil {
		t.Fatal("PrintReport() error = nil, want writer error")
	}
}
