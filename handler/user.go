package handler

import (
	"net/http"

	"github.com/DulatKuntu/bilim/requestHandler"
)

//"github.com/DulatKuntu/bilim/utils"

func (h *AppHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	userToken := requestHandler.GetToken(r)
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
	userToken := requestHandler.GetToken(r)
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
	userToken := requestHandler.GetToken(r)
	userID, err := h.Repo.GetIDByToken(userToken)
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
