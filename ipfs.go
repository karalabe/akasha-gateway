package main

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	peer "gx/ipfs/QmWNY7dV54ZDYmTA1ykVdwNCqC11mpU4zSUp6XDpLTH9eG/go-libp2p-peer"
	leveldb "gx/ipfs/QmYnCBXxoyoS38vtNQjjpRwZTiUnpuuKpapxMNaDfyQRLf/go-ds-leveldb"
	crypto "gx/ipfs/QmaPbCnUMBohSGo3KnxEa2bHqyJVVeEEcwtqJAYxerieBo/go-libp2p-crypto"
	"gx/ipfs/QmbBhyDKsY4mbY6xsKt3qu9Y7FPvMJ6qbD8AMjYYvPRw1g/goleveldb/leveldb/opt"
	"gx/ipfs/QmdHG8MAuARdGHxx4rPQASLcvhz24fzjSQq7AJRAQEorq5/go-datastore/sync"
	"path/filepath"

	"github.com/ipfs/go-ipfs/core"
	"github.com/ipfs/go-ipfs/merkledag"
	"github.com/ipfs/go-ipfs/path"
	"github.com/ipfs/go-ipfs/repo"
	cfg "github.com/ipfs/go-ipfs/repo/config"
)

// ipfs is a small utilty wrapper around the IPFS library.
type ipfs struct {
	node *core.IpfsNode
}

// makeIpfs assembles a go-ipfs in-process node to interact with the IPFS network.
func makeIpfs(datadir string) (*ipfs, error) {
	// Configure the data repository layer for IPFS
	ds, err := leveldb.NewDatastore(filepath.Join(datadir, "ipfs"), &leveldb.Options{
		OpenFilesCacheCapacity: 128,
		BlockCacheCapacity:     512 * opt.MiB,
		WriteBuffer:            256 * opt.MiB, // Two of these are used internally
	})
	if err != nil {
		return nil, err
	}
	priv, pub, err := crypto.GenerateKeyPairWithReader(crypto.RSA, 1024, rand.Reader)
	if err != nil {
		return nil, err
	}
	pid, err := peer.IDFromPublicKey(pub)
	if err != nil {
		return nil, err
	}
	privkeyb, err := priv.Bytes()
	if err != nil {
		return nil, err
	}
	conf := cfg.Config{
		Bootstrap: cfg.DefaultBootstrapAddresses,
	}
	conf.Addresses.Swarm = []string{"/ip4/0.0.0.0/tcp/4001"}
	conf.Identity.PeerID = pid.Pretty()
	conf.Identity.PrivKey = base64.StdEncoding.EncodeToString(privkeyb)

	// Assemble the IPFS node and return it
	node, err := core.NewNode(context.TODO(), &core.BuildCfg{
		Online:    true,
		Permanent: true,
		Repo:      &repo.Mock{D: sync.MutexWrap(ds), C: conf},
	})
	if err != nil {
		return nil, err
	}
	return &ipfs{
		node: node,
	}, nil
}

// Close implements the closer interface.
func (n *ipfs) Close() error {
	return n.node.Close()
}

// Links retrieves the links gathered together by an IPFS root object.
func (n *ipfs) Links(ctx context.Context, multihash string) (map[string]string, error) {
	obj, err := core.Resolve(ctx, n.node.Namesys, n.node.Resolver, path.Path(multihash))
	if err != nil {
		return nil, err
	}
	links := make(map[string]string)
	for _, link := range obj.Links() {
		links[link.Name] = link.Cid.String()
	}
	return links, nil
}

// Content retrieves the raw data content of an IPFS object.
func (n *ipfs) Content(ctx context.Context, multihash string) ([]byte, error) {
	obj, err := core.Resolve(ctx, n.node.Namesys, n.node.Resolver, path.Path(multihash))
	if err != nil {
		return nil, err
	}
	return obj.(*merkledag.ProtoNode).Data(), nil
}
