package gen

import (
	libuuid "github.com/google/uuid"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

type (
	Generator interface {
		Nanoid() string
		UUID() string
	}

	gen struct{}
)

func New() Generator {
	return &gen{}
}

func (*gen) Nanoid() string {
	return gonanoid.Must(11)
}

func (*gen) UUID() string {
	return libuuid.NewString()
}
