package controllers

import (
	"learning-freemarket-api/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type IItemController interface {
	FindAll(ctx *gin.Context)
	FindByID(ctx *gin.Context)
}

// 具体的な実装
type ItemController struct {
	service services.IItemService
}

// 3層アーキテクチャの別の層では、具体的な実装（services.ItemService）を参照しない
// インターフェース（services.IItemService）に依存する作りにする
func NewItemController(service services.IItemService) IItemController {
	return &ItemController{
		service: service,
	}
}

func (c *ItemController) FindAll(ctx *gin.Context) {
	items, err := c.service.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": items,
	})
}

func (c *ItemController) FindByID(ctx *gin.Context) {
	itemId, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Id"})
		return
	}
	item, err := c.service.FindByID(uint(itemId))
	if err != nil {
		if err.Error() == "Item not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": item,
	})
}
