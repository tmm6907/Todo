package handlers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func (h Handler) GetTodoItems(ctx *gin.Context) {

}
