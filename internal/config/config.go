package config

import "os"

type Config struct {
	Host      string
	Port      string
	TZ        string
	ENV       string
	Image     string
	Mongos    []Mongo
	JWTSecret string
}

type Mongo struct {
	URI            string
	ConnectionName string
	DatabaseName   string
}

func New() *Config {
	config := &Config{}

	config.Host = os.Getenv("HOST")
	config.Port = os.Getenv("PORT")
	config.TZ = os.Getenv("TZ")
	config.ENV = os.Getenv("ENV")
	config.Image = os.Getenv("IMAGE")

	config.JWTSecret = os.Getenv("JWT_SECRET")

	config.Mongos = append(config.Mongos, Mongo{
		URI:            os.Getenv("MONGO_URI"),
		ConnectionName: os.Getenv("MONGO_CONNECTION_NAME"),
		DatabaseName:   os.Getenv("MONGO_DATABASE_NAME"),
	})

	return config
}
