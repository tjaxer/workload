package handler

import (
	"fmt"
	"log"
	"net/http"
	"workshop/internal/api"
)

type Handler struct {
	jokeClient api.Client
}

func NewHandler(jokeClient api.Client) *Handler {
	return &Handler{
		jokeClient: jokeClient,
	}
}

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	log.Printf("Fetching joke")
	joke, err := h.jokeClient.GetJoke()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		_, err = fmt.Fprint(w, joke.Joke)
		log.Println(err)
	}
}
