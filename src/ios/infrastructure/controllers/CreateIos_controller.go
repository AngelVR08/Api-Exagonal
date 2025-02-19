package controllers

import (
	"Angel/src/ios/application"
	"Angel/src/ios/domain/entities"
	"github.com/gin-gonic/gin"
)

type CreateIosController struct {
	useCase *application.CreateIos
}

func NewCreateIosController(useCase *application.CreateIos) *CreateIosController {
	return &CreateIosController{useCase: useCase}
}

func (cs_a *CreateIosController) Execute(c *gin.Context) {
	var ios entities.Ios

	if err := c.ShouldBindJSON(&ios); err != nil {
		c.JSON(400, gin.H{"error": "Datos invalidos"})
		return
	}

	if err := cs_a.useCase.Execute(ios); err != nil {
		c.JSON(500, gin.H{"error": "No se pudo crear el telefono"})
		return
	}

	c.JSON(201, gin.H{"message": "Creado correctamente"})
}
