package main

import (
	"context"
	"net/http"
)

// Slowloris performs single threaded slow loris attack. If you want to run distributed
// attack, just run multiple calls of the function over the same URL.
func Slowloris(ctx context.Context, url string) error {
	_, err := http.Get(url)
	return err
}
