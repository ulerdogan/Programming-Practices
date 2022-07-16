package app

import (
	"mvc_balance/controllers"
)

func mapUrls() {
	router.GET("/address/:addr", controllers.GetAddress)
}
