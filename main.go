package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/SurkovIlya/statistics-app/internal/orders"
	"github.com/SurkovIlya/statistics-app/internal/server"
	"github.com/SurkovIlya/statistics-app/internal/storage/pg"
	postg "github.com/SurkovIlya/statistics-app/pkg/postgres"
)

func main() {
	pgParams := postg.DBParams{
		Host:     "db",
		Port:     "5432",
		Username: os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Database: os.Getenv("POSTGRES_DB"),
	}
	storage, err := postg.Initialize(pgParams)
	if err != nil {
		panic(err)
	}

	pgq := pg.New(storage)

	om := orders.New(pgq)

	srv := server.New("8080", om)

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
