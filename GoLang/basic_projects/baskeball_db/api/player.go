package api

import (
	"basketball_db/db/sqlc"
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

type createPlayerRequest struct {
	Name string `json:"name" binding:"required"`
	Role string `json:"role" binding:"required"`
}

func (server *Server) createPlayer(ctx *gin.Context) {
	var req createPlayerRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreatePlayerParams{
		Name: req.Name,
		Role: db.Roles(req.Role),
	}

	player, err := server.store.CreatePlayer(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, player)
}

type getPlayerRequest struct {
	ID int64 `uri:"id" binding:"required"`
}

func (server *Server) getPlayer(ctx *gin.Context) {
	var req getPlayerRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	player, err := server.store.GetPlayer(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, player)
}

type listTeamsPlayersRequest struct {
	Team     int64 `form:"team" binding:"required"`
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listTeamsPlayers(ctx *gin.Context) {
	var req listTeamsPlayersRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.GetTeamsPlayersParams{
		Team:   sql.NullInt64{Int64: req.Team, Valid: true},
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	players, err := server.store.GetTeamsPlayers(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, players)
}
