package controllersandroid

import (
	"Angel/src/android/applicationandroid"
	"github.com/gin-gonic/gin"
)

type ListAndroidController struct {
	useCase *applicationandroid.ListAndroid
}

func NewListAndroidController(useCase *applicationandroid.ListAndroid) *ListAndroidController {
	return &ListAndroidController{useCase: useCase}
}

func (la_c *ListAndroidController) Execute(c *gin.Context) {
	android, err := la_c.useCase.Execute()
	if err != nil {
		c.JSON(500, gin.H{"error": "Error al obtener datos de los celulares"})
		return
	}
	c.JSON(200, android)
}
