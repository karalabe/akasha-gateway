# Akasha Gateway

Restful HTTP API for the Akasha decentralized social network. Work in progress!

## Installation

If you have a local Go development environment set up, the Akasha gateway can be downloaded and built from source with the usual `go get` workflow. Please note, the gateway contains certain C code too (secp256k1 crypto library) and as such requires a C compiler too.

```
$ go get -v github.com/karalabe/akasha-gateway
```

Alternatively, the Akasha gateway is also available as an automated build from Docker Hub. This is based on the Alpine Linux distribution to keep the image size small (12MB download).

```
$ docker pull karalabe/akasha-gateway
```

## Usage

The Akasha gateway has 3 important options that need to be set:

 * The host name on which the RESTful HTTP API server is listening for remote requests (to advertise the correct URL on the documentation page). This is normally the domain of the machine running the gateway (e.g. akasha.karalabe.com), defaulting to `localhost`.
 * The port number on which to listen for remote HTTP connections. This is normally set to `80` (requires low-port privileges), defaulting to `8080`.
 * The data directory where the Ethereum blockchain and the IPFS caches will end up in. This is normally set to a reliable data partition, defaulting to `$HOME/.akasha`.

Using the Go binary directly, the above options can be set via command line flags:

```
$ akasha-gateway --host=akasha.karalabe.com --port=80 --datadir=/my/data/dir
```

Using docker is a bit more involved, as the port cannot be changed via a flag, rather docker needs to be told to route it differently. As for the data directory, docker needs to explicitly mount an external folder into the container to ensure the data survives a restart:

```
$ docker run -p 80:8080 -v /my/data/dir:/root/.akasha:rw karalabe/akasha-gateway --host=akasha.karalabe.com
```

For a list of flags, please consult `akasha-gateway --help` or `docker run karalabe/akasha-gateway --help`.
