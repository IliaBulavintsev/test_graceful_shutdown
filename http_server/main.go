package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	TIMEOUT = 10 * time.Second
)

func main() {

	hs, logger := setup()

	go func() {
		logger.Printf("Listening on http://0.0.0.0%s\n", hs.Addr)
		hs.ListenAndServe()
	}()

	graceful(hs, logger)
}

func setup() (*http.Server, *log.Logger) {
	addr := ":" + os.Getenv("PORT")
	if addr == ":" {
		addr = ":2017"
	}

	logger := log.New(os.Stdout, "", 0)

	return &http.Server{
		Addr:           ":4430",
		Handler:        NewServer(Logger(logger)),
		ReadTimeout:    TIMEOUT,
		WriteTimeout:   TIMEOUT,
		MaxHeaderBytes: 1 << 20,
	}, logger
}

func graceful(hs *http.Server, logger *log.Logger) {
	stop := make(chan os.Signal, 1)

	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(TIMEOUT))
	defer cancel()

	logger.Printf("\nShutdown with timeout: %s\n", TIMEOUT)

	if err := hs.Shutdown(ctx); err != nil {
		logger.Printf("Error: %v\n", err)
	} else {
		logger.Println("Server stopped")
	}

}
