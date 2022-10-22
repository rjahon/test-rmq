package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/rjahon/labs-rmq/storage/api"
	"github.com/rjahon/labs-rmq/storage/config"

	"github.com/rjahon/labs-rmq/storage/storage"
)

func main() {
	cfg := config.Load()

	// Postgres connection
	psqlUrl := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDatabase,
	)

	psqlConn, err := sqlx.Connect("postgres", psqlUrl)
	if err != nil {
		log.Fatal("error while connecting to postgres", err)
	}
	defer psqlConn.Close()

	strg := storage.NewStoragePg(psqlConn)

	router := api.New(&api.RouterOptions{
		Cfg:     &cfg,
		Storage: strg,
	})

	apiServer := &http.Server{
		Addr:    cfg.HttpPort,
		Handler: router,
	}

	log.Printf("HTTP Port: %s", apiServer.Addr)
	go func() {
		if err := apiServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal("could not start api server", err)
		}
	}()

	shutdownChan := make(chan os.Signal, 1)
	defer close(shutdownChan)
	signal.Notify(shutdownChan, syscall.SIGTERM, syscall.SIGINT)
	sig := <-shutdownChan

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), time.Second*10)
	defer shutdownCancel()

	log.Println("received os signal", sig)
	if err := apiServer.Shutdown(shutdownCtx); err != nil {
		log.Fatal("could not shutdown http server", err)
	}

	log.Println("server shutdown successfully")
}
