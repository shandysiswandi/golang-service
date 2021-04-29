package env

import (
	"os"

	"github.com/joho/godotenv"
)

func Load(path string) error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	dir = dir + "/../../"

	if err := godotenv.Load(dir + path); err != nil {
		return err
	}
	return nil
}
