package controllers

import (
	"mvc_balance/services"
	"mvc_balance/utils"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

func GetAddress(c *gin.Context) {
	addr := c.Param("addr")
	if !utils.IsValidAddress(addr) || utils.IsZeroAddress(addr) {
		utils.RespondError(c, &utils.ApplicationError{
			Message: "Not a valid address",
			StatusCode: http.StatusBadRequest,
			Code: "bad_request",
		})
		return
	}
	
	balance, err := services.UsersService.GetBalance(common.HexToAddress(addr))
	if err != nil {
		utils.RespondError(c, err)
		return
	}

	utils.Respond(c, http.StatusOK, balance)
}