module workshop

go 1.13

require github.com/go-chi/chi v4.1.2+incompatible // indirect

require internal/handler v1.0.0

replace internal/handler => ./internal/handler
