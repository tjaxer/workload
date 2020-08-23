package api

// Client iterface with third party API
type Client interface {
	// GetJoke returns one joke
	GetJoke() (*JokeResponce, error)
}
