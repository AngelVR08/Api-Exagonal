package application

import (
	"Angel/src/ios/domain"
	"Angel/src/ios/domain/entities"
)

type ListIos struct {
	db domain.IIos
}

func NewListIos(db domain.IIos) *ListIos {
	return &ListIos{db: db}
}

func (la *ListIos) Execute() ([]entities.Ios, error) {
	return la.db.GetAll()
}
