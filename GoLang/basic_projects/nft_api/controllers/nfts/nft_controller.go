package nfts

import (
	"fmt"
	"nft_api/domain"
	"nft_api/services"
	"nft_api/utils/errors"
	eth_u "nft_api/utils/ethereum_utils"

	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

func GetNfts(c *gin.Context) {
	nftAddr := common.HexToAddress(c.Param("nft"))
	userAddr := common.HexToAddress(c.Param("usr"))

	fmt.Println("aaaa", nftAddr, userAddr)

	if !(eth_u.IsValidAddress(nftAddr) && eth_u.IsValidAddress(userAddr)) {
		apiErr := errors.NewBadRequestError("Invalid addresses input")
		c.JSON(apiErr.Status(), apiErr)
		return
	}
	
	check := domain.NftQuery{NftAddr: nftAddr, UserAddr: userAddr}

	response, apiErr := services.NftService.GetBalance(check)
	if apiErr != nil {
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	c.JSON(http.StatusCreated, response)
}
