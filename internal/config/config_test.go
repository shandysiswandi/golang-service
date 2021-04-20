package config_test

import (
	"os"
	"testing"

	"github.com/shandysiswandi/echo-service/internal/config"
	"github.com/stretchr/testify/assert"
)

func Test_New(t *testing.T) {
	// setup data
	host := "localhost"
	port := "3000"
	tz := "UTC"
	env := "development"
	image := "echo-service"
	mgouri := "echo-service"
	mgoconanme := "echo-service"
	mgodbname := "echo-service"
	jwtsecret := "fake-secret-jwt"

	// set env variable
	os.Setenv("HOST", host)
	os.Setenv("PORT", port)
	os.Setenv("ENV", env)
	os.Setenv("TZ", tz)
	os.Setenv("IMAGE", image)

	os.Setenv("MONGO_URI", mgouri)
	os.Setenv("MONGO_CONNECTION_NAME", mgoconanme)
	os.Setenv("MONGO_DATABASE_NAME", mgodbname)

	os.Setenv("JWT_SECRET", jwtsecret)

	// testing & assertion
	cfg := config.New()
	assert.NotNil(t, cfg)
	assert.Equal(t, host, cfg.Host)
	assert.Equal(t, port, cfg.Port)
	assert.Equal(t, env, cfg.ENV)
	assert.Equal(t, tz, cfg.TZ)
	assert.Equal(t, image, cfg.Image)

	assert.Equal(t, mgouri, cfg.Mongos[0].URI)
	assert.Equal(t, mgoconanme, cfg.Mongos[0].ConnectionName)
	assert.Equal(t, mgodbname, cfg.Mongos[0].DatabaseName)

	assert.Equal(t, jwtsecret, cfg.JWTSecret)
}
