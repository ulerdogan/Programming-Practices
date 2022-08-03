package domain

import "github.com/ethereum/go-ethereum/common"

type NftQuery struct {
	UserAddr common.Address
	NftAddr  common.Address
}

type NftCheck struct {
	UserAddr common.Address
	NftAddr  common.Address
	Balance string
}