litter
=======

Litter is not about anonymity or privacy.  It is about allodial rights
to your identity.

	> litter help
	NAME:
   		litter - Spreading litter across the dark web

	USAGE:
   		litter [global options] command [command options] [arguments...]

	VERSION:
   		0.0.0

	COMMANDS:
   		blocknotify, b       notify litter of new block
   		start, d             start the daemon process
   		register, d          register <name>
   		help, h              Shows a list of commands or help for one command
   
	GLOBAL OPTIONS:
   		--help, -h           show help
   		--version, -v        print the version

# TODO
	* implement name registration flow
		- namecoind calls litter on new blocks
		- init namecoind folder automatically echo -e "..."
		- gpg sign sent messages
		- check gpg sig on received messages
		- SOCKS5 connection to tor: http://blog.matee.net/post/tor-proxy-with-go
		- persistent connections
# Notes

https://github.com/prettymuchbryce/hellobitcoin

https://sourcegraph.com/blog/building-a-testable-webapp

Github releases: https://github.com/blog/1547-release-your-software

Namecoin: https://github.com/namecoin/namecoin

build for osx: https://wiki.namecoin.info/index.php?title=Build_Namecoin_From_Source

API Calls:
https://dot-bit.org/Client_API_calls_list
https://en.bitcoin.it/wiki/Original_Bitcoin_client/API_Calls_list

Config File (for callbacks):

https://en.bitcoin.it/wiki/Running_Bitcoin

Examples of custom commands:
https://github.com/conformal/btcws/blob/85cc323e34e694615c4364ebe97010d7c3197952/cmds.go

Example bitcoin client in [golang](https://en.bitcoin.it/wiki/API_reference_%28JSON-RPC%29)


forget about building rpc service for now with go

0.  Register namecoin address
1.  Launch go server
2.  Register tor hidden service and point it to go server: http://dominicbunch.wordpress.com/2014/03/14/how-to-make-a-tor-hidden-service/
3.  Get .onion address and update namecoin

ncurses: https://code.google.com/p/goncurses/

shell: http://golang.org/pkg/os/exec/#example_Cmd_Start

testnet faucet: https://nmctest.net/

RESTful go:

http://www.gorillatoolkit.org/pkg/mux
https://github.com/emicklei/go-restful

Persistence:

https://github.com/mattn/go-sqlite3

Go app structure:

https://medium.com/@benbjohnson/structuring-applications-in-go-3b04be4ff091

Go upnp:

https://github.com/huin/goupnp

ZeroMQ:

https://github.com/pebbe/zmq4

Linking go statically:

http://blog.hashbangbash.com/2014/04/linking-golang-statically/

JSON Unmarshalling:

http://blog.golang.org/json-and-go

GPG signing commits:

http://stackoverflow.com/questions/10077996/sign-git-commits-with-gpg

Command line lib:
https://github.com/codegangsta/cli

# commands:

curl -v --socks5-hostname 127.0.0.1:9050 acxjf2dhepeps7ts.onion:9191
curl --socks5-hostname 127.0.0.1:9050 acxjf2dhepeps7ts.onion:9191 -H "Accept: application/json" -X POST --data @test/fixtures/msg.json

bin/tor --SOCKSPort 9070 --DataDirectory ./.tor --HiddenServiceDir ./hidden_service --HiddenServicePort 7777

bin/namecoind -testnet -datadir=namecoin/ -dbcache=400 -printtoconsole -walletpath=./testnet-wallet.dat

go test -v ./...

name_firstupdate id/augustus d2eb8a142ed154d500 $(cat test/fixtures/record.json)
