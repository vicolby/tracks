package main

import (
	"embed"
	"log"
	"net/http"
	"tracks/handler"
	"tracks/pkg/sb"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

//go:embed public
var FS embed.FS

func main() {
	if err := InitEverything(); err != nil {
		log.Fatal(err)

	}
	router := chi.NewMux()

	router.Handle("/*", http.StripPrefix("/", http.FileServer(http.FS(FS))))
	router.Get("/", handler.MakeHandler(handler.HandleHomeIndex))
	router.Get("/login", handler.MakeHandler(handler.HandleLoginIndex))
	router.Post("/login", handler.MakeHandler(handler.HandleLogin))

	log.Fatal(http.ListenAndServe("localhost:3000", router))
}

func InitEverything() error {
	if err := godotenv.Load(); err != nil {
		return err
	}
	return sb.Init()
}
