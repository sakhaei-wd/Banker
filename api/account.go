package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	db "github.com/sakhaei-wd/banker/db/sqlc"
)

type createAccountRequest struct {
	Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required,currency"`  //use currency custom tag instead of : oneof=USD EUR"` which we can use the oneof condition to check only between USD and EUR
}

type getAccountRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

//we’re not getting these parameters from uri, but from query string instead,
//so we cannot use the uri tag. Instead, we should use form tag
type listAccountRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
	//for all validation : https://pkg.go.dev/github.com/go-playground/validator
}

func (server *Server) createAccount(ctx *gin.Context) {
	var req createAccountRequest

	//This function will return an error
	//If the error is not nil, then it means that the client has provided invalid data
	//So we should send a 400 Bad Request response to the client
	//To do that, we just call ctx.JSON() function to send a JSON response.
	if err := ctx.ShouldBindJSON(&req); err != nil {

		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		//The second argument is the JSON object that we want to send to the client. Here we just want to send the error, so we will need a function to convert this error into a key-value object so that Gin can serialize it to JSON before returning to the client.
		return
	}

	arg := db.CreateAccountParams{
		Owner:    req.Owner,
		Currency: req.Currency,
		Balance:  0,
	}

	account, err := server.store.CreateAccount(ctx, arg)
	if err != nil {
		//if an error is returned, we will try to convert it to pq.Error type, and assign the result to pqErr variable
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "foreign_key_violation" , "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
                return
			}
            log.Println(pqErr.Code.Name())
        }
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

func (server *Server) getAccount(ctx *gin.Context) {
	var req getAccountRequest

	// here instead of ShouldBindJSON, we should call ShouldBindUri
	if err := ctx.ShouldBindUri(&req); err != nil {

		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	
	account, err := server.store.GetAccount(ctx, req.ID)
	if err != nil {

		//the account with that specific input ID doesn’t exist
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		//some internal error when querying data from the database.
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

func (server *Server) listAccount(ctx *gin.Context) {
	var req listAccountRequest

	// tell Gin to get data from query string
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListAccountsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	accounts, err := server.store.ListAccounts(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, accounts)
}
