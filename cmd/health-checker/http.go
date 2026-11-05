package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

// refs:
// https://codewithhugo.com/parallel-http-request-golang/

var ErrHealthCheckRequestFailed = errors.New("health check request failed")

type Client struct {
	endpoints  Endpoints
	httpClient http.Client
}

func NewClient(endpoints Endpoints, timeout time.Duration) *Client {
	return &Client{endpoints: endpoints, httpClient: http.Client{Timeout: timeout}}
}

func (c *Client) CheckEndpoints(ctx context.Context) error {
	errs := make(chan error, len(c.endpoints))
	var wg sync.WaitGroup

	for _, endpoint := range c.endpoints {
		wg.Add(1) // increase counter 1 per worker
		go func(u url.URL) {
			defer wg.Done()
			if err := c.healthCheckRequest(ctx, u); err != nil {
				errs <- err
			}
		}(endpoint)
	}

	// Close only after every worker has finished sending.
	go func() {
		wg.Wait()
		close(errs)
	}()

	var failures []error
	for err := range errs {
		failures = append(failures, err)
	}

	return errors.Join(failures...)
}

const maxRespBody = 4 << 10 // 4 KiB

func (c *Client) healthCheckRequest(ctx context.Context, u url.URL) (err error) {
	formattedUrl, err := url.JoinPath(u.String(), "health")
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, "GET", formattedUrl, nil)
	if err != nil {
		return err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil {
			err = errors.Join(err, fmt.Errorf("closing response body for %s: %w", u.String(), closeErr))
		}
	}()

	if resp.StatusCode < http.StatusOK || resp.StatusCode > http.StatusIMUsed {
		body, err := io.ReadAll(io.LimitReader(resp.Body, maxRespBody))
		if err == nil {
			return fmt.Errorf(
				"%w: %s returned status %d (body: %s)",
				ErrHealthCheckRequestFailed,
				u.String(),
				resp.StatusCode,
				strings.TrimSpace(string(body)),
			)
		}

		return fmt.Errorf(
			"%w: %s returned status %d",
			ErrHealthCheckRequestFailed,
			u.String(),
			resp.StatusCode,
		)
	}

	return nil
}
