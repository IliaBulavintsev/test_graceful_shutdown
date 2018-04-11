package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

func main() {

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())
	var g errgroup.Group
	g.Go(func() error {
		return Listen(g, ctx, 1)
	})
	g.Go(func() error {
		return Listen(g, ctx, 2)
	})
	select {
	case sig := <-signalChannel:
		fmt.Printf("Received signal: %s\n", sig)
		cancel()
		break
	}
	// wait for all errgroup goroutines
	if err := g.Wait(); err == nil {
		fmt.Println("finished clean")
	} else {
		fmt.Printf("received error: %v", err)
	}

}
