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

package main

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"gx/ipfs/QmNwUEK7QbwSqyKBu3mMtToo8SUc6wQJ7gdZq4gGGJqfnf/go-ipld-format"
	peer "gx/ipfs/QmWNY7dV54ZDYmTA1ykVdwNCqC11mpU4zSUp6XDpLTH9eG/go-libp2p-peer"
	leveldb "gx/ipfs/QmYnCBXxoyoS38vtNQjjpRwZTiUnpuuKpapxMNaDfyQRLf/go-ds-leveldb"
	crypto "gx/ipfs/QmaPbCnUMBohSGo3KnxEa2bHqyJVVeEEcwtqJAYxerieBo/go-libp2p-crypto"
	"gx/ipfs/QmbBhyDKsY4mbY6xsKt3qu9Y7FPvMJ6qbD8AMjYYvPRw1g/goleveldb/leveldb/opt"
	dbsync "gx/ipfs/QmdHG8MAuARdGHxx4rPQASLcvhz24fzjSQq7AJRAQEorq5/go-datastore/sync"
	"path/filepath"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/log"
	"github.com/ipfs/go-ipfs/core"
	"github.com/ipfs/go-ipfs/merkledag"
	"github.com/ipfs/go-ipfs/path"
	"github.com/ipfs/go-ipfs/repo"
	cfg "github.com/ipfs/go-ipfs/repo/config"
)

// ipfs is a small utilty wrapper around the IPFS library.
type ipfs struct {
	node *core.IpfsNode

	pend map[string]struct{}
	lock sync.Mutex
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
		Repo:      &repo.Mock{D: dbsync.MutexWrap(ds), C: conf},
	})
	if err != nil {
		return nil, err
	}
	return &ipfs{
		node: node,
		pend: make(map[string]struct{}),
	}, nil
}

// Close implements the closer interface.
func (n *ipfs) Close() error {
	if err := n.node.Pinning.Flush(); err != nil {
		log.Error("Failed to flush IPFS pinner", "err", err)
	}
	return n.node.Close()
}

// Links retrieves the links gathered together by an IPFS root object.
func (n *ipfs) Links(ctx context.Context, multihash string) (map[string]string, error) {
	obj, err := n.resolve(ctx, multihash)
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
	obj, err := n.resolve(ctx, multihash)
	if err != nil {
		return nil, err
	}
	return obj.(*merkledag.ProtoNode).Data(), nil
}

// resolve attempts to resolve the requested multihash until the context times out,
// in case of which it returns an error but starts a longer background resolution.
func (n *ipfs) resolve(ctx context.Context, multihash string) (format.Node, error) {
	obj, err := core.Resolve(ctx, n.node.Namesys, n.node.Resolver, path.Path(multihash))
	if err != nil {
		select {
		case <-ctx.Done():
			// Request timed out, retrieve in the background
			go func() {
				// Make sure we only have one pending fetch per resource
				n.lock.Lock()
				if _, ok := n.pend[multihash]; ok {
					n.lock.Unlock()
					return
				}
				n.pend[multihash] = struct{}{}
				n.lock.Unlock()

				// Make sure we clean up the resource lock
				defer func() {
					n.lock.Lock()
					defer n.lock.Unlock()
					delete(n.pend, multihash)
				}()

				// Try to resolve the resource with a significantly larger timeout
				ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
				defer cancel()

				if _, err = core.Resolve(ctx, n.node.Namesys, n.node.Resolver, path.Path(multihash)); err != nil {
					log.Warn("Background IPFS resolution failed", "multihash", multihash)
					return
				}
				log.Info("Background IPFS resolution succeeded", "multihash", multihash)
			}()
		default:
		}
		return nil, err
	}
	// Object successfully resolved, pin it locally to retain for future calls
	if _, pinned, _ := n.node.Pinning.IsPinned(obj.Cid()); !pinned {
		if err := n.node.Pinning.Pin(ctx, obj, false); err != nil {
			log.Warn("Failed to ping IPFS object", "hash", multihash, "err", err)
		} else {
			log.Info("Pinned requested IPFS object", "hash", multihash)
		}
	}
	return obj, nil
}
