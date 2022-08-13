package app

import (
	"market_store_db/controllers/status"
)

func mapUrls() {
	router.GET("/status", status.Status)
}