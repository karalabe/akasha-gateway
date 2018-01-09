package main

import (
	"bytes"
	"flag"
	"fmt"
	"html/template"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/fdlimit"
	"github.com/ethereum/go-ethereum/log"
	"github.com/karalabe/akasha/spec"
	colorable "github.com/mattn/go-colorable"
)

var (
	hostFlag = flag.String("host", "localhost", "Host name to advertise the API in the docs")
	baseFlag = flag.String("base", "/beta/", "Base path on which to serve the RESTful API")
	portFlag = flag.Int("port", 8080, "Port number on which to listen for web requests")

	dataFlag = flag.String("data", filepath.Join(os.Getenv("HOME"), ".akasha"), "Data directory to use for Ethereum and IPFS")

	logLevelFlag = flag.Int("loglevel", int(log.LvlInfo), "Log verbosity to use by Ethereum and Akasha")
	fdLimitFlag  = flag.Int("fdlimit", 2048, "Number of file descriptors to request from the OS")
)

func main() {
	// Set up the process for the Akasha API
	flag.Parse()
	fdlimit.Raise(uint64(*fdLimitFlag))
	log.Root().SetHandler(log.LvlFilterHandler(log.Lvl(*logLevelFlag), log.StreamHandler(colorable.NewColorableStderr(), log.TerminalFormat(true))))

	// Create and boot up the Ethereum node
	geth, err := makeGeth(*dataFlag)
	if err != nil {
		log.Crit("Failed to create Ethereum node", "err", err)
	}
	if err = geth.Start(); err != nil {
		log.Crit("Failed to start Ethereum node", "err", err)
	}
	defer geth.Stop()

	// Create and boot up the IPFS node
	ipfs, err := makeIpfs(*dataFlag)
	if err != nil {
		log.Crit("Failed to create IPFS node", "err", err)
	}
	defer ipfs.Close()

	// Interface the Akasha smart contracts and IPFS backends
	akasha, err := makeAkasha(geth, ipfs, &config{
		AETHAddress:      common.HexToAddress("0x398a7a69f3c59181a1ffe34bed11dcb5df863a8a"),
		EssenceAddress:   common.HexToAddress("0xd8c6dd6d69ad4e37d6ea1b58b3ea9e06857c2550"),
		ResolverAddress:  common.HexToAddress("0xa6100e99dda74e8aad319f4c4bae098694c910a4"),
		RegistrarAddress: common.HexToAddress("0x3f4df77876fb393975daa074301d2913699c28dc"),
		EntriesAddress:   common.HexToAddress("0x315406c6e19a0781a65d65dba2f40b4d96dd8952"),
	})
	if err != nil {
		log.Crit("Failed to interface Akasha", "err", err)
	}
	if err = akasha.Prefetch(); err != nil {
		log.Crit("Failed to start Akasha prefetcher", "err", err)
	}
	// Create the web handler to server the API specs and actual website
	api, err := makeAPI(akasha, *baseFlag)
	if err != nil {
		log.Crit("Failed to interface Akasha", "err", err)
	}
	http.Handle("/", http.HandlerFunc(docsHandler))
	http.Handle(*baseFlag, api)

	http.ListenAndServe(fmt.Sprintf(":%d", *portFlag), nil)
}

// docsHandler handles all non-API requests, flattening and returning the docs.
func docsHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve the data content the user wishes to load
	path := r.URL.String()
	if path == "/" {
		path = "/index.html"
	}
	blob, err := spec.Asset(strings.TrimPrefix(path, "/"))
	if err != nil {
		http.NotFound(w, r)
		return
	}
	// If we're serving the docs, expand any template variables
	if path == "/index.html" {
		tmpl := template.Must(template.New("").Parse(string(blob)))
		buff := new(bytes.Buffer)

		host := *hostFlag
		if *portFlag != 80 && *portFlag != 443 {
			host += fmt.Sprintf(":%d", *portFlag)
		}
		if err = tmpl.Execute(buff, map[string]string{"Host": host, "Base": *baseFlag}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		blob = buff.Bytes()
	}
	// Resolve the file's mime type and deliver
	w.Header().Set("Content-Type", mime.TypeByExtension(filepath.Ext(path)))
	w.Write(blob)
}
