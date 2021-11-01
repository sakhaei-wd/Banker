package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/sakhaei-wd/banker/db/sqlc"
)

//This Server will serves all HTTP requests for our banking service
type Server struct {
	store  *db.Store   //It will allow us to interact with the database when processing API requests from clients.
	router *gin.Engine //This router will help us send each API request to the correct handler for processing.
}

//This function will create a new Server instance, and setup all HTTP API routes for our service on that server.
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default() //we create a new router by  calling gin.Default()

	// TODO: add routes to router
	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)
	
	server.router = router
	return server
}

//This function will take an error as input, and it will return a gin.H object
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

//Its role is to run the HTTP server on the input address to start listening for API requests
//Note that the server.router field is private, so it cannot be accessed from outside of this api package. That’s one of the reasons we have this public Start() function.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
