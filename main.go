package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/SurkovIlya/statistics-app/internal/server"
)

func main() {
	srv := server.New("8080")

	go func() {
		if err := srv.Run(); err != nil {
			log.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	log.Print("app Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Print("app Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Panicf("error occured on server shutting down: %s", err.Error())
	}
}
