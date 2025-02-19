package application

import (
	"Angel/src/ios/domain"
)

type DeleteIos struct {
	db domain.IIos
}

func NewDeleteIos(db domain.IIos) *DeleteIos {
	return &DeleteIos{db: db}
}

func (da *DeleteIos) Execute(id int) error {
	return da.db.Delete(id)
}
