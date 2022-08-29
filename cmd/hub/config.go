package main

import "github.com/kelseyhightower/envconfig"

type Config struct {
	SpotifyClientID string `envconfig:"SPOTIFY_CLIENT_ID"`
	SpotifySecret   string `envconfig:"SPOTIFY_SECRET"`
}

func GetConfig() Config {
	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		panic(err)
	}
	return cfg
}
