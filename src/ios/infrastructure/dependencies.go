package infrastructure

import (
	"Angel/src/ios/application"
	"Angel/src/ios/infrastructure/controllers"
	"Angel/src/ios/infrastructure/database"
)

type DependenciesIos struct {
	CreateIosController *controllers.CreateIosController
	ListIosController   *controllers.ListIosController
	UpdateIosController *controllers.UpdateIosController
	DeleteIosController *controllers.DeleteIosController
}

func InitIos() *DependenciesIos {
	ps := database.NewMySQL()

	return &DependenciesIos{
		CreateIosController: controllers.NewCreateIosController(application.NewCreateIos(ps)),
		ListIosController:   controllers.NewListIosController(application.NewListIos(ps)),
		UpdateIosController: controllers.NewUpdateIosController(application.NewUpdateIos(ps)),
		DeleteIosController: controllers.NewDeleteIoSController(application.NewDeleteIos(ps)),
	}
}
