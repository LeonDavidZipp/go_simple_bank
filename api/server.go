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

// Starts http server on specified address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

// formats error response into json body
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
