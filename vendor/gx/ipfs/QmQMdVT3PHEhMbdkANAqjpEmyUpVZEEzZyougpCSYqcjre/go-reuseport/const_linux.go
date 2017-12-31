// +build linux

package reuseport

import (
	"gx/ipfs/QmPXvegq26x982cQjSfbTvSzZXn7GiaMwhhVPHkeTEhrPT/sys/unix"
)

var soReusePort = 15 // this is not defined in unix go pkg.
var soReuseAddr = unix.SO_REUSEADDR
