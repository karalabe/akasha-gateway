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

//go:generate abigen -abi ./aeth.abi -pkg contracts -type AETH -out aeth.go
//go:generate abigen -abi ./essence.abi -pkg contracts -type Essence -out essence.go
//go:generate abigen -abi ./profileResolver.abi -pkg contracts -type ProfileResolver -out profileResolver.go
//go:generate abigen -abi ./profileRegistrar.abi -pkg contracts -type ProfileRegistrar -out profileRegistrar.go
//go:generate abigen -abi ./entries.abi -pkg contracts -type Entries -out entries.go
//go:generate abigen -abi ./feed.abi -pkg contracts -type Feed -out feed.go

// Package contracts are the Go-bound Akasha smart contracts.
package contracts
