package mongodb_test

import (
	"testing"

	"github.com/shandysiswandi/echo-service/internal/config"
	"github.com/shandysiswandi/echo-service/internal/infrastructure/mongodb"
	"github.com/stretchr/testify/assert"
)

func Test_New(t *testing.T) {
	cfg := &config.Config{}
	act := mongodb.New(cfg)
	exp := &mongodb.MongoDB{}

	assert.NotEqual(t, exp, act)
}

func Test_Connect(t *testing.T) {
	mongo := []config.Mongo{{
		URI:            "mongodb://localhost:27017/",
		ConnectionName: "default",
		DatabaseName:   "echo-service-test",
	}}
	cfg := &config.Config{Mongos: mongo}
	mdb := mongodb.New(cfg)
	act := mdb.Connect()

	assert.NoError(t, act)

	// error connect to mongo beaculse invalid `uri`
	mongo = []config.Mongo{{
		URI:            "mongdb://localhost:27010/",
		ConnectionName: "default",
		DatabaseName:   "echo-service-test",
	}}
	cfg = &config.Config{Mongos: mongo}
	mdb = mongodb.New(cfg)
	act = mdb.Connect()

	assert.Error(t, act)

	// error connect to mongo can not ping mongo
	mongo = []config.Mongo{{
		URI:            "mongodb://localhost:27010/",
		ConnectionName: "default",
		DatabaseName:   "echo-service-test",
	}}
	cfg = &config.Config{Mongos: mongo}
	mdb = mongodb.New(cfg)
	act = mdb.Connect()

	assert.Error(t, act)
}

func Test_GetDB(t *testing.T) {
	mongo := []config.Mongo{{
		URI:            "mongodb://localhost:27017/",
		ConnectionName: "default",
		DatabaseName:   "echo-service-test",
	}}
	cfg := &config.Config{Mongos: mongo}
	mdb := mongodb.New(cfg)
	act := mdb.GetDB()

	assert.NotNil(t, act)
	mdb.Close()
}

func Test_Close(t *testing.T) {
	mongo := []config.Mongo{{
		URI:            "mongodb://localhost:27017/",
		ConnectionName: "default",
		DatabaseName:   "echo-service-test",
	}}
	cfg := &config.Config{Mongos: mongo}
	mdb := mongodb.New(cfg)

	mdb.Close()
}
