package applicationandroid

import (
	"Angel/src/android/domainandroid"
	"Angel/src/android/domainandroid/entities"
)

type ListAndroid struct {
	db domainandroid.IAndroid
}

func NewListAndroid(db domainandroid.IAndroid) *ListAndroid {
	return &ListAndroid{db: db}
}

func (la *ListAndroid) Execute() ([]entities.Android, error) {
	return la.db.GetAll()
}
