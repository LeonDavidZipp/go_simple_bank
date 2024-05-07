package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/leondavidzipp/simple_bank/db/sqlc"
)


type createAccountRequest struct {
	Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required,oneof=USD EUR CAD"`
}

func (server *Server) createAccount(ctx *gin.Context) {
	var req createAccountRequest

	if err := ctx.ShouldBindJson(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return 
	}

	arg := db.CreateAccountParams{
		Owner : req.Owner,
		Currency : req.Currency,
		Balance : int64(0),
	}

	account, err := server.store.CreateAccount(ctx, arg)
	if err != nil {
		ctx.JSON(
			http.StatusInternelServerError, errorResponse(err)
		)
	}

	ctx.JSON(http.StatusOK, account)
}
