package controller

import (
	"context"

	"github.com/DulatKuntu/bilim/model"
	"github.com/DulatKuntu/bilim/utils"
	"gopkg.in/mgo.v2/bson"
)

func (r *DatabaseRepository) CreatePost(PostData *model.RequestBuddyPost) *model.BuddyPost {
	postsCollection := r.db.Collection(utils.CollectionPosts)

	id, err := GetMaxID([]bson.M{{"$group": bson.M{"_id": nil, "id": bson.M{"$max": "$id"}}}}, true, 1, postsCollection)

	if err != nil {
		return nil
	}
	var newPost *model.BuddyPost
	newPost.UserID = PostData.UserID
	newPost.Name = PostData.Name
	newPost.Surname = PostData.Surname
	newPost.Description = PostData.Description
	newPost.ID = id
	_, err = postsCollection.InsertOne(
		context.TODO(),
		newPost,
	)

	return newPost
}
