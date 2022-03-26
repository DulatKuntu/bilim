package handler

import (
	"log"
	"net/http"

	"github.com/DulatKuntu/bilim/requestHandler"
)

//"github.com/DulatKuntu/bilim/utils"

func (h *AppHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	var a interface{}
	userToken := requestHandler.GetToken(r, a)
	userID, err := h.Repo.GetIDByToken(userToken)
	if err != nil {
		DefaultErrorHandler(err, w)
		return
	}
	user, err := h.Repo.GetUserByID(userID)

	if err != nil {
		DefaultErrorHandler(err, w)
		return
	}

	SendGeneral(user, w)
}

func (h *AppHandler) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	user, err := requestHandler.GetUser(r)
	var a interface{}
	userToken := requestHandler.GetToken(r, a)
	userID, err := h.Repo.GetIDByToken(userToken)

	if err != nil {
		DefaultErrorHandler(err, w)
		return
	}

	err = h.Repo.UpdateProfile(userID, user)

	if err != nil {
		DefaultErrorHandler(err, w)
		return
	}
	SendGeneral(user, w)
}

func (h *AppHandler) AddInterests(w http.ResponseWriter, r *http.Request) {
	type str struct {
		Token      string `json:"token" bson:"token"`
		InterestID int    `json:"id" bson:"id"`
		Name       string `json:"name" bson:"name"`
	}
	var s str
	userToken := requestHandler.GetToken(r, s).(str)
	userID, err := h.Repo.GetIDByToken(userToken.Token)
	log.Print(err, userToken.Token)
	if err != nil {
		DefaultErrorHandler(err, w)
		return
	}
	interests, err := requestHandler.GetInterest(r)
	if err != nil {
		DefaultErrorHandler(err, w)
		return
	}
	err = h.Repo.AddInterest(interests.InterestID, userID)

	if err != nil {
		DefaultErrorHandler(err, w)
		return
	}

	SendGeneral(interests, w)
}

func (h *AppHandler) GetPosts(w http.ResponseWriter, r *http.Request) {
	var a interface{}
	userToken := requestHandler.GetToken(r, a)
	userID, err := h.Repo.GetIDByToken(userToken)
	if err != nil {
		DefaultErrorHandler(err, w)
		return
	}
	mentors, err := h.Repo.GetPosts(userID)

	if err != nil {
		DefaultErrorHandler(err, w)
		return
	}

	SendGeneral(mentors, w)
}
