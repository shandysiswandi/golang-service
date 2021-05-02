package nanoid

import gonanoid "github.com/matoous/go-nanoid/v2"

type IDGenerator interface {
	Generate() string
}

type nanoid struct{}

func New() IDGenerator {
	return &nanoid{}
}

func (*nanoid) Generate() string {
	id, err := gonanoid.New(11)
	if err != nil {
		panic("can't generate nanoid")
	}
	return id
}
