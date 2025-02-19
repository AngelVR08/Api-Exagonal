package domain

import (
	"Angel/src/ios/domain/entities"
)

type IIos interface {
	Save(ios entities.Ios) error
	GetAll() ([]entities.Ios, error)
	Update(ios entities.Ios) error
	Delete(id int) error
}
