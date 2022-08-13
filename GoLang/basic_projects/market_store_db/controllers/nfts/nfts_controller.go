package nfts

import (
	"market_store_db/utils/rest_errors"
	"market_store_db/domain/nfts"
	"market_store_db/services"
	
	"net/http"
	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {

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

}

func Update(c *gin.Context) {

}

func Index(c *gin.Context) {

}