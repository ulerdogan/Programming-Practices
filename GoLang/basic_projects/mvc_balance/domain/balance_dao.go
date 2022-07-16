package domain

import (
	"mvc_balance/balance"
	"mvc_balance/utils"
	"net/http"
	"github.com/ethereum/go-ethereum/common"
)

type balanceDao struct {}

var BalanceDao balanceDao

func (blnc *balanceDao) GetBalance(addr common.Address) (*Balance, *utils.ApplicationError) {
	b, _ := balance.BalanceOf(addr)
	if b != "" {
		return &Balance{
			Balance: b,
			Nonce: 0,
			Network: "Avalanche",
		} , nil
	}

	return nil, &utils.ApplicationError{
		Message: "Balance error",
		StatusCode: http.StatusNotFound,
		Code: "not_found",
	}
}