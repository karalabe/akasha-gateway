//go:generate abigen -abi ./aeth.abi -pkg contracts -type AETH -out aeth.go
//go:generate abigen -abi ./essence.abi -pkg contracts -type Essence -out essence.go
//go:generate abigen -abi ./profileResolver.abi -pkg contracts -type ProfileResolver -out profileResolver.go
//go:generate abigen -abi ./profileRegistrar.abi -pkg contracts -type ProfileRegistrar -out profileRegistrar.go
//go:generate abigen -abi ./entries.abi -pkg contracts -type Entries -out entries.go
//go:generate abigen -abi ./feed.abi -pkg contracts -type Feed -out feed.go

package contracts
