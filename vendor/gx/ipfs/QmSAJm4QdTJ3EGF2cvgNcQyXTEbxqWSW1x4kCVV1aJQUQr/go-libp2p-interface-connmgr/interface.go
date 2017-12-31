package ifconnmgr

import (
	"context"
	"time"

	inet "gx/ipfs/QmU4vCDZTPLDqSDKguWbHCiUe46mZUtmM2g2suBZ9NE8ko/go-libp2p-net"
	peer "gx/ipfs/QmWNY7dV54ZDYmTA1ykVdwNCqC11mpU4zSUp6XDpLTH9eG/go-libp2p-peer"
)

type ConnManager interface {
	TagPeer(peer.ID, string, int)
	UntagPeer(peer.ID, string)
	GetTagInfo(peer.ID) *TagInfo
	TrimOpenConns(context.Context)
	Notifee() inet.Notifiee
}

type TagInfo struct {
	FirstSeen time.Time
	Value     int
	Tags      map[string]int
	Conns     map[string]time.Time
}
