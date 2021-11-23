package main

import (
	"context"
	"os"
	"syscall"

	"github.com/fizzbuzz-endpoint/config"
	"github.com/fizzbuzz-endpoint/internal/service"

	"github.com/gol4ng/signal"
)

func main() {
	cfg := config.NewConfig(context.Background())
	container := service.NewContainer(cfg)

	server := container.GetHTTPServer()
	go server.ListenAndServe()

	idleConnsClosed := make(chan struct{})
	defer signal.Subscribe(func(signal os.Signal) {
		println(signal.String(), "signal received. stopping...")
		if err := server.Shutdown(); err != nil {
			println(err)
		}
		close(idleConnsClosed)
	}, os.Interrupt, os.Kill, syscall.SIGTERM)()

	<-idleConnsClosed
}
