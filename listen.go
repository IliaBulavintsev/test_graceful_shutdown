package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func Listen(g errgroup.Group, ctx context.Context, sec int) error {
	ticker := time.NewTicker(time.Duration(sec) * time.Second)
	g.Go(func() error {
		return Listen2(ctx, sec)
	})
	for {
		select {
		case <-ticker.C:
			fmt.Printf("ticker  %ds ticked\n", sec)
		case <-ctx.Done():
			fmt.Printf("closing ticker %ds goroutine\n", sec)
			return ctx.Err()
		}
	}
}

func Listen2(ctx context.Context, sec int) error {
	ticker := time.NewTicker(time.Duration(sec) * time.Second)
	for {
		select {
		case <-ticker.C:
			fmt.Printf("ticker  2-%ds ticked\n", sec)
			time.Sleep(10 * time.Second)
		case <-ctx.Done():
			fmt.Printf("closing ticker  2-%ds goroutine\n", sec)
			return ctx.Err()
		}
	}
}
