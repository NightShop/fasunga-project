package api

import (
	"database/sql"
	"net/http"

	db "github.com/NightShop/fasunga-project/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createItemRequest struct {
	UserEmail   string `json:"user_email" binding:"required,email"`
	Description string `json:"description" binding:"required,min=1"`
	GroupKey    string `json:"group_key" binding:"required,min=1"`
}

func (server *Server) createItem(ctx *gin.Context) {
	var req createItemRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateItemParams{
		UserEmail:   req.UserEmail,
		GroupKey:    req.GroupKey,
		Description: req.Description,
	}

	item, err := server.store.CreateItem(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, item)
}

type listItemRequest struct {
	GroupKey string `uri:"group_key" binding:"required,min=1"`
}

func (server *Server) listItems(ctx *gin.Context) {
	var req listItemRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	items, err := server.store.ListItems(ctx, req.GroupKey)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, items)
}

type updateItemRequest struct {
	ID      int64 `json:"id" binding:"required,min=1"`
	Checked bool  `json:"checked" binding:"required"`
}

func (server *Server) updateItem(ctx *gin.Context) {
	var req updateItemRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateItemParams{
		ID:      req.ID,
		Checked: req.Checked,
	}

	item, err := server.store.UpdateItem(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, item)
}
