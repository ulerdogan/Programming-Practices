package services

import (
	"mvc_balance/utils"
	"mvc_balance/domain"
	"github.com/ethereum/go-ethereum/common"
)

type usersService struct{}

var  UsersService usersService

func (u *usersService) GetBalance(addr common.Address) (*domain.Balance, *utils.ApplicationError) {
	balance, err := domain.BalanceDao.GetBalance(addr)
	if err != nil {
		return nil, err
	}
	return balance, nil
}