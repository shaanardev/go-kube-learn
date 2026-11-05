package main

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"
)

func endpointForTest(t *testing.T, server *httptest.Server) url.URL {
	t.Helper()

	parsed, err := url.Parse(server.URL)
	if err != nil {
		t.Fatalf("parse test server URL: %v", err)
	}
	return *parsed
}

func TestClient_CheckEndpoints_OK(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("request method = %q, want %q", r.Method, http.MethodGet)
		}
		if r.URL.Path != "/health" {
			t.Errorf("request path = %q, want %q", r.URL.Path, "/health")
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client := NewClient(Endpoints{endpointForTest(t, server)}, time.Second)
	if err := client.CheckEndpoints(context.Background()); err != nil {
		t.Fatalf("CheckEndpoints() error = %v, want nil", err)
	}
}

func TestClient_CheckEndpoints_JoinsErrors(t *testing.T) {
	serverOne := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusServiceUnavailable)
		_, _ = w.Write([]byte("service one unavailable"))
	}))
	defer serverOne.Close()

	serverTwo := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("service two failed"))
	}))
	defer serverTwo.Close()

	client := NewClient(Endpoints{
		endpointForTest(t, serverOne),
		endpointForTest(t, serverTwo),
	}, time.Second)
	err := client.CheckEndpoints(context.Background())

	if !errors.Is(err, ErrHealthCheckRequestFailed) {
		t.Fatalf("CheckEndpoints() error = %v, want ErrHealthCheckRequestFailed", err)
	}
	for _, message := range []string{"service one unavailable", "service two failed"} {
		if !strings.Contains(err.Error(), message) {
			t.Errorf("CheckEndpoints() error %q does not contain %q", err, message)
		}
	}
}

func TestClient_CheckEndpoints_UsesPerRequestTimeout(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(_ http.ResponseWriter, r *http.Request) {
		<-r.Context().Done()
	}))
	defer server.Close()

	client := NewClient(Endpoints{endpointForTest(t, server)}, 20*time.Millisecond)
	err := client.CheckEndpoints(context.Background())
	if err == nil {
		t.Fatal("CheckEndpoints() error = nil, want timeout error")
	}
}
