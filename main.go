package main

import (
	"blog-rest/internal/handlers"
	"blog-rest/internal/storage/memory"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	store := memory.NewBlogStorage()

	userHandler := handlers.NewUserHandler(store)
	postHandler := handlers.NewPostHandler(store)

	http.HandleFunc("/users", userHandler.HandleUsers)
	http.HandleFunc("/user", userHandler.HandleUserCreate)
	http.HandleFunc("/user/:id", userHandler.HandleUserUpdate) // PATCH /user/:id
	http.HandleFunc("/user/", userHandler.HandleUserDelete)    // DELETE /user/:id
	http.HandleFunc("/posts", postHandler.HandlePostCreate)    // Пример для POST /posts

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", os.Getenv("PORT")),
		Handler: nil,
	}

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Printf("Server started on :%s", os.Getenv("PORT"))
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not listen on %s: %v\n", server.Addr, err)
		}
	}()

	<-stopChan
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exiting")
}
