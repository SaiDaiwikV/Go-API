package main

import (
	"context"
	// "fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/SaiDaiwikV/Go-API/internal/config"
	"github.com/SaiDaiwikV/Go-API/internal/http/handlers/student"
	"github.com/SaiDaiwikV/Go-API/internal/storage/sqlite"
)

func main() {
	// load config

	cfg := config.MustLoad()
	// database setup

	storage, err := sqlite.New(cfg)

	if err != nil {
		log.Fatal(err)
	}
	slog.Info("Storage initialized",slog.String("env",cfg.Env),slog.String("version", "1.0.0"))
	// setup router
	router := http.NewServeMux()

	router.HandleFunc("POST /api/students", student.New(storage))


	// setup server
	server := http.Server{
		Addr: cfg.Addr,
		Handler: router,
	}
	
	slog.Info("Server Started",slog.String("address",cfg.Addr))

	done := make(chan os.Signal,1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT,syscall.SIGTERM)

	go func(){
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("Failed to start server")
		}
	}()

	<- done


	slog.Info("Shutting down the server")

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	err = server.Shutdown(ctx)

	if err != nil {
		slog.Error("Failed to shutdown the server",slog.String("error", err.Error()))
	}

	slog.Info("Server shutdown successfully")

}
