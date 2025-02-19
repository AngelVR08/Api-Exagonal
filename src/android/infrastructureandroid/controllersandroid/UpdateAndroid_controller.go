package controllersandroid

import (
	"Angel/src/android/applicationandroid"
	"Angel/src/android/domainandroid/entities"
	"github.com/gin-gonic/gin"
	"strconv"
)

type UpdateAndroidController struct {
	useCase *applicationandroid.UpdateAndroid
}

func NewUpdateAndroidController(useCase *applicationandroid.UpdateAndroid) *UpdateAndroidController {
	return &UpdateAndroidController{useCase: useCase}
}

func (us_c *UpdateAndroidController) Execute(c *gin.Context) {
	var android entities.Android
	if err := c.ShouldBindJSON(&android); err != nil {
		c.JSON(400, gin.H{"error": "Datos invalidos"})
		return
	}
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "ID no encontrado"})
		return
	}
	android.ID = idInt
	if err := us_c.useCase.Execute(android); err != nil {
		c.JSON(500, gin.H{"error": "No se pudo actualizar el celular"})
	}
	c.JSON(200, gin.H{"message": "Celular actualizado"})
}
