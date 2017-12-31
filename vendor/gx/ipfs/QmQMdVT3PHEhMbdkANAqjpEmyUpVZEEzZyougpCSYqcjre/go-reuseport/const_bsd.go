// +build darwin freebsd dragonfly netbsd openbsd

package reuseport

import (
	"gx/ipfs/QmPXvegq26x982cQjSfbTvSzZXn7GiaMwhhVPHkeTEhrPT/sys/unix"
)

var soReusePort = unix.SO_REUSEPORT
var soReuseAddr = unix.SO_REUSEADDR
