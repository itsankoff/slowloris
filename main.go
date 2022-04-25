package main

import (
	"errors"
	"flag"
	"fmt"
	"net/url"
	"os"
	"strings"
	"time"
)

func checkHelp(args []string) {
	for _, arg := range args {
		if arg == "help" || arg == "-help" || arg == "--help" {
			flag.PrintDefaults()
			os.Exit(0)
		}
	}
}

func main() {
	// cmd flags
	var rawURL *string = flag.String("url", "", "URL to perform attack")
	var count *int64 = flag.Int64("count", int64(10), "Number of parallel workers")
	var interval *time.Duration = flag.Duration("interval", 1*time.Second, "Interval for sending data")
	var timeout *time.Duration = flag.Duration("timeout", 10*time.Second, "Timeout for the whole operation")
	var userAgent *string = flag.String("user-agent", "random", "Custom User-Agent header. Default 'random' which sends random header for each worker")

	var err error
	defer func() {
		if err != nil {
			fmt.Println(err)
			flag.PrintDefaults()
		}
	}()

	flag.Parse()
	if !flag.Parsed() {
		return
	}

	checkHelp(os.Args)

	if !strings.Contains(*rawURL, "http") {
		err = errors.New("no scheme provided. use (http or https)")
		return
	}
	parsed, err := url.Parse(*rawURL)
	if err != nil {
		return
	}

	fmt.Println("slowloris...")
	fmt.Printf("creating zoo of %d loris for %s\n", *count, *rawURL)
	if err := Zoo(Options{
		URL:       parsed,
		Count:     *count,
		Interval:  *interval,
		Timeout:   *timeout,
		UserAgent: *userAgent,
	}); err != nil {
		return
	}

	return
}
