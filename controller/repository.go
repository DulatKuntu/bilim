package controller

import "go.mongodb.org/mongo-driver/mongo"

type DatabaseRepository struct {
	db     *mongo.Database
	Client *mongo.Client
}

func NewDatabaseRepository(db *mongo.Database, client *mongo.Client) *DatabaseRepository {
	return &DatabaseRepository{db: db, Client: client}
}
