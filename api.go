package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

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
	// Prep the request for processing
	path := strings.Split(strings.TrimSuffix(strings.TrimPrefix(r.URL.Path, api.base), "/"), "/")

	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	timeout := 3 * time.Second
	if val := r.FormValue("timeout"); val != "" {
		n, err := strconv.Atoi(val)
		if err != nil {
			http.Error(w, fmt.Sprintf("invalid timeout: %v", err), http.StatusBadRequest)
			return
		}
		if n < 0 || n > 60 {
			http.Error(w, "invalid timeout range", http.StatusBadRequest)
			return
		}
		timeout = time.Duration(n) * time.Second
	}
	// Handle the first component of the API path
	var res interface{}

	switch path[0] {
	case "users":
		switch {
		case len(path) == 2:
			// User profile requested, look up and return
			user, err := api.user(path[1], timeout)
			if err != nil {
				if err == errUnknownUser {
					http.Error(w, err.Error(), http.StatusNotFound)
					return
				}
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			res = user

		case len(path) == 3 && path[2] == "entries":
			// User's entry list requested, filter for and return
			entries, err := api.entries(path[1], timeout)
			if err != nil {
				if err == errUnknownUser || err == errUnknownEntry {
					http.Error(w, err.Error(), http.StatusNotFound)
					return
				}
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			res = entries

		case len(path) == 4 && path[2] == "entries":
			// User's particular entry post requested, locate and return
			entry, err := api.entry(path[1], common.HexToHash(path[3]), timeout)
			if err != nil {
				if err == errUnknownUser || err == errUnknownEntry {
					http.Error(w, err.Error(), http.StatusNotFound)
					return
				}
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			res = entry
		}

	case "entries":
		// Anonymous entry requested, look up and return
		switch {
		case len(path) == 2:
			entry, err := api.akasha.Entry(common.HexToHash(path[1]), timeout)
			if err != nil {
				if err == errUnknownEntry {
					http.Error(w, err.Error(), http.StatusNotFound)
					return
				}
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			res = entry
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

// user tries to figure out whether id is an Akasha username or an Ethereum
// address and retrieves the Akasha profile associated with it.
func (api *api) user(id string, timeout time.Duration) (*User, error) {
	if common.IsHexAddress(id) {
		return api.akasha.UserByAddress(common.HexToAddress(id), timeout)
	}
	return api.akasha.UserByName(id, timeout)
}

// entries tries to figure out whether id is an Akasha username or an Ethereum
// address and retrieves the Akasha entries posted by it.
func (api *api) entries(id string, timeout time.Duration) ([]common.Hash, error) {
	if common.IsHexAddress(id) {
		return api.akasha.EntriesByAddress(common.HexToAddress(id), timeout)
	}
	return api.akasha.EntriesByName(id, timeout)
}

// entry tries to figure out whether id is an Akasha username or an Ethereum
// address and retrieves the Akasha entry belonging to it.
func (api *api) entry(id string, hash common.Hash, timeout time.Duration) (*Entry, error) {
	if common.IsHexAddress(id) {
		return api.akasha.EntryByAddress(common.HexToAddress(id), hash, timeout)
	}
	return api.akasha.EntryByName(id, hash, timeout)
}
