package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	var domain *string = flag.String("domain", "", "Domain to perform attack")
	var count *int64 = flag.Int64("count", int64(10), "Number of parallel workers")
	var interval *time.Duration = flag.Duration("interval", 1*time.Second, "Interval for sending data")
	var timeout *time.Duration = flag.Duration("timeout", 10*time.Second, "Timeout for the whole operation")

	flag.Parse()
	if !flag.Parsed() {
		Help()
		return
	}

	if *domain == "" {
		Help()
		return
	}

	sanitized, err := SanitizeDomain(*domain)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("slowloris...")
	fmt.Printf("creating zoo of %d loris for %s\n", *count, sanitized)
	if err := Zoo(Config{
		URL:      sanitized,
		Count:    *count,
		Interval: *interval,
		Timeout:  *timeout,
	}); err != nil {
		fmt.Println(err)
	}

	return
}
