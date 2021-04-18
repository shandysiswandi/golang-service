package uuid

import lib "github.com/google/uuid"

type UUID interface {
	Generate() string
}

type uuid struct{}

func New() UUID {
	return &uuid{}
}

func (uuid) Generate() string {
	return lib.New().String()
}
