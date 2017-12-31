package testutil

import (
	"context"
	"testing"

	inet "gx/ipfs/QmU4vCDZTPLDqSDKguWbHCiUe46mZUtmM2g2suBZ9NE8ko/go-libp2p-net"
	swarm "gx/ipfs/QmUhvp4VoQ9cKDVLqAxciEKdm8ymBx2Syx4C1Tv6SmSTPa/go-libp2p-swarm"
	ma "gx/ipfs/QmW8s4zTsUoX1Q6CeYxVKPyqSKbF7H1YDUyTostBtZ8DaG/go-multiaddr"
	pstore "gx/ipfs/QmYijbtjCxFEjSXaudaQAUz3LN5VKLssm8WCUsRoqzXmQR/go-libp2p-peerstore"
	tu "gx/ipfs/QmeDA8gNhvRTsbrjEieay5wezupJDiky8xvCzDABbsGzmp/go-testutil"
)

func GenSwarmNetwork(t *testing.T, ctx context.Context) *swarm.Network {
	p := tu.RandPeerNetParamsOrFatal(t)
	ps := pstore.NewPeerstore()
	ps.AddPubKey(p.ID, p.PubKey)
	ps.AddPrivKey(p.ID, p.PrivKey)
	n, err := swarm.NewNetwork(ctx, []ma.Multiaddr{p.Addr}, p.ID, ps, nil)
	if err != nil {
		t.Fatal(err)
	}
	ps.AddAddrs(p.ID, n.ListenAddresses(), pstore.PermanentAddrTTL)
	return n
}

func DivulgeAddresses(a, b inet.Network) {
	id := a.LocalPeer()
	addrs := a.Peerstore().Addrs(id)
	b.Peerstore().AddAddrs(id, addrs, pstore.PermanentAddrTTL)
}

var RandPeerID = tu.RandPeerID
