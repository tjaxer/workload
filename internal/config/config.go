package config

type Server struct {
	Host    string `yaml:"host" env:"HOST" env-default:"0.0.0.0"`
	Port    string `yaml:"port" env:"PORT" env-default:"8080"`
	JokeURL string `yaml:"jokeurl"`
}
