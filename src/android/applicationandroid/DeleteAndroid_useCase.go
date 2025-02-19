package applicationandroid

import (
	"Angel/src/android/domainandroid"
)

type DeleteAndroid struct {
	db domainandroid.IAndroid
}

func NewDeleteAndroid(db domainandroid.IAndroid) *DeleteAndroid {
	return &DeleteAndroid{db: db}
}

func (da *DeleteAndroid) Execute(id int) error {
	return da.db.Delete(id)
}
