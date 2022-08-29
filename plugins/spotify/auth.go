package spotify

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/oauth2"
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

}

func (m *Model) OauthHandler(w http.ResponseWriter, req *http.Request) {
	code := req.URL.Query().Get("code")
	tok, err := m.oauthCfg.Exchange(context.Background(), code, oauth2.AccessTypeOffline)
	if err != nil {
		panic(err)
	}

	m.client = m.oauthCfg.Client(context.Background(), tok)
	fmt.Fprintf(w, "Success! You can now close this window.")
}
