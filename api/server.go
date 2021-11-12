package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	db "github.com/sakhaei-wd/banker/db/sqlc"
)

//This Server will serves all HTTP requests for our banking service
type Server struct {
	store  db.Store    //It will allow us to interact with the database when processing API requests from clients.
	router *gin.Engine //This router will help us send each API request to the correct handler for processing.
}

//This function will create a new Server instance, and setup all HTTP API routes for our service on that server.
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default() //we create a new router by  calling gin.Default()

	//get the current validator engine that Gin is using (binding is a sub-package of Gin)
	//we have to convert the output to a validator.Validate object pointer. If it is ok then we can call v.RegisterValidation() to register our custom validate function.
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
		//now we have custom currency tag
	}

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)

	router.POST("/entries", server.createEntry)
	router.GET("/entries/:id", server.getEntry)
	router.GET("/entries", server.listEntry)

	router.POST("/transfers", server.createTransfer)

	router.GET("/user/:username", server.getUser)
	router.POST("/user", server.createUser)

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
