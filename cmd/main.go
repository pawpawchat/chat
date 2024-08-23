package main

import (
	"context"
	"flag"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/pawpawchat/chat/config"
	"github.com/pawpawchat/chat/internal/app"
)

func main() {
	flag.Parse()
	cfg, err := config.LoadDefaultConfig()
	if err != nil {
		log.Fatal(err)
	}

	config.ConfigureLogger(cfg)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		stop := make(chan os.Signal, 1)
		signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
		<-stop
		cancel()
	}()

	if err := app.Run(ctx, cfg); err != nil {
		slog.Error("app was terminated with an error", "err", err)
		os.Exit(1)
	}
}
