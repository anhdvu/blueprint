package repository

import "github.com/anhdvu/blueprint/internal/model"

type ThingStore interface {
	Create(t *model.Thing) error
	Read(id int64) (model.Thing, error)
	Update(t *model.Thing) error
	Delete(id int64) error
	All() ([]model.Thing, error)
}

type Repositories struct {
	Things ThingStore
}

func New(ts ThingStore) *Repositories {
	return &Repositories{
		Things: ts,
	}
}
