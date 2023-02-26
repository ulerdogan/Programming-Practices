package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (server *Server) ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "runs")
}
