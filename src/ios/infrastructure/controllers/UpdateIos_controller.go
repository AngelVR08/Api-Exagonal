package controllers

import (
	"Angel/src/ios/application"
	"Angel/src/ios/domain/entities"
	"github.com/gin-gonic/gin"
	"strconv"
)

type UpdateIosController struct {
	useCase *application.UpdateIos
}

func NewUpdateIosController(useCase *application.UpdateIos) *UpdateIosController {
	return &UpdateIosController{useCase: useCase}
}

func (us_c *UpdateIosController) Execute(c *gin.Context) {
	var ios entities.Ios
	if err := c.ShouldBind(&ios); err != nil {
		c.JSON(400, gin.H{"error": "Datos invalidos"})
		return
	}
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "ID no encontrado"})
		return
	}
	ios.ID = idInt
	if err := us_c.useCase.Execute(ios); err != nil {
		c.JSON(500, gin.H{"error": "No se pudo actualizar el celular"})
		return
	}
	c.JSON(200, gin.H{"message": "Celular actualizado"})
}
