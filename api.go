// Akasha Gateway - API for legacy (web 2.0) applications
// Copyright (c) 2018 Péter Szilágyi. All rights reserved.
//
// The Akasha gateway is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or (at your
// option) any later version.
//
// The Akasha gateway is distributed in the hope that it will be useful, but
// WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY
// or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Lesser General Public
// License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the Akasha gateway. If not, see <http://www.gnu.org/licenses/>.

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

		case len(path) == 5 && path[2] == "entries" && path[4] == "comments":
			// User's entry's comment list requested, filter for and return
			comments, err := api.akasha.CommentsByEntry(common.HexToHash(path[3]), timeout)
			if err != nil {
				if err == errUnknownUser || err == errUnknownEntry {
					http.Error(w, err.Error(), http.StatusNotFound)
					return
				}
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			res = comments

		case len(path) == 6 && path[2] == "entries" && path[4] == "comments":
			// User's entry's comment requested, look up and return
			comment, err := api.akasha.CommentByEntry(common.HexToHash(path[3]), common.HexToHash(path[5]), timeout)
			if err != nil {
				if err == errUnknownUser || err == errUnknownEntry || err == errUnknownComment {
					http.Error(w, err.Error(), http.StatusNotFound)
					return
				}
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			res = comment

		case len(path) == 3 && path[2] == "comments":
			// User's comment list requested, filter for and return
			comments, err := api.comments(path[1], timeout)
			if err != nil {
				if err == errUnknownUser {
					http.Error(w, err.Error(), http.StatusNotFound)
					return
				}
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			res = comments

		case len(path) == 4 && path[2] == "comments":
			// User's comment list requested, filter for and return
			comments, err := api.comment(path[1], common.HexToHash(path[3]), timeout)
			if err != nil {
				if err == errUnknownUser {
					http.Error(w, err.Error(), http.StatusNotFound)
					return
				}
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			res = comments
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

		case len(path) == 3 && path[2] == "comments":
			// Anonymous entry's comment list requested, filter for and return
			comments, err := api.akasha.CommentsByEntry(common.HexToHash(path[1]), timeout)
			if err != nil {
				if err == errUnknownEntry {
					http.Error(w, err.Error(), http.StatusNotFound)
					return
				}
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			res = comments

		case len(path) == 4 && path[2] == "comments":
			// Anonymous entry's comment requested, look up and return
			comment, err := api.akasha.CommentByEntry(common.HexToHash(path[1]), common.HexToHash(path[3]), timeout)
			if err != nil {
				if err == errUnknownEntry || err == errUnknownComment {
					http.Error(w, err.Error(), http.StatusNotFound)
					return
				}
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			res = comment
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

// comments tries to figure out whether id is an Akasha username or an Ethereum
// address and retrieves the Akasha comments posted by it.
func (api *api) comments(id string, timeout time.Duration) ([]common.Hash, error) {
	if common.IsHexAddress(id) {
		return api.akasha.CommentsByAddress(common.HexToAddress(id), timeout)
	}
	return api.akasha.CommentsByName(id, timeout)
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

// comment tries to figure out whether id is an Akasha username or an Ethereum
// address and retrieves the Akasha comment belonging to it.
func (api *api) comment(id string, hash common.Hash, timeout time.Duration) (*Comment, error) {
	if common.IsHexAddress(id) {
		return api.akasha.CommentByAddress(common.HexToAddress(id), hash, timeout)
	}
	return api.akasha.CommentByName(id, hash, timeout)
}
