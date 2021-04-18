package mongodb

import (
	"context"

	"github.com/shandysiswandi/echo-service/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoDB struct {
	config  *config.Config
	context context.Context
	client  *mongo.Client
	db      *mongo.Database
}

func New(cfg *config.Config, ctx context.Context) *MongoDB {
	return &MongoDB{
		config:  cfg,
		context: ctx,
		client:  nil,
		db:      nil,
	}
}

func (mDB *MongoDB) Connect() error {
	client, err := mongo.Connect(mDB.context, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return err
	}

	if err := client.Ping(mDB.context, readpref.Primary()); err != nil {
		return err
	}

	mDB.client = client
	mDB.db = client.Database("echo-service")

	return nil
}

func (mDB *MongoDB) Close() {
	if mDB.client == nil {
		return
	}

	if err := mDB.client.Disconnect(mDB.context); err != nil {
		panic(err)
	}
}

func (mDB *MongoDB) GetDB() *mongo.Database {
	if mDB.client == nil || mDB.db == nil {
		mDB.Connect()
	}

	return mDB.db
}
