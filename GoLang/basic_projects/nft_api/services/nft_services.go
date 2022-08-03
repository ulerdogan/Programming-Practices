package services

import (
	"nft_api/domain"
	"nft_api/utils/errors"
	"nft_api/domain/ethereum"
)

type nftService struct{}

type reposServiceInterface interface {
	GetBalance(NQ domain.NftQuery) (*domain.NftCheck, errors.ApiError)
}

var  NftService reposServiceInterface

func init() {
	NftService = &nftService{}
}

func (nq *nftService) GetBalance(NQ domain.NftQuery) (*domain.NftCheck, errors.ApiError) {
	request, err := ethereum.NftBalance.BalanceOf(NQ)
	if err != nil {
		return nil, errors.NewInternalServerError("Balance query errror.")
	}
	return request, nil
}
