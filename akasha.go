package main

import (
	"context"
	"encoding/json"
	"errors"
	"strings"
	"time"

	"github.com/btcsuite/btcutil/base58"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/node"
	"github.com/karalabe/akasha/contracts"
)

var (
	// errUnknownUser is returned if a user cannot be found in the Akasha network.
	errUnknownUser = errors.New("unknown user")
)

// akasha represents the interface to the Akasha smart contracts.
type akasha struct {
	ipfs      *ipfs
	aeth      *contracts.AETH
	essence   *contracts.Essence
	resolver  *contracts.ProfileResolver
	registrar *contracts.ProfileRegistrar
	feed      *contracts.Feed
}

// config represents the configurations for the Akasha smart contract system.
type config struct {
	AETHAddress      common.Address
	EssenceAddress   common.Address
	ResolverAddress  common.Address
	RegistrarAddress common.Address
}

// makeAkasha creates a programatic interface to the Akasha smart contracts.
func makeAkasha(geth *node.Node, ipfs *ipfs, conf *config) (*akasha, error) {
	// Attach to the Geth client and bind the Akasha contracts
	rpc, err := geth.Attach()
	if err != nil {
		log.Crit("Failed to attach to Ethereum client", "err", err)
	}
	client := ethclient.NewClient(rpc)

	aeth, err := contracts.NewAETH(conf.AETHAddress, client)
	if err != nil {
		return nil, err
	}
	essence, err := contracts.NewEssence(conf.EssenceAddress, client)
	if err != nil {
		return nil, err
	}
	resolver, err := contracts.NewProfileResolver(conf.ResolverAddress, client)
	if err != nil {
		return nil, err
	}
	registrar, err := contracts.NewProfileRegistrar(conf.RegistrarAddress, client)
	if err != nil {
		return nil, err
	}
	return &akasha{
		ipfs:      ipfs,
		aeth:      aeth,
		essence:   essence,
		resolver:  resolver,
		registrar: registrar,
	}, nil
}

// user represents all the known information about an Akasha user. The reason
// beind the nullable strings is to allow signalling unreachable IPFS content.
type user struct {
	User    string         `json:"user"`
	Name    *string        `json:"name"`
	Address common.Address `json:"address"`
	About   *string        `json:"about"`
	Avatar  *string        `json:"avatar"`
	Cover   *string        `json:"cover"`
	Links   []string       `json:"links"`
	Tips    bool           `json:"tips"`
	Aether  *hexutil.Big   `json:"aether"`
	Bonded  *hexutil.Big   `json:"bonded"`
	Cycling *hexutil.Big   `json:"cycling"`
	Mana    *hexutil.Big   `json:"mana"`
	Spent   *hexutil.Big   `json:"spent"`
	Essence *hexutil.Big   `json:"essence"`
	Karma   *hexutil.Big   `json:"karma"`
}

// userByID does an ENS lookup to get the registration node of the user and
// retrieves all known infos associated with it.
func (a *akasha) userByID(id string) (*user, error) {
	var label [32]byte
	copy(label[:], id)

	node, err := a.registrar.Hash(nil, label)
	if err != nil {
		return nil, err
	}
	return a.user(node)
}

// userByAddress does a reverse ENS lookup to get the registration node of the
// user and retrieves all known infos associated with it.
func (a *akasha) userByAddress(addr common.Address) (*user, error) {
	node, err := a.resolver.Reverse(nil, addr)
	if err != nil {
		return nil, err
	}
	return a.user(node)
}

// user retrieves all the known infos about a user identified by it's ENS node.
func (a *akasha) user(node [32]byte) (*user, error) {
	// Retrieve the profile infos from the profile resolver
	profile, err := a.resolver.Resolve(nil, node)
	if err != nil {
		return nil, err
	}
	if profile.Addr == (common.Address{}) {
		return nil, errUnknownUser
	}
	// Retrieve profile details from the IPFS profile objects (failures are fine-ish)
	root := base58.Encode(append([]byte{profile.Fn, profile.DigestSize}, profile.Hash[:]...))

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var (
		name   *string
		about  *string
		avatar *string
		cover  *string
		links  []string
	)
	objs, err := a.ipfs.Links(ctx, root)
	if err == nil {
		avatar = new(string)
		*avatar = "https://ipfs.io/ipfs/" + objs["avatar"]

		var (
			blob []byte
			prof struct {
				FirstName string `json:"firstName"`
				LastName  string `json:"lastName"`
			}
			bg struct {
				XXL struct {
					Src string `json:"src"`
				} `json:"xxl"`
			}
			urls []struct {
				Url string `json:"url"`
			}
		)
		if blob, err = a.ipfs.Content(ctx, root); err == nil {
			if err = json.Unmarshal(blob, &prof); err != nil {
				return nil, err
			}
			name = new(string)
			*name = prof.FirstName + " " + prof.LastName
		}
		if blob, err = a.ipfs.Content(ctx, objs["about"]); err == nil {
			if err = json.Unmarshal(blob, &about); err != nil {
				return nil, err
			}
		}
		if blob, err = a.ipfs.Content(ctx, objs["backgroundImage"]); err == nil {
			if err = json.Unmarshal(blob, &bg); err != nil {
				return nil, err
			}
			if bg.XXL.Src != "" {
				cover = new(string)
				*cover = "https://ipfs.io/ipfs/" + bg.XXL.Src
			}
		}
		if blob, err = a.ipfs.Content(ctx, objs["links"]); err == nil {
			if err = json.Unmarshal(blob, &urls); err != nil {
				return nil, err
			}
			for _, url := range urls {
				links = append(links, url.Url)
			}
		}
	}
	// Return the token infos from the ledger contract
	balances, err := a.aeth.GetTokenRecords(nil, profile.Addr)
	if err != nil {
		return nil, err
	}
	credits, err := a.essence.GetCollected(nil, profile.Addr)
	if err != nil {
		return nil, err
	}
	mana, err := a.essence.Mana(nil, profile.Addr)
	if err != nil {
		return nil, err
	}
	return &user{
		User:    strings.TrimRight(string(profile.AkashaId[:]), string([]byte{0})),
		Name:    name,
		Address: profile.Addr,
		About:   about,
		Avatar:  avatar,
		Cover:   cover,
		Links:   links,
		Tips:    profile.DonationsEnabled,
		Aether:  (*hexutil.Big)(balances.Free),
		Bonded:  (*hexutil.Big)(balances.Bonded),
		Cycling: (*hexutil.Big)(balances.Cycling),
		Mana:    (*hexutil.Big)(mana.Total),
		Spent:   (*hexutil.Big)(mana.Spent),
		Essence: (*hexutil.Big)(credits.Essence),
		Karma:   (*hexutil.Big)(credits.Karma),
	}, nil
}
