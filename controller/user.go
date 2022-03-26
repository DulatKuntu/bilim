package controller

import (
	"context"
	"log"

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

func (r *DatabaseRepository) GetMentorByUsername(username string) (*model.Mentor, error) {
	mentorsCollection := r.db.Collection(utils.CollectionMentor)

	var userData model.Mentor

	err := mentorsCollection.FindOne(
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

func (r *DatabaseRepository) GetMentorByEmail(Email string) (*model.User, error) {
	mentorsCollection := r.db.Collection(utils.CollectionMentor)

	var mentorData model.User

	err := mentorsCollection.FindOne(
		context.TODO(),
		bson.M{"email": Email},
	).Decode(&mentorData)

	return &mentorData, err
}

func (r *DatabaseRepository) CreateUser(signData *model.RequestUser) (*model.User, error) {
	usersCollection := r.db.Collection(utils.CollectionUser)

	id, err := GetMaxID([]bson.M{{"$group": bson.M{"_id": nil, "id": bson.M{"$max": "$id"}}}}, true, 1, usersCollection)

	if err != nil {
		return nil, err
	}

	//var newUser model.User
	// newUser.Email = signData.Email
	// newUser.Username = signData.Username
	// newUser.Password = signData.Password
	// newUser.Name = signData.Name
	// newUser.Surname = signData.Surname
	// newUser.Bio = signData.Bio
	// newUser.Interests = signData.Interests
	var newUser model.User
	newUser.ID = id
	newUser.Email = signData.Email
	newUser.Username = signData.Username
	newUser.Name = signData.Name
	newUser.Surname = signData.Surname
	newUser.Bio = signData.Bio

	_, err = usersCollection.InsertOne(
		context.TODO(),
		newUser,
	)
	log.Print(err)

	return r.GetUserByEmail(signData.Email)
}

func (r *DatabaseRepository) CreateMentor(signData *model.RequestMentor) (*model.User, error) {
	mentorsCollection := r.db.Collection(utils.CollectionMentor)

	id, err := GetMaxID([]bson.M{{"$group": bson.M{"_id": nil, "id": bson.M{"$max": "$id"}}}}, true, 1, mentorsCollection)

	if err != nil {
		return nil, err
	}

	var newMentor model.User
	newMentor.ID = id
	newMentor.Email = signData.Email
	newMentor.Username = signData.Username
	newMentor.Name = signData.Name
	newMentor.Surname = signData.Surname
	newMentor.Bio = signData.Bio
	_, err = mentorsCollection.InsertOne(
		context.TODO(),
		newMentor,
	)

	return r.GetMentorByEmail(signData.Email)
}

func (r *DatabaseRepository) Sign() {
	usersCollection := r.db.Collection(utils.CollectionUser)

	usersCollection.InsertOne(
		context.TODO(),
		bson.M{"userId": 2},
	)
}
