package main

import (
	"backend/interface/handler"
	"backend/repo/db"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func main() {
	// Load port from environment variable or default to 8080
	port, exists := os.LookupEnv("PORT")
	if !exists {
		port = "8080"
	}

	// Initialize the database connection
	dbConn, err := db.ConnectToDB()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer dbConn.Close()

	// Register handlers
	handler.RegisterHandler()

	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", port),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Channel to capture OS signals for graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Printf("Listening on port: %s\n", port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed: %v", err)
		}
	}()

	// Wait for termination signal
	<-stop
	log.Println("Shutting down server...")

	// Context for graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown failed: %v", err)
	}

	log.Println("Server exited gracefully")
}
