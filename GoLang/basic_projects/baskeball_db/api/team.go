package api

import (
	"basketball_db/db/sqlc"
	"basketball_db/token"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) createTeam(ctx *gin.Context) {
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	coach, err := server.store.GetCoachByUsername(ctx, authPayload.Username)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	team, err := server.store.CreateTeam(ctx, coach.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, team)
}

type getTeamRequest struct {
	ID int64 `uri:"id" binding:"required"`
}

func (server *Server) getTeam(ctx *gin.Context) {
	var req getTeamRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	team, err := server.store.GetTeam(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, team)
}

type IncreaseWinRequest struct {
	Amount int64 `json:"amount" binding:"required"`
}

func (server *Server) IncreaseWin(ctx *gin.Context) {
	var req IncreaseWinRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	coach, err := server.store.GetCoachByUsername(ctx, authPayload.Username)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	team, err := server.store.GetTeamByCoach(ctx, coach.ID)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	arg := db.IncreaseWinTeamParams{
		ID:   team.ID,
		Wins: int32(req.Amount),
	}

	team, err = server.store.IncreaseWinTeam(ctx, arg)
	ctx.JSON(http.StatusOK, team)
}
