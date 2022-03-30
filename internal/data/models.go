package data

type Thing struct {
	ID int64
}

type ThingStore interface {
	Create(t *Thing) error
	Read(id int64) (*Thing, error)
	Update(t *Thing) error
	Delete(id int64) error
}

type Repositories struct {
	Things ThingStore
}

func New() *Repositories {
	return &Repositories{
		Things: NewPGStore(),
	}
}
