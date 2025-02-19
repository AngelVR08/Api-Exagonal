package application

import (
	"Angel/src/ios/domain"
	"Angel/src/ios/domain/entities"
)

type UpdateIos struct {
	db domain.IIos
}

func NewUpdateIos(db domain.IIos) *UpdateIos {
	return &UpdateIos{db: db}
}

func (ua *UpdateIos) Execute(ios entities.Ios) error {
	return ua.db.Update(ios)
}
