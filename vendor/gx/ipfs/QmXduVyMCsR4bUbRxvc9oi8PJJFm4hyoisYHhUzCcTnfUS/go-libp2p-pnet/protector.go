package pnet

import (
	"io"

	tconn "gx/ipfs/QmQGRkVSA5vTXt9WpJ6nBGnrvq9zRNsfjoNPq8uQrhnBoq/go-libp2p-transport"
	ipnet "gx/ipfs/QmXcN1kXchSvodd2MGypWXXirXk7GigQ7WVyWGYpukag6J/go-libp2p-interface-pnet"
)

var _ ipnet.Protector = (*protector)(nil)

// NewProtector creates ipnet.Protector instance from a io.Reader stream
// that should include Multicodec encoded V1 PSK.
func NewProtector(input io.Reader) (ipnet.Protector, error) {
	psk, err := decodeV1PSK(input)
	if err != nil {
		return nil, err
	}
	return NewV1ProtectorFromBytes(psk)
}

// NewV1ProtectorFromBytes creates ipnet.Protector of the V1 version.
func NewV1ProtectorFromBytes(psk *[32]byte) (ipnet.Protector, error) {
	return &protector{
		psk:         psk,
		fingerprint: fingerprint(psk),
	}, nil
}

type protector struct {
	psk         *[32]byte
	fingerprint []byte
}

func (p protector) Protect(in tconn.Conn) (tconn.Conn, error) {
	return newPSKConn(p.psk, in)
}
func (p protector) Fingerprint() []byte {
	return p.fingerprint
}
