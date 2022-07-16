package main

import (
	"errors
	"fmt"
	"net/http"
	"regexp"
	"strings"

	// "context"
	"math/big"

	// ierc20 "priceapi/ierc20"
	irouter "priceapi/irouter"
	util "priceapi/util"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const JOE_ROUTER string = "0x60aE616a2155Ee3d9A68541Ba4544862310933d4"
const USDC string = "0xB97EF9Ef8734C71904D8002F8b6Bc66Dd9c48a6E"
const WAVAX string = "0xb31f66aa3c1e785363f0875a1b74e27b85fd66c7"

func main() {
	http.HandleFunc("/", readPath)
	http.ListenAndServe(":8080", nil)
}

func readPath(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	divides := strings.Split(path, "/")
	if (len(divides) == 2 && divides[1] != "") || (len(divides) == 3 && divides[2] == "") {
		//io.WriteString(w, divides[1])

		err := isValidAddress(divides[1])
		if err != nil {
			fmt.Fprintf(w, "error")
		}

		price, err := queryAVAX(common.HexToAddress(divides[1]))
		if err != nil {
			fmt.Fprintf(w, "error")
		}
		
		// io.WriteString(w, price.String())
		fmt.Fprintf(w, "%v", util.ToDecimal(price, 6))
	} else {
		fmt.Fprintf(w, "error")
	}
}

func queryAVAX(address common.Address) (*big.Int, error) {
	client, err := ethclient.Dial("https://api.avax.network/ext/bc/C/rpc")
	if err != nil {
		return nil, err
	}

	// balance, err := client.BalanceAt(context.Background(), address, nil)
	// if err != nil {
	// 	return nil, err
	// }

	// tokenAddress := common.HexToAddress("0xB97EF9Ef8734C71904D8002F8b6Bc66Dd9c48a6E")
	// token, err := ierc20.NewIERC20(tokenAddress, client)
	// if err != nil {
	// 	log.Fatalf("Failed to instantiate a Token contract: %v", err)
	// }

	// balance, err:= token.BalanceOf(nil, address)
	// if err != nil {
	// 	log.Fatalf("Failed to query the contract: %v", err)
	// }

	addrRouter := common.HexToAddress(JOE_ROUTER)
	router, err := irouter.NewIPancakeRouter01(addrRouter, client)
	if err != nil {
		return nil, err
	}

	value := big.NewInt(1000000000000000000)
	route := []common.Address{address, common.HexToAddress(WAVAX), common.HexToAddress(USDC)}
	price, err := router.GetAmountsOut(nil, value, route)
	if err != nil {
		return nil, err
	}

	return price[2], nil
}

func isValidAddress(address string) error {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	if re.MatchString(address) {
		return nil
	} else {
		return errors.New("invalid address")
	}
}

//// abigen --sol IRouter.sol --pkg store --out irouter.go
