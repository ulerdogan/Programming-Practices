package app

import (
	"nft_api/controllers/status"
	"nft_api/controllers/nfts"
)

func mapUrls() {
	router.GET("/", status.Status)
	router.GET("/nft", nfts.GetNfts)
}
