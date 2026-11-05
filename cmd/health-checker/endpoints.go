package main

import (
	"fmt"
	"net/url"
	"strings"
)

type Endpoints []url.URL

func NewEndpoints(u url.URL) *Endpoints {
	return &Endpoints{u}
}

func (s Endpoints) String() string {
	values := make([]string, len(s))
	for i, u := range s {
		values[i] = u.String()
	}
	return strings.Join(values, ", ")
}

func (i *Endpoints) Set(value string) error {
	parsed, err := parse(value)
	if err != nil {
		return err
	}
	*i = append(*i, *parsed)
	return nil
}

func parse(value string) (*url.URL, error) {
	parsedURL, err := url.Parse(value)
	if err != nil {
		return nil, err
	}

	if parsedURL.Scheme == "" {
		return nil, fmt.Errorf("invalid endpoint %q: no scheme found", value)
	}
	if parsedURL.Host == "" {
		return nil, fmt.Errorf("invalid endpoint %q: no host found", value)
	}

	return parsedURL, nil
}
