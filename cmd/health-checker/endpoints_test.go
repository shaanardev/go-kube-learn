package main

import "testing"

func TestEndpoints_Set(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantURL string
		wantErr bool
	}{
		{name: "scheme without host", input: "https", wantErr: true},
		{name: "scheme with empty host", input: "https://", wantErr: true},
		{name: "empty string", input: "", wantErr: true},
		{name: "http with short host", input: "http://www", wantURL: "http://www", wantErr: false},
		{name: "http with domain", input: "http://www.dumpsters.com", wantURL: "http://www.dumpsters.com", wantErr: false},
		{name: "https with explicit port", input: "https://www.dumpsters.com:443", wantURL: "https://www.dumpsters.com:443", wantErr: false},
		{name: "absolute path without scheme", input: "/testing-path", wantErr: true},
		{name: "relative path without scheme", input: "testing-path", wantErr: true},
		{name: "malformed relative URL", input: "alskjff#?asf//dfas", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got Endpoints
			err := got.Set(tt.input)
			if (err != nil) != tt.wantErr {
				t.Fatalf("Set() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr && len(got) != 0 {
				t.Errorf("Set() added %d endpoints for invalid input, want 0", len(got))
			}
			if !tt.wantErr && len(got) != 1 {
				t.Fatalf("Set() added %d endpoints, want 1", len(got))
			}
			if !tt.wantErr && got[0].String() != tt.wantURL {
				t.Errorf("Set() parsed URL = %q, want %q", got[0].String(), tt.wantURL)
			}
		})
	}
}
