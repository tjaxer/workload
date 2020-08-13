package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"internal/handler"
	"log"
	"net/http"
)



func main() {
	h := handler.NewHandler()
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/test", h.Home)
	err:=http.ListenAndServe(":8080", r)
	if err!=nil {
		log.Fatal(err)
	}
}