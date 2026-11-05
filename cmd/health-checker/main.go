package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	var httpEndpoints Endpoints
	flag.Var(&httpEndpoints, "endpoints", "HTTP endpoint/s to check (repeatable)")
	timeout := flag.Duration("timeout", 5*time.Second, "Per request timeout")
	mockServer := flag.Bool("mock-server", false, "Start and check a local mock health server")
	flag.Parse()

	if *mockServer && len(httpEndpoints) > 0 {
		fmt.Fprintln(os.Stderr, "-mock-server cannot be combined with -endpoints")
		os.Exit(1)
	}

	if *timeout <= 0 {
		fmt.Fprintln(os.Stderr, "-timeout must be greater than zero")
		os.Exit(1)
	}

	if *mockServer {
		endpoint, shutdown, err := startMockServer()
		if err != nil {
			fmt.Fprintf(os.Stderr, "starting mock server: %v\n", err)
			os.Exit(1)
		}
		defer shutdown()
		httpEndpoints = *NewEndpoints(endpoint)
	} else if len(httpEndpoints) == 0 {
		fmt.Fprintln(os.Stderr, "provide at least one -endpoints value or use -mock-server")
		os.Exit(1)
	}

	for _, endpoint := range httpEndpoints {
		fmt.Printf("checking %s/health\n", endpoint.String())
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	client := NewClient(httpEndpoints, *timeout)
	if err := client.CheckEndpoints(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "health checks failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("all health checks passed")
	os.Exit(0)
}
