package addr

import (
	"bytes"

	manet "gx/ipfs/QmSGL5Uoa6gKHgBBwQG8u1CWKUC8ZnwaZiLgFVTFBR2bxr/go-multiaddr-net"
	mafmt "gx/ipfs/QmUfnfCzL17Lfft4NtZZGGU34PP72xgecFVfusYt6WNwYx/mafmt"
	ma "gx/ipfs/QmW8s4zTsUoX1Q6CeYxVKPyqSKbF7H1YDUyTostBtZ8DaG/go-multiaddr"
)

func isFDCostlyTransport(a ma.Multiaddr) bool {
	return mafmt.TCP.Matches(a)
}

type AddrList []ma.Multiaddr

func (al AddrList) Len() int {
	return len(al)
}

func (al AddrList) Swap(i, j int) {
	al[i], al[j] = al[j], al[i]
}

func (al AddrList) Less(i, j int) bool {
	a := al[i]
	b := al[j]

	// dial localhost addresses next, they should fail immediately
	lba := manet.IsIPLoopback(a)
	lbb := manet.IsIPLoopback(b)
	if lba {
		if !lbb {
			return true
		}
	}

	// dial utp and similar 'non-fd-consuming' addresses first
	fda := isFDCostlyTransport(a)
	fdb := isFDCostlyTransport(b)
	if !fda {
		if fdb {
			return true
		}

		// if neither consume fd's, assume equal ordering
		return false
	}

	// if 'b' doesnt take a file descriptor
	if !fdb {
		return false
	}

	// if 'b' is loopback and both take file descriptors
	if lbb {
		return false
	}

	// for the rest, just sort by bytes
	return bytes.Compare(a.Bytes(), b.Bytes()) > 0
}
