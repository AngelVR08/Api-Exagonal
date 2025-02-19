package controllers

import (
	"Angel/src/ios/application"
	"github.com/gin-gonic/gin"
	"strconv"
)

type DeleteIosController struct {
	useCase *application.DeleteIos
}

func NewDeleteIoSController(useCase *application.DeleteIos) *DeleteIosController {
	return &DeleteIosController{useCase: useCase}
}

func (da_c *DeleteIosController) Execute(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "ID invalido"})
		return
	}
	if err := da_c.useCase.Execute(idInt); err != nil {
		c.JSON(500, gin.H{"error": "No se pudo eliminar el telefono"})
		return
	}
	c.JSON(200, gin.H{"message": "Celular eliminado correctamente"})
}
