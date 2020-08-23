package jokes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"workshop/internal/api"
)

const getJokePath = "/api?format=json"

type JokeClient struct {
	url string
}

// NewJokeClient creates new client
func NewJokeClient(baseUrl string) *JokeClient {
	var client = JokeClient{url: baseUrl}
	return &client
}

func (jc *JokeClient) GetJoke() (*api.JokeResponce, error) {
	urlPath := jc.url + getJokePath
	resp, err := http.Get(urlPath)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(
			"Api request status: %s",
			http.StatusText(resp.StatusCode))
	}
	var data api.JokeResponce

	err = json.NewDecoder(resp.Body).Decode(&data)

	if err != nil {
		return nil, err
	}
	return &data, nil

}
