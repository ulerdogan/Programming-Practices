package nfts

import (
	"market_store_db/domain/nfts"
	"market_store_db/services"
	"market_store_db/utils/rest_errors"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
)


func Get(c *gin.Context) {
	id, idErr := checkId(c.Param("id"))
	if idErr != nil {
		c.JSON(idErr.Status(), idErr)
		return
	}

	nft, getErr := services.NftsService.GetListing(id)
	if getErr != nil {
		c.JSON(getErr.Status(), getErr)
		return
	}

	c.JSON(http.StatusOK, nft)
}

func checkId(id string) (int64, rest_errors.RestErr) {
	nId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return 0, rest_errors.NewBadRequestError("invalid id")
	}
	return nId, nil
}

func Create(c *gin.Context) {
	var listing nfts.Listing
	if err := c.ShouldBindJSON(&listing); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}

	result, saveErr := services.NftsService.CreateListing(listing)
	if saveErr != nil {
		c.JSON(saveErr.Status(), saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func Delete(c *gin.Context) {
	id, idErr := checkId(c.Param("id"))
	if idErr != nil {
		c.JSON(idErr.Status(), idErr)
		return
	}

	if err := services.NftsService.DeleteListing(id); err != nil {
		c.JSON(err.Status(), err)
		return
	}
}

func Update(c *gin.Context) {
	id, idErr := checkId(c.Param("id"))
	if idErr != nil {
		c.JSON(idErr.Status(), idErr)
		return
	}

	var listing nfts.Listing
	if err := c.ShouldBindJSON(&listing); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}

	listing.Id = id

	isPartial := c.Request.Method == http.MethodPatch
	result, err := services.NftsService.UpdateListing(isPartial, listing)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, result)

}

func Index(c *gin.Context) {

}