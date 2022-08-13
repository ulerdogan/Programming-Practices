package status

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Status(c *gin.Context) {
	c.String(http.StatusOK, "Server is Running.")
}
