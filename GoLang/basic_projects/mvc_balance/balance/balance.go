package balance

import (
	"context"
	"mvc_balance/utils"
	"net/http"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)


func BalanceOf(addr common.Address) (string, *utils.ApplicationError) {
	client, err := ethclient.Dial("https://api.avax.network/ext/bc/C/rpc")
	if err != nil {
		return "", &utils.ApplicationError{
			Message: "Conn error",
			StatusCode: http.StatusBadRequest,
			Code: "no_conn",
		}
	}

	balance, err := client.BalanceAt(context.Background(), addr, nil)
	if err != nil {
		return "", &utils.ApplicationError{
			Message: "Query error",
			StatusCode: http.StatusBadRequest,
			Code: "q_err",
		}
	}

	return utils.ToDecimal(balance, 18).String(), nil
}