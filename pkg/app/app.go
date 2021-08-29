package app

import (
	"context"
	"go-twitter-downloader/pkg/config"
	"go-twitter-downloader/pkg/delivery/http"
	"go-twitter-downloader/pkg/server"
	"go-twitter-downloader/pkg/service"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run(configPath string) {
	cfg, err := config.Init(configPath)
	if err != nil {
		log.Fatal(err)
		return
	}

	services := service.NewServices()
	handlers := http.NewHandler(services)

	server := server.NewServer(cfg, handlers.Init(cfg.HTTP.Host, cfg.HTTP.Port))
	go func() {
		if err := server.Run(); err != nil {
			log.Printf("Http server: %s\n", err.Error())
		}
	}()

	log.Print("Server started")

	// Graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)

	// Block until a signal is received.
	sig := <-c
	log.Println("Got signal:", sig)
	log.Print("Shutting down server")

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	server.Stop(ctx)
}
