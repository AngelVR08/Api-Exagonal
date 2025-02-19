package applicationandroid

import (
	"Angel/src/android/domainandroid"
	"Angel/src/android/domainandroid/entities"
)

type CreateAndroid struct {
	db domainandroid.IAndroid
}

func NewCreateAndroid(db domainandroid.IAndroid) *CreateAndroid {
	return &CreateAndroid{db: db}
}

func (ca *CreateAndroid) Execute(android entities.Android) error {
	return ca.db.Save(android)
}
