package applicationandroid

import (
	"Angel/src/android/domainandroid"
	"Angel/src/android/domainandroid/entities"
)

type UpdateAndroid struct {
	db domainandroid.IAndroid
}

func NewUpdateAndroid(db domainandroid.IAndroid) *UpdateAndroid {
	return &UpdateAndroid{db: db}
}

func (ua *UpdateAndroid) Execute(android entities.Android) error {
	return ua.db.Update(android)
}
