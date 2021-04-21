package mongodb

import (
	"context"
	"time"

	"github.com/shandysiswandi/echo-service/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoDB struct {
	config *config.Config
	client *mongo.Client
	db     *mongo.Database
}

func New(cfg *config.Config) *MongoDB {
	return &MongoDB{
		config: cfg,
		client: nil,
		db:     nil,
	}
}

func (mDB *MongoDB) Connect() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mDB.config.Mongos[0].URI))
	if err != nil {
		return err
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}

	mDB.client = client
	mDB.db = client.Database(mDB.config.Mongos[0].DatabaseName)

	return nil
}

func (mDB *MongoDB) Close() {
	if mDB.client == nil {
		return
	}

	if err := mDB.client.Disconnect(context.Background()); err != nil {
		println(err.Error())
	}
}

func (mDB *MongoDB) GetDB() *mongo.Database {
	if mDB.client == nil || mDB.db == nil {
		mDB.Connect()
	}

	return mDB.db
}
