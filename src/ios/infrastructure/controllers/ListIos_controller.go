package controllers

import (
	"Angel/src/ios/application"
	"github.com/gin-gonic/gin"
)

type ListIosController struct {
	useCase *application.ListIos
}

func NewListIosController(useCase *application.ListIos) *ListIosController {
	return &ListIosController{useCase: useCase}
}

func (la_c *ListIosController) Execute(c *gin.Context) {
	ios, err := la_c.useCase.Execute()
	if err != nil {
		c.JSON(500, gin.H{"error": "Error al obtener los celulares"})
		return
	}
	c.JSON(200, ios)
}
