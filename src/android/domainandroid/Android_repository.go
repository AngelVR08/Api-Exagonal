package domainandroid

import (
	"Angel/src/android/domainandroid/entities"
)

type IAndroid interface {
	Save(android entities.Android) error
	GetAll() ([]entities.Android, error)
	Update(android entities.Android) error
	Delete(id int) error
}
