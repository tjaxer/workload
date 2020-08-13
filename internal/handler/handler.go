package handler

import (
	"fmt"
	"net/http"
)

type Handler struct {
	
}

func  NewHandler() *Handler  {
	return &Handler{}
}

func (*Handler) Home (w http.ResponseWriter, r *http.Request)  {
	fmt.Fprint(w, "Hello world")
}