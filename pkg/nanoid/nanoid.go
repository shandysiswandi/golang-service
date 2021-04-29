package nanoid

import gonanoid "github.com/matoous/go-nanoid/v2"

func Generate(l ...int) (string, error) {
	return gonanoid.New(l...)
}
