package ethereum

import (
	"nft_api/domain"
	"nft_api/domain/ethereum/contracts/nft_minter_contract"
	"nft_api/utils/errors"

	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type nftMinter struct{}

var NftMinter nftMinter

const CONTRACT_ADDRESS string = "ADDR"

func (n *nftMinter) MintNft(NR domain.NftRequest) (*domain.NftsMinted, errors.ApiError) {
	instance, err := igotest.NewIgotest(common.HexToAddress(CONTRACT_ADDRESS), Client)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}

	nonce, err := Client.PendingNonceAt(context.Background(), User.FromAddress)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}

	gasPrice, err := Client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}

	auth := User.Auth

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(300000)
	auth.GasPrice = gasPrice

	_, err = instance.Mint(auth, NR.ReceiverAddr)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}

	return &domain.NftsMinted{TokenIds: []string{"lrm", "ipsm"}}, nil
}
