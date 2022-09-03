package spotify

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"sync"

	"github.com/mitchellh/go-homedir"
	"github.com/pkg/browser"
	"golang.org/x/oauth2"
)

func writeToCache(path string, t *oauth2.Token) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Println(err.Error())
		}
	}()

	return json.NewEncoder(f).Encode(t)
}

func readFromCache(path string) (*oauth2.Token, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Println(err.Error())
		}
	}()

	var t oauth2.Token
	if err := json.NewDecoder(f).Decode(&t); err != nil {
		return nil, err
	}

	return &t, nil
}

func getCachePath() (string, bool) {
	home, err := homedir.Dir()
	if err != nil {
		panic(err)
	}

	termhubConf := path.Join(home, ".termhub")

	entries, err := os.ReadDir(termhubConf)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			err = os.MkdirAll(termhubConf, 0755)
			if err != nil {
				panic(err)
			}
			return path.Join(termhubConf, "spotify.json"), false
		}
	}
	for _, entry := range entries {
		if entry.Name() == "spotify.json" {
			return path.Join(termhubConf, "spotify.json"), true
		}
	}
	return path.Join(termhubConf, "spotify.json"), false
}

func ListenForCode(clientID, clientSecret string) *http.Client {
	cfg := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint: oauth2.Endpoint{
			TokenURL: "https://accounts.spotify.com/api/token",
			AuthURL:  "https://accounts.spotify.com/authorize",
		},
		RedirectURL: "http://localhost:8080/spotify/redirect",
		Scopes:      []string{"user-read-playback-state", "user-modify-playback-state"},
	}

	// check cache (ok means it exists already)
	cachePath, ok := getCachePath()
	if ok {
		t, err := readFromCache(cachePath)
		if err != nil {
			log.Println("an error occurred checking cached credentials, logging in again")
		}
		if t.Valid() {
			return cfg.Client(context.Background(), t)
		}
	}

	// setup server
	var client *http.Client
	var wg sync.WaitGroup

	wg.Add(1)

	srv := http.Server{
		Addr: ":8080",
	}

	srv.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/spotify/redirect" {
			code := r.URL.Query().Get("code")
			if code == "" {
				log.Println("no code found")
				return
			}
			token, err := cfg.Exchange(context.Background(), code)
			if err != nil {
				log.Println(err)
				return
			}
			client = cfg.Client(context.Background(), token)
			fmt.Fprintf(w, "Success! You can close this window now.")
			err = writeToCache(cachePath, token)
			if err != nil {
				log.Printf("An unexpected error occurred while caching credentials\n%s\n", err.Error())
				log.Println("unable to cache credentials")
			}
			wg.Done()
		}
	})

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				panic(err)
			}
		}
	}()

	// prompt login
	url := cfg.AuthCodeURL("state", oauth2.AccessTypeOffline)
	if err := browser.OpenURL(url); err != nil {
		panic(err)
	}

	wg.Wait()
	srv.Shutdown(context.Background())

	return client
}
