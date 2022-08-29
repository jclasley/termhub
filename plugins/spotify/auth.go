package spotify

import (
	"encoding/json"
	"github.com/pkg/browser"
	"golang.org/x/oauth2"
	"net/http"
)

func getToken() string {
	req, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token?grant_type=client_credentials", nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth("6e6edf8d952f42fca28dc9f5f2f8751f", "9059472d718c49f1ac31acb8926d6a5f")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	data := struct {
		Token string `json:"access_token"`
	}{}
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	return data.Token
}

func oauthToken() *oauth2.Config {
	//ctx := context.Background()
	cfg := &oauth2.Config{
		ClientID:     "6e6edf8d952f42fca28dc9f5f2f8751f",
		ClientSecret: "9059472d718c49f1ac31acb8926d6a5f",
		Endpoint: oauth2.Endpoint{
			TokenURL: "https://accounts.spotify.com/api/token",
			AuthURL:  "https://accounts.spotify.com/authorize",
		},
		RedirectURL: "http://localhost:8080/redirect",
		Scopes:      []string{"user-read-playback-state"},
	}

	url := cfg.AuthCodeURL("state", oauth2.AccessTypeOffline)

	// Open the authorization URL in the user's browser
	if err := browser.OpenURL(url); err != nil {
		panic(err)
	}

	return cfg
}
