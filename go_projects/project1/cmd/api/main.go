package main

import (
	"context"
	"fmt"
	"go_project1/internal/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Initialize handler
	apiHandler := handlers.NewAPIHandler()

	// Create new serve mux
	mux := http.NewServeMux()

	// Register routes
	mux.HandleFunc("/hello", apiHandler.Hello)
	mux.HandleFunc("/headers", apiHandler.Headers)
	mux.HandleFunc("GET /greet", apiHandler.HandleGet)
	mux.HandleFunc("POST /data", apiHandler.HandlePost)
	mux.HandleFunc("PUT /update", apiHandler.HandlePut)
	mux.HandleFunc("DELETE /delete", apiHandler.HandleDelete)
	mux.HandleFunc("PATCH /modify", apiHandler.HandlePatch)
	mux.HandleFunc("OPTIONS /options", apiHandler.HandleOptions)
	mux.HandleFunc("HEAD /head", apiHandler.HandleHead)
	mux.HandleFunc("TRACE /trace", apiHandler.HandleTrace)

	// Create server with configuration
	server := &http.Server{
		Addr:         ":8090",
		Handler:      mux,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Start server in a goroutine
	go func() {
		fmt.Printf("Server starting on port %s\n", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed to start: %v\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// Create a deadline for graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	fmt.Println("\nShutting down server...") 
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v\n", err)
	}

	fmt.Println("Server exited properly")
}