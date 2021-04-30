package uuid

import lib "github.com/google/uuid"

func Generate() string {
	return lib.New().String()
}
