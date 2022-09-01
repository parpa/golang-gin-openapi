package main

import (
	"context"
	"log"
	"net/http"
	"time"

	infra "github.com/parpa/golang-gin-openapi/server/infrastructure"
	sw "github.com/parpa/golang-gin-openapi/server/service/router"
)

func main() {
	// init infrastructure
	infra.Init()

	serverPort := infra.Getenv("SERVER_PORT", "8080")
	router := sw.NewRouter()
	srv := &http.Server{
		Addr:    ":" + serverPort,
		Handler: router,
	}

	go func() {
		// service connections
		log.Printf("[server] Server started on port %s", serverPort)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("[server] listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	<-infra.Quit(func() {
		defer func() {
			log.Println("[server] Shutdown Server ...")

			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			if err := srv.Shutdown(ctx); err != nil {
				log.Fatal("Server Shutdown:", err)
			}
		}()
	})
	log.Println("[server] Server exiting")
}
