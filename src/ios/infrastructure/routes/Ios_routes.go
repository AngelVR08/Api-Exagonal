package routes

import (
	"Angel/src/ios/infrastructure/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigureRoutesIos(
	r *gin.Engine,
	createIosController *controllers.CreateIosController,
	listIosController *controllers.ListIosController,
	updateIosController *controllers.UpdateIosController,
	deleteIosController *controllers.DeleteIosController,
) {
	r.GET("/ios", listIosController.Execute)
	r.POST("/ios", createIosController.Execute)
	r.PUT("/ios/:id", updateIosController.Execute)
	r.DELETE("/ios/:id", deleteIosController.Execute)
}
