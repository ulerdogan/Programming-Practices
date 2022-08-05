package ethereum

import (
	"crypto/ecdsa"
	"log"
	"nft_api/domain"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
)

var User domain.User

const PRIVATE_KEY string = "PRIVATE_KEY"
const CHAIN_ID int64 = 43113

func init() {

	privateKey, err := crypto.HexToECDSA(PRIVATE_KEY)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()

	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(CHAIN_ID))
	if err != nil {
		log.Fatal(err)
	}

	User = domain.User{
		FromAddress: fromAddress,
		Auth:        auth,
	}
}
