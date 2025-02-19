package application

import (
	"Angel/src/ios/domain"
	"Angel/src/ios/domain/entities"
)

type CreateIos struct {
	db domain.IIos
}

func NewCreateIos(db domain.IIos) *CreateIos {
	return &CreateIos{db: db}
}

func (ca *CreateIos) Execute(ios entities.Ios) error {
	return ca.db.Save(ios)
}
