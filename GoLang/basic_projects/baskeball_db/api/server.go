package api

import (
	"basketball_db/db/sqlc"
	"basketball_db/token"
	"basketball_db/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

type Server struct {
	config utils.Config
	store  db.Store
	token  token.Maker
	router *gin.Engine
}

func NewServer(store db.Store, config utils.Config) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker")
	}

	server := &Server{
		config: config,
		store:  store,
		token:  tokenMaker,
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()
	router.GET("/ping", server.ping)
	router.POST("/coaches", server.createAccount)
	router.POST("/coaches/login", server.loginUser)
	
	authRoutes := router.Group("/").Use(authMiddleware(server.token))
	authRoutes.POST("/coaches/team", server.createTeam)
	authRoutes.GET("/coaches", server.getAccount)
	authRoutes.POST("/coaches/win", server.IncreaseWin)

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}