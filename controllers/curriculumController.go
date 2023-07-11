package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	categoryService string
}

// declare constructor
func NewCategoryController() *CategoryController {
	return &CategoryController{
		categoryService: "Hello",
	}
}

// method
func (categoryController CategoryController) GetListCategory(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "Hello Gin Framework")
}
