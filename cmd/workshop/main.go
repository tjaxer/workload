package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/ilyakaznacheev/cleanenv"
	"internal/handler"
	"log"
	"net/http"
	"workshop/internal/config"
)

func main() {
	cfg := config.Server{}

	err := cleanenv.ReadConfig("config.yml", &cfg)
	if err != nil {
		log.Fatal(err)
	}
	h := handler.NewHandler()
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/test", h.Home)
	path := cfg.Host + ":" + cfg.Port
	err = http.ListenAndServe(path, r)
	if err != nil {
		log.Fatal(err)
	}
}
