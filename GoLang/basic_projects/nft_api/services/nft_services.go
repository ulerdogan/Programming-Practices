package services

import (
	"nft_api/domain"
	"nft_api/domain/ethereum"
	"nft_api/utils/errors"
)

type nftService struct{}

type reposServiceInterface interface {
	GetBalance(NQ domain.NftQuery) (*domain.NftCheck, errors.ApiError)
	MintNfts(R domain.NftRequest) (*domain.NftsMinted, errors.ApiError)
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

func (nq *nftService) MintNfts(R domain.NftRequest) (*domain.NftsMinted, errors.ApiError) {
	result, err := ethereum.NftMinter.MintNft(R)
	if err != nil {
		return nil, errors.NewInternalServerError("Minting error.")
	}

	return result, nil
}