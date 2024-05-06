package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/leondavidzipp/simple_bank/db/sqlc"
)


// server for serving http requests to our banking service
type Server struct {
	store *db.Store
	router *gin.Engine
}

// creates new server instance
func NewServer(store *db.Store) *Server {
	server := &Server{
		store : store
	}

	router := gin.Default()
	router.POST("/accounts", server.createAccount)

	server.router = router

	return server
}
