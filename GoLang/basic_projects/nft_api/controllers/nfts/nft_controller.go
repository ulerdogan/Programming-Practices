package nfts

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetNfts(c *gin.Context) {
	c.String(http.StatusOK, "nft")
}
