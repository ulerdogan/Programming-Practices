package domain

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

type User struct {
    FromAddress common.Address
    Auth *bind.TransactOpts
}
