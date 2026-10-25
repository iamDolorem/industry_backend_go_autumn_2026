package main

import (
	"context"
	"errors"
	"sync"
)

type Result[R any] struct {
	Value R
	Err   error
}

var ErrInvalidWorkers = errors.New("workers must be positive")

func ParallelMap[T any, R any](ctx context.Context, workers int, in []T, fn func(context.Context, T) (R, error)) ([]Result[R], error) {
	if workers <= 0 {
		return nil, ErrInvalidWorkers
	}
	err := ctx.Err()
	if err != nil {
		return nil, err
	}
	if len(in) == 0 {
		return []Result[R]{}, nil
	}
	results := make([]Result[R], len(in))
	jobs := make(chan int)
	var wg sync.WaitGroup
	workerCount := min(workers, len(in))
	for w := 0; w < workerCount; w++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			for i := range jobs {
				if ctx.Err() != nil {
					return
				}
				value, err := fn(ctx, in[i])
				if err != nil {
					results[i].Err = err
					continue
				}
				results[i].Value = value
			}
		}()
	}
dispatch:
	for i := 0; i < len(in); i++ {
		if ctx.Err() != nil {
			break
		}

		select {
		case jobs <- i:
		case <-ctx.Done():
			break dispatch
		}
	}
	close(jobs)
	wg.Wait()
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	return results, nil
}
