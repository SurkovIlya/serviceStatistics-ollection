package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/SurkovIlya/statistics-app/docs"
	"github.com/SurkovIlya/statistics-app/internal/orders"
	"github.com/SurkovIlya/statistics-app/internal/server"
	"github.com/SurkovIlya/statistics-app/internal/storage/pg"
	st "github.com/SurkovIlya/statistics-app/pkg/postgres"
)

const serverPort = "8080"

// @title Statistics-App API
// @version 1.0
// @description API server statistics-App.
// @host localhost:8080
// @BasePath /
func main() {
	pgParams := st.DBParams{
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     os.Getenv("POSTGRES_PORT"),
		Username: os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Database: os.Getenv("POSTGRES_DB"),
	}

	conn, err := st.Connect(pgParams)
	if err != nil {
		panic(err)
	}

	db := st.New(conn)

	pgq := pg.New(db)

	orderManager := orders.New(pgq)

	srv := server.New(serverPort, orderManager)

	go func() {
		if err := srv.Run(); err != nil {
			log.Panicf("error occured while running http server: %s", err.Error())
		}
	}()

	log.Println("app Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Println("app Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Panicf("error occured on server shutting down: %s", err.Error())
	}
}
