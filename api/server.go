package api

import (
	db "github.com/NightShop/fasunga-project/db/sqlc"
	"github.com/gin-gonic/gin"
)

// Server serves HTTP requests for our banking service
type Server struct {
	store  *db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/items", server.createItem)
	router.GET("/items/:group_key", server.listItems)
	router.PATCH("/items", server.updateItem)
	router.POST("/users", server.createUser)
	router.GET("/users", server.getUser)

	server.router = router

	return server
}

// Start runs the http server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
