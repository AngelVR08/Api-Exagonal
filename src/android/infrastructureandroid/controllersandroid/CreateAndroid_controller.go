package controllersandroid

import (
	"Angel/src/android/applicationandroid"
	"Angel/src/android/domainandroid/entities"
	"github.com/gin-gonic/gin"
)

type CreateAndroidController struct {
	useCase *applicationandroid.CreateAndroid
}

func NewCreateAndroidController(useCase *applicationandroid.CreateAndroid) *CreateAndroidController {
	return &CreateAndroidController{useCase: useCase}
}

func (cs_a *CreateAndroidController) Execute(c *gin.Context) {
	var android entities.Android
	if err := c.ShouldBindJSON(&android); err != nil {
		c.JSON(400, gin.H{"error": "Datos invalidos"})
	}
	if err := cs_a.useCase.Execute(android); err != nil {
		c.JSON(500, gin.H{"error": "No se pudo crear el telefono"})
	}
	c.JSON(201, gin.H{"message": "Creado con exito!"})
}
