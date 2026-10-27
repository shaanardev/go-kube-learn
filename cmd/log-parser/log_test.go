package main

import (
	"errors"
	"testing"
)

func TestAddMalformedRecord(t *testing.T) {
	t.Run("adding fails when record already exists", func(t *testing.T) {
		summary := &Summary{
			malformedRecords: map[int]error{
				1: errors.New("mock err"),
			},
		}

		err := summary.AddMalformedRecord(1, errors.New("mock err"))
		if err == nil {
			t.Fatal("AddMalformedRecord() error = nil, want an error")
		}
	})
}
