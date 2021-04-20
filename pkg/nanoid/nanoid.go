package nanoid

import gonanoid "github.com/matoous/go-nanoid/v2"

func Generate() (string, error) {
	return gonanoid.New()
}
