package controller

import (
	"context"
	"errors"
	"log"

	"github.com/DulatKuntu/bilim/model"
	"github.com/DulatKuntu/bilim/utils"
	"gopkg.in/mgo.v2/bson"
)

func (r *DatabaseRepository) GetMentorByUsername(username string) (*model.Mentor, error) {
	mentorsCollection := r.db.Collection(utils.CollectionMentor)

	var userData model.Mentor

	err := mentorsCollection.FindOne(
		context.TODO(),
		bson.M{"username": username},
	).Decode(&userData)

	return &userData, err
}

func (r *DatabaseRepository) GetMentorByEmail(Email string) (*model.Mentor, error) {
	mentorsCollection := r.db.Collection(utils.CollectionMentor)

	var mentorData model.Mentor

	err := mentorsCollection.FindOne(
		context.TODO(),
		bson.M{"email": Email},
	).Decode(&mentorData)

	return &mentorData, err
}
func (r *DatabaseRepository) CheckPasswordMentor(Username, Password string) (*model.Mentor, error) {
	mentorsCollection := r.db.Collection(utils.CollectionMentor)

	var mentorData model.Mentor

	err := mentorsCollection.FindOne(
		context.TODO(),
		bson.M{"username": Username, "password": Password},
	).Decode(&mentorData)
	log.Println(err)
	if err != nil {
		return nil, errors.New("username or password are not correct")
	}

	return &mentorData, err
}

func (r *DatabaseRepository) CreateMentor(signData *model.RequestMentor) (*model.Mentor, error) {
	mentorsCollection := r.db.Collection(utils.CollectionMentor)

	id, err := GetMaxID([]bson.M{{"$group": bson.M{"_id": nil, "id": bson.M{"$max": "$id"}}}}, true, 1, mentorsCollection)

	if err != nil {
		return nil, err
	}

	var newMentor model.Mentor
	newMentor.ID = id
	newMentor.Email = signData.Email
	newMentor.Username = signData.Username
	newMentor.Password = signData.Password
	newMentor.Name = signData.Name
	newMentor.Surname = signData.Surname
	newMentor.Bio = signData.Bio
	newMentor.Interests = signData.Interests
	newMentor.Students = signData.Students
	newMentor.Pending = signData.Pending
	_, err = mentorsCollection.InsertOne(
		context.TODO(),
		newMentor,
	)

	return r.GetMentorByEmail(signData.Email)
}

func (r *DatabaseRepository) GetMentorByID(mentorID int) (*model.Mentor, error) {
	mentorsCollection := r.db.Collection(utils.CollectionMentor)

	var Mentor model.Mentor

	err := mentorsCollection.FindOne(
		context.TODO(),
		bson.M{"id": mentorID},
	).Decode(&Mentor)

	if err != nil {
		return nil, err
	}

	return &Mentor, nil
}

func (r *DatabaseRepository) GetMentors(userID int) ([]*model.Mentor, error) {
	usersCollection := r.db.Collection(utils.CollectionUser)

	pipeline := []bson.M{
		bson.M{"$match": bson.M{"id": userID}},
		bson.M{"$lookup": bson.M{"from": "mentors", "localField": "interests", "foreignField": "interests", "as": "result"}},
		bson.M{"$project": bson.M{"result": 1, "_id": 0}},
	}

	cursor, err := usersCollection.Aggregate(
		context.TODO(),
		pipeline,
	)

	type resultParse struct {
		Result []*model.Mentor `json:"result" bson:"result"`
	}
	var mentors []*model.Mentor

	for cursor.Next(context.TODO()) {
		var result resultParse
		err = cursor.Decode(&result)
		if err != nil {
			return nil, errors.New("no comment exists")
		}
		for _, item := range result.Result {
			mentors = append(mentors, item)
		}
	}
	if err != nil {
		return nil, errors.New("username or password are not correct")
	}

	return mentors, err
}

func (r *DatabaseRepository) GetPosts(userID int) ([]*model.Mentor, error) {
	usersCollection := r.db.Collection(utils.CollectionUser)

	pipeline := []bson.M{
		bson.M{"$match": bson.M{"id": userID}},
		bson.M{"$lookup": bson.M{"from": "mentors", "localField": "interests", "foreignField": "interests", "as": "result"}},
		bson.M{"$project": bson.M{"result": 1, "_id": 0}},
	}

	cursor, err := usersCollection.Aggregate(
		context.TODO(),
		pipeline,
	)

	type resultParse struct {
		Result []*model.Mentor `json:"result" bson:"result"`
	}
	var mentors []*model.Mentor

	for cursor.Next(context.TODO()) {
		var result resultParse
		err = cursor.Decode(&result)
		if err != nil {
			return nil, errors.New("no comment exists")
		}
		for _, item := range result.Result {
			mentors = append(mentors, item)
		}
	}
	if err != nil {
		return nil, errors.New("username or password are not correct")
	}

	return mentors, err
}

func (r *DatabaseRepository) GetMentorIDByToken(token string) (int, error) {
	mentorsCollection := r.db.Collection(utils.CollectionMentor)

	var Mentor model.Mentor

	err := mentorsCollection.FindOne(
		context.TODO(),
		bson.M{"token": token},
	).Decode(&Mentor)

	log.Print(err)
	if err != nil {
		return 0, err
	}
	if token == "" {
		return 0, errors.New("Token not set")
	}
	return Mentor.ID, nil

}

func (r *DatabaseRepository) AddInterestMentor(interestID, userID int) error {
	mentorsCollection := r.db.Collection(utils.CollectionMentor)

	_, err := mentorsCollection.UpdateOne(
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

func (r *DatabaseRepository) AddMentor(userID, mentorID int) error {
	/*mentorsCollection := r.db.Collection(utils.CollectionMentor)

	_, err := mentorsCollection.InsertOne(
		context.TODO(),
		bson.M{
			"id": userID,
		},
		bson.M{
			"$addToSet": bson.M{"interests": interestID},
		},
	)

	return err
	*/
	return nil
}
