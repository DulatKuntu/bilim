package controller

import (
	"context"

	"github.com/DulatKuntu/bilim/model"
	"github.com/DulatKuntu/bilim/utils"
	"gopkg.in/mgo.v2/bson"
)

func (r *DatabaseRepository) GetUserByUsername(username string) (*model.User, error) {
	usersCollection := r.db.Collection(utils.CollectionUser)

	var userData model.User

	err := usersCollection.FindOne(
		context.TODO(),
		bson.M{"username": username},
	).Decode(&userData)

	return &userData, err
}

func (r *DatabaseRepository) GetUserByEmail(Email string) (*model.User, error) {
	usersCollection := r.db.Collection(utils.CollectionUser)

	var userData model.User

	err := usersCollection.FindOne(
		context.TODO(),
		bson.M{"email": Email},
	).Decode(&userData)

	return &userData, err
}

func (r *DatabaseRepository) CreateUser(signData *model.User) (*model.User, error) {
	usersCollection := r.db.Collection(utils.CollectionUser)

	id, err := GetMaxID([]bson.M{{"$group": bson.M{"_id": nil, "id": bson.M{"$max": "$id"}}}}, true, 1, usersCollection)

	if err != nil {
		return nil, err
	}

	var newUser model.User
	newUser.ID = id
	newUser.Email = signData.Email
	newUser.Username = signData.Username

	usersCollection.InsertOne(
		context.TODO(),
		newUser,
	)

	return r.GetUserByEmail(signData.Email)
}
