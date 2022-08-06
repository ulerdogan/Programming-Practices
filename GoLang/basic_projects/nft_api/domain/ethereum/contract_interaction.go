package ethereum

import (
	"fmt"
	"log"
	"nft_api/domain"
	"nft_api/domain/ethereum/contracts/nft_minter_contract"
	"nft_api/utils/errors"
	"sync"

	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type nftMinter struct{}

var NftMinter nftMinter
var Instance *igotest.Igotest
var accData *domain.AccData

const CONTRACT_ADDRESS string = "0x57da8bBB52c705b42D700EA68125077845b9637d"

func init() {
	var err error
	Instance, err = igotest.NewIgotest(common.HexToAddress(CONTRACT_ADDRESS), Client)
	if err != nil {
		log.Fatalln(err)
	}
}

func (n *nftMinter) MintNft(NR domain.NftRequest) (*domain.NftsMinted, errors.ApiError) {
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

	_, err = Instance.Mint(auth, NR.ReceiverAddr)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}

	return &domain.NftsMinted{TokenIds: []string{"1"}}, nil
}

func (n *nftMinter) MintNfts(NR domain.NftRequest) (*domain.NftsMinted, errors.ApiError) {
	if accData == nil {
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

		accData = &domain.AccData{
			Nonce:    nonce,
			GasPrice: gasPrice,
			UserAuth: auth,
		}
	}

	var wg sync.WaitGroup
	s := make([]string, 0)

	for i := 0; i < int(NR.Amount); i++ {
		wg.Add(1)
		go n.MintNftConc(NR.ReceiverAddr, uint(i), &wg, &s)
	}

	wg.Wait()

	accData = nil
	return &domain.NftsMinted{TokenIds: s}, nil
}

func (n *nftMinter) MintNftConc(receiver common.Address, i uint, wg *sync.WaitGroup, s *[]string) {
	_, err := Instance.Mint(accData.UserAuth, receiver)
	if err == nil {
		*s = append(*s, fmt.Sprintf("%d", i+1))
	}

	wg.Done()
}
