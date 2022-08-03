package services

import (
	"nft_api/domain"
	"nft_api/utils/errors"
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
	
}
