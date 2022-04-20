package main

import (
	"context"
	"net"
	"net/url"
)

// Slowloris performs single threaded slow loris attack. If you want to run distributed
// attack, just run multiple calls of the function over the same URL.
func Slowloris(ctx context.Context, host string) error {
	_, err := net.Dial("tcp", host)
	return err
}

// SanitizeDomain extracts the host from the URL
func SanitizeDomain(rawURL string) (string, error) {
	parsed, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}

	return parsed.Host, nil
}
