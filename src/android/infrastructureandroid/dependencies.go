package infrastructureandroid

import (
	"Angel/src/android/applicationandroid"
	"Angel/src/android/infrastructureandroid/controllersandroid"
	"Angel/src/android/infrastructureandroid/databaseandroid"
)

type DependenciesAndroid struct {
	CreateAndroidController *controllersandroid.CreateAndroidController
	ListAndroidController   *controllersandroid.ListAndroidController
	UpdateAndroidController *controllersandroid.UpdateAndroidController
	DeleteAndroidController *controllersandroid.DeleteAndroidController
}

func InitAndroid() *DependenciesAndroid {
	ps := databaseandroid.NewMySQL()

	return &DependenciesAndroid{
		CreateAndroidController: controllersandroid.NewCreateAndroidController(applicationandroid.NewCreateAndroid(ps)),
		ListAndroidController:   controllersandroid.NewListAndroidController(applicationandroid.NewListAndroid(ps)),
		UpdateAndroidController: controllersandroid.NewUpdateAndroidController(applicationandroid.NewUpdateAndroid(ps)),
		DeleteAndroidController: controllersandroid.NewDeleteAndroidController(applicationandroid.NewDeleteAndroid(ps)),
	}
}
