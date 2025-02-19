package main

import (
	androidInfra "Angel/src/android/infrastructureandroid"
	androidRoutes "Angel/src/android/infrastructureandroid/routesandroid"
	iosInfra "Angel/src/ios/infrastructure"
	iosRoutes "Angel/src/ios/infrastructure/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	androidDependencies := androidInfra.InitAndroid()
	androidRoutes.ConfigureRoutesAndroid(r,
		androidDependencies.CreateAndroidController,
		androidDependencies.ListAndroidController,
		androidDependencies.UpdateAndroidController,
		androidDependencies.DeleteAndroidController)

	ioDependencies := iosInfra.InitIos()
	iosRoutes.ConfigureRoutesIos(r,
		ioDependencies.CreateIosController,
		ioDependencies.ListIosController,
		ioDependencies.UpdateIosController,
		ioDependencies.DeleteIosController)

	r.Run(":8080")
}
