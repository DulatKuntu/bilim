package controller

import (
	"github.com/DulatKuntu/bilim/model"
	"github.com/DulatKuntu/bilim/utils"
)

func (r *DatabaseRepository) CreatePost(PostData *model.BuddyPost) {
	postsCollection := r.db.Collection(utils.CollectionPosts)

	newPost := model.BuddyPost*

	_, err := postsCollection.InsertOne(
		context.TODO(),
		newPost,
	)
	log.Print(err)

	return r.GetPostById(signData.Email)
}
