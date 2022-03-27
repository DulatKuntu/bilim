package controller

import (
	"context"
	"errors"
	"log"
	"os"

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

func (r *DatabaseRepository) CheckPassword(Username, Password string) (*model.User, error) {
	usersCollection := r.db.Collection(utils.CollectionUser)

	var userData model.User

	err := usersCollection.FindOne(
		context.TODO(),
		bson.M{"username": Username, "password": Password},
	).Decode(&userData)
	log.Print(err)
	if err != nil {
		return nil, errors.New("username of password are not correct")
	}

	return &userData, err
}

func (r *DatabaseRepository) CreateUser(signData *model.RequestUser) (*model.User, error) {
	usersCollection := r.db.Collection(utils.CollectionUser)

	id, err := GetMaxID([]bson.M{{"$group": bson.M{"_id": nil, "id": bson.M{"$max": "$id"}}}}, true, 1, usersCollection)

	if err != nil {
		return nil, err
	}
	var newUser model.User
	newUser.ID = id
	newUser.Email = signData.Email
	newUser.Username = signData.Username
	newUser.Name = signData.Name
	newUser.Surname = signData.Surname
	newUser.Bio = signData.Bio
	newUser.Password = signData.Password
	_, err = usersCollection.InsertOne(
		context.TODO(),
		newUser,
	)
	log.Print(err)

	return r.GetUserByEmail(signData.Email)
}
func (r *DatabaseRepository) InsertToken(UserID int, token string) error {
	usersCollection := r.db.Collection(utils.CollectionUser)

	_, err := usersCollection.UpdateOne(
		context.TODO(),
		bson.M{"id": UserID},
		bson.M{
			"$set": bson.M{
				"token": token,
			},
		},
	)

	return err
}

func (r *DatabaseRepository) GetIDByToken(token interface{}) (int, error) {
	usersCollection := r.db.Collection(utils.CollectionUser)

	var User model.User

	err := usersCollection.FindOne(
		context.TODO(),
		bson.M{"token": token},
	).Decode(&User)
	log.Print(token)
	if err != nil {
		return 0, err
	}
	log.Print(User)
	if token == "" {
		return 0, errors.New("Token not set")
	}
	return User.ID, nil

}
func (r *DatabaseRepository) GetUserByID(userID int) (*model.User, error) {
	usersCollection := r.db.Collection(utils.CollectionUser)

	var User model.User

	err := usersCollection.FindOne(
		context.TODO(),
		bson.M{"id": userID},
	).Decode(&User)

	if err != nil {
		return nil, err
	}

	return &User, nil
}

func (r *DatabaseRepository) UpdateProfile(userID int, user *model.User) error {

	usersCollection := r.db.Collection(utils.CollectionUser)

	var toChange model.User
	err := usersCollection.FindOne(
		context.TODO(),
		bson.M{
			"id": userID,
		},
	).Decode(&toChange)

	if err != nil {
		return err
	}
	if user.Bio != "" {
		toChange.Bio = user.Bio
	}
	if user.Name != "" {
		toChange.Name = user.Name
	}
	if user.Surname != "" {
		toChange.Surname = user.Surname
	}
	if user.Image != "" {
		toChange.Image = user.Image
	}

	usersCollection.UpdateOne(
		context.TODO(),
		bson.M{"userID": userID},
		bson.M{"bio": toChange.Bio, "name": toChange.Name, "surname": toChange.Surname, "image": toChange.Image},
	)
	locationImage, exists := os.LookupEnv("LocationProfileImage")

	if !exists {
		return errors.New("enviroment variable is not set")
	}

	err = os.Remove(locationImage + toChange.Image)
	if err != nil {
		return err
	}

	return nil
}

func (r *DatabaseRepository) AddInterest(interestID, userID int) error {
	usersCollection := r.db.Collection(utils.CollectionUser)

	_, err := usersCollection.UpdateOne(
		context.TODO(),
		bson.M{
			"id": userID,
		},
		bson.M{
			"$addToSet": bson.M{"interests": interestID},
		},
	)

	return err
}
