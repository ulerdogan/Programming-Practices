package app

import (
	"market_store_db/controllers/nfts"
	"market_store_db/controllers/status"
	"net/http"

	"github.com/gin-gonic/gin"
)

func mapUrls() {
	notImplemented := func(c *gin.Context) { c.String(http.StatusOK, "Not implemented yet.") }
	_ = notImplemented
	
	router.GET("/status", status.Status)
	router.GET("/nfts/:id", nfts.Get)

	router.POST("/nfts", nfts.Create)
}
