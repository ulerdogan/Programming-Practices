package api

import (
	"basketball_db/db/sqlc"
	"basketball_db/utils"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createCoachRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,alphanum"`
}

func (server *Server) createAccount(ctx *gin.Context) {
	var req createCoachRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateCoachParams{
		Username:       req.Username,
		HashedPassword: utils.HashPassword(req.Password),
	}

	coach, err := server.store.CreateCoach(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, coach)
}

type getAccountRequest struct {
	ID int64 `uri:"id" binding:"required"`
}

func (server *Server) getAccount(ctx *gin.Context) {
	var req getAccountRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	coach, err := server.store.GetCoachById(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, coach)
}

type loginCoachRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,alphanum"`
}

type loginCoachResponse struct {
	AccessToken string `json:"access_token"`
	Username    string `json:"username"`
}

func (server *Server) loginUser(ctx *gin.Context) {
	var req loginCoachRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	coach, err := server.store.GetCoachByUsername(ctx, req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = utils.CheckPassword(req.Password, coach.HashedPassword)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	accesToken, err := server.token.CreateToken(
		coach.Username,
		server.config.AccessTokenDuration,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := loginCoachResponse{
		AccessToken: accesToken,
		Username:    coach.Username,
	}

	ctx.JSON(http.StatusOK, rsp)
}
