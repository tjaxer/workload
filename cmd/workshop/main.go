package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"net/http"
	"workshop/internal/api/jokes"
	"workshop/internal/config"
	"workshop/internal/handler"
)

func main() {
	cfg := config.Server{}

	err := cleanenv.ReadConfig("config.yml", &cfg)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Joke base url: %v\n", cfg.JokeURL)
	apiClient := jokes.NewJokeClient(cfg.JokeURL)
	h := handler.NewHandler(apiClient)
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/test", h.Home)
	path := cfg.Host + ":" + cfg.Port
	err = http.ListenAndServe(path, r)
	if err != nil {
		log.Fatal(err)
	}
}
