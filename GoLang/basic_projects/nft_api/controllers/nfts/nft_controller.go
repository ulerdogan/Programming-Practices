package nfts

import (
	"nft_api/domain"
	"nft_api/services"
	eth_u "nft_api/utils/ethereum_utils"
	"nft_api/utils/errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetNfts(c *gin.Context) {
	nftAddr := common.HexToAddress(c.Param("nft"))
	userAddr := common.HexToAddress(c.Param("user"))

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
