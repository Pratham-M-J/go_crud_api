package main

import (
	"fmt"
	"log"
	"net/http"
	"os"           // for os.Signal
	"os/signal"    // listen for OS shutdown signals
	"syscall"      // SIGTERM/SIGINT
	"log/slog"
	"context"
	"time"         // timeout duration

	"github.com/Pratham-M-J/crud_api/internal/config"
	"github.com/Pratham-M-J/crud_api/internal/Handler/student"
)

func main() {
	fmt.Println("Welcome to crud api")

	cfg := config.MustLoad() // load YAML config into struct

	router := http.NewServeMux() // create request router

	router.HandleFunc("POST /api/students", student.New())  // response for GET /home
	

	server := http.Server{
		Addr:    cfg.Addr, // server port/address from config
		Handler: router,   // attach router to server
	}

	slog.Info("Server Started", slog.String("address",cfg.Addr))

	done := make(chan os.Signal, 1) // channel waits for shutdown signal

	signal.Notify(
		done,
		os.Interrupt,   // Ctrl+C
		syscall.SIGINT, // interrupt signal
		syscall.SIGTERM,// termination signal (docker/k8s)
	)

	go func() {
		err := server.ListenAndServe() // start server in background
		if err != nil && err != http.ErrServerClosed {
			log.Fatal("failed to start server")
		}
	}()

	<-done // block until signal received

	slog.Info("Shutting down the Server")

	ctx, cancel := context.WithTimeout( //returns 2 values: ctx → the context object, cancel → a function 
		context.Background(),
		5*time.Second, // allow 5 sec for graceful shutdown
	)
	defer cancel()

	err := server.Shutdown(ctx) // stop server gracefully
	if err != nil {
		slog.Error("failed to shutdown server")
		log.Fatal(err)
	}

	slog.Info("Server shutdown successfully")
}