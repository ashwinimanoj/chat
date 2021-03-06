package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		log.Fatal("no arguments passed. Check help for more information")
	}

	cmd := args[0]
	switch cmd {
	case "help":
		log.Println("Use serve to start the application.")
	case "serve":
		serve()
	default:
		log.Fatal("Invalid command. Use help to see available options")
	}

}

func serve() {
	port := 8080
	log.Printf("Starting Server on port %d", port)
	server := http.Server{
		Addr: fmt.Sprintf(":%d", port),
	}
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("Server Error: %v", err)
	}
}
