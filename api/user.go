package api

import (
	"database/sql"
	"net/http"

	db "github.com/NightShop/fasunga-project/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createUserRequest struct {
	Email          string `json:"email" binding:"required,email"`
	GroupKey       string `json:"group_key" binding:"required,min=1"`
	HashedPassword string `json:"hashed_password" binging:"required, min=1"`
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateUserParams{
		Email:          req.Email,
		GroupKey:       req.GroupKey,
		HashedPassword: req.HashedPassword,
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, user)
}

type getUserRequest struct {
	Email string `json:"email" binding:"required,email"`
}

func (server *Server) getUser(ctx *gin.Context) {
	var req getUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.store.GetUser(ctx, req.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, user)
}
