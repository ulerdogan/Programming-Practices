package ethereum

import (
    "github.com/ethereum/go-ethereum/ethclient"
	"log"
)

var Client *ethclient.Client

func init() {
	var err error
	Client, err = ethclient.Dial("https://api.avax.network/ext/bc/C/rpc")
	if err != nil {
		log.Fatalln(err)
		return
	}
}