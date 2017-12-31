package websocket

import (
	"context"

	tpt "gx/ipfs/QmQGRkVSA5vTXt9WpJ6nBGnrvq9zRNsfjoNPq8uQrhnBoq/go-libp2p-transport"
	manet "gx/ipfs/QmSGL5Uoa6gKHgBBwQG8u1CWKUC8ZnwaZiLgFVTFBR2bxr/go-multiaddr-net"
	ma "gx/ipfs/QmW8s4zTsUoX1Q6CeYxVKPyqSKbF7H1YDUyTostBtZ8DaG/go-multiaddr"
	ws "gx/ipfs/QmZH5VXfAJouGMyCCHTRPGCT3e5MG9Lu78Ln3YAYW1XTts/websocket"
)

type dialer struct{}

func (d *dialer) Dial(raddr ma.Multiaddr) (tpt.Conn, error) {
	return d.DialContext(context.Background(), raddr)
}

func (d *dialer) DialContext(ctx context.Context, raddr ma.Multiaddr) (tpt.Conn, error) {
	wsurl, err := parseMultiaddr(raddr)
	if err != nil {
		return nil, err
	}

	wscon, _, err := ws.DefaultDialer.Dial(wsurl, nil)
	if err != nil {
		return nil, err
	}

	mnc, err := manet.WrapNetConn(NewConn(wscon, nil))
	if err != nil {
		wscon.Close()
		return nil, err
	}

	return &wsConn{
		Conn: mnc,
	}, nil
}

func (d *dialer) Matches(a ma.Multiaddr) bool {
	return WsFmt.Matches(a)
}
