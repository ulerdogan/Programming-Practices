package ethereum

import (
	"nft_api/domain/ethereum/contracts/nft_interface_contract"
	"nft_api/domain"
	"nft_api/utils/errors"
)

type nftBalance struct {}

var NftBalance nftBalance

func (n *nftBalance) BalanceOf(NQ domain.NftQuery) (*domain.NftCheck, errors.ApiError) {
	instance, err := erc721.NewErc721(NQ.NftAddr, Client)
    if err != nil {
        return nil, errors.NewInternalServerError(err.Error())
    }

	balance, err := instance.BalanceOf(nil, NQ.UserAddr)
	if err != nil {
        return nil, errors.NewInternalServerError(err.Error())
    }

	check := &domain.NftCheck{UserAddr: NQ.UserAddr, NftAddr: NQ.NftAddr, Balance: balance.String()}

	return check, nil
}