package controller

import (
	"context"

	"github.com/DulatKuntu/bilim/model"
	"github.com/DulatKuntu/bilim/utils"
	"gopkg.in/mgo.v2/bson"
)

func (r *DatabaseRepository) CreateBuddy(PostData *model.RequestBuddyPost, userID int) *model.BuddyPost {
	postsCollection := r.db.Collection(utils.CollectionPosts)
	usersCollection := r.db.Collection(utils.CollectionUser)

	var user model.User
	err := usersCollection.FindOne(
		context.TODO(),
		bson.M{"id": userID},
	).Decode(&user)

	id, err := GetMaxID([]bson.M{{"$group": bson.M{"_id": nil, "id": bson.M{"$max": "$id"}}}}, true, 1, postsCollection)

	if err != nil {
		return nil
	}
	var newPost model.BuddyPost
	newPost.UserID = userID
	newPost.Name = user.Name
	newPost.Surname = user.Surname
	newPost.Description = PostData.Description
	newPost.ID = id
	_, err = postsCollection.InsertOne(
		context.TODO(),
		newPost,
	)

	return &newPost
}
