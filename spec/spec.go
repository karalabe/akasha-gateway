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

// Package spec contains the embedded static API doc website.
package spec

//go:generate sh -c "npm list spectacle-docs || npm install spectacle-docs"
//go:generate ./node_modules/.bin/spectacle api.yaml
//go:generate go-bindata -nometadata -pkg spec -o assets.go -prefix public ./public/...
