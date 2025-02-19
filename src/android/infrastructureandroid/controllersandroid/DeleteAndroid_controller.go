package controllersandroid

import (
	"Angel/src/android/applicationandroid"
	"github.com/gin-gonic/gin"
	"strconv"
)

type DeleteAndroidController struct {
	useCase *applicationandroid.DeleteAndroid
}

func NewDeleteAndroidController(useCase *applicationandroid.DeleteAndroid) *DeleteAndroidController {
	return &DeleteAndroidController{useCase: useCase}
}

func (da_c *DeleteAndroidController) Execute(c *gin.Context) {
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
	c.JSON(200, gin.H{"message": "celular eliminado correctamente"})
}
