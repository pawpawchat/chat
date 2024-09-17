package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/pawpawchat/chat/config"
	"github.com/pawpawchat/chat/internal/app"
)

var cfg *string

func init() {
	cfg = flag.String("cfg", "config.yaml", "path to config file")
}

func main() {
	flag.Parse()
	cfg, err := config.LoadConfig(*cfg)
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

	app.Run(ctx, cfg)
}
