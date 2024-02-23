package main

import (
	"embed"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/596050/dreampicai/handler"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

//go:embed public
var FS embed.FS

func main() {
	if err := initService(); err != nil {
		log.Fatalf("failed to initialize service: %v", err)
	}
	router := chi.NewMux()

	router.Handle("/*", http.StripPrefix("/", http.FileServer(http.FS(FS))))

	router.Get("/", handler.MakeHandler(handler.HandleHomeIndex))

	port := os.Getenv("HTTP_LISTEN_ADDR")
	protocol := os.Getenv("PROTOCOL")
	domain := os.Getenv("DOMAIN")
	slog.Info("Starting server at " + protocol + "://" + domain + port)
	log.Fatal(http.ListenAndServe(port, router))
}

func initService() error {
	if err := godotenv.Load(); err != nil {
		return err
	}
	os.Setenv("HTTP_LISTEN_ADDR", ":8080")
	os.Setenv("LOG_LEVEL", "debug")
	os.Setenv("PROTOCOL", "http")
	os.Setenv("DOMAIN", "localhost")
	return nil
}
