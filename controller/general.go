package controller

import (
	"context"

	"github.com/DulatKuntu/bilim/model"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

func GetMaxID(pipe []bson.M, WithIncrement bool, FirstID int, collection *mongo.Collection) (int, error) {
	var resID []*model.ReqID
	var newId int

	cursor, err := collection.Aggregate(context.TODO(), pipe)

	if err != nil { //Case aggregation with unwind gives error
		return 1, nil
	}

	if err = cursor.All(context.TODO(), &resID); err != nil {
		return 1, nil //Case aggregation with unwind gives error
	}

	if len(resID) > 0 {
		newId = resID[0].ID
		if WithIncrement {
			newId++
		}
	} else {
		newId = FirstID
	}

	return newId, nil
}
