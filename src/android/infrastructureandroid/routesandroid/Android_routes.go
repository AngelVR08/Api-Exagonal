package routesandroid

import (
	"Angel/src/android/infrastructureandroid/controllersandroid"
	"github.com/gin-gonic/gin"
)

func ConfigureRoutesAndroid(
	r *gin.Engine,
	createAndroidController *controllersandroid.CreateAndroidController,
	listAndroidController *controllersandroid.ListAndroidController,
	updateAndroidController *controllersandroid.UpdateAndroidController,
	deleteAndroidController *controllersandroid.DeleteAndroidController,
) {
	r.GET("/android", listAndroidController.Execute)
	r.POST("/android", createAndroidController.Execute)
	r.PUT("/android/:id", updateAndroidController.Execute)
	r.DELETE("/android/:id", deleteAndroidController.Execute)
}
