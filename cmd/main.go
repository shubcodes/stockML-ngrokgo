package main

import (
	"context"
	"log"
	"net/http"

	"shub.codes/ngrokgo-stock/internal/handlers"

	"github.com/joho/godotenv"

	"github.com/gorilla/mux"
	"golang.ngrok.com/ngrok"
	"golang.ngrok.com/ngrok/config"
)

func main() {
	if err := run(context.Background()); err != nil {
		log.Fatal(err)
	}
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func run(ctx context.Context) error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := mux.NewRouter()

	// Define routes
	r.HandleFunc("/stock", handlers.StockHandler).Methods(http.MethodPost)
	r.HandleFunc("/news", handlers.NewsHandler).Methods(http.MethodPost)

	// Serve static files
	fs := http.FileServer(http.Dir("static"))
	r.PathPrefix("/").Handler(fs)

	srv := &http.Server{
		Addr:    ":8000",
		Handler: r,
	}

	// Set up ngrok tunnel
	tun, err := ngrok.Listen(ctx,
		config.HTTPEndpoint(),
		ngrok.WithAuthtokenFromEnv(),
	)
	if err != nil {
		return err
	}

	log.Println("tunnel created:", tun.URL())

	// Serve HTTP traffic through the ngrok tunnel
	go func() {
		log.Fatal(http.Serve(tun, srv.Handler))
	}()

	// Serve HTTP traffic directly
	log.Println("Starting server at :8000")
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

	return nil
}
