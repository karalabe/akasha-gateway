package main

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

// api represents a RESTful web API to act as a bridge into the Akasha platform.
type api struct {
	akasha *akasha
	base   string
}

// makeAPI creates, starts and returns a RESTful HTTP API into the Akasha platform.
func makeAPI(akasha *akasha, base string) (http.Handler, error) {
	return &api{
		akasha: akasha,
		base:   base,
	}, nil
}

// ServeHTTP is responsible for serving a single HTTP REST request from a browser.
func (api *api) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(strings.TrimSuffix(strings.TrimPrefix(r.URL.Path, api.base), "/"), "/")

	// Handle the first component of the API path
	var res interface{}

	switch path[0] {
	case "users":
		// Ensure we have an actual ID to look up
		if len(path) != 2 {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}
		// Depending on the ID type, either look up by address or username
		if common.IsHexAddress(path[1]) {
			user, err := api.akasha.userByAddress(common.HexToAddress(path[1]))
			if err != nil {
				if err == errUnknownUser {
					http.Error(w, err.Error(), http.StatusNotFound)
					return
				}
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			res = user
		} else {
			user, err := api.akasha.userByID(path[1])
			if err != nil {
				if err == errUnknownUser {
					http.Error(w, err.Error(), http.StatusNotFound)
					return
				}
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			res = user
		}

	default:
		http.NotFound(w, r)
		return
	}
	// Handled and response generated, serialize to user
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	enc.Encode(res)
}
