package handler

import (
	"log"
	"net/http"

	"github.com/DulatKuntu/bilim/requestHandler"
)

//"github.com/DulatKuntu/bilim/utils"

func (h *AppHandler) GetProfileMentor(w http.ResponseWriter, r *http.Request) {
	mentorToken := requestHandler.GetToken(r)
	mentorID, err := h.Repo.GetIDByToken(mentorToken)
	if err != nil {
		DefaultErrorHandler(err, w)
		return
	}
	mentor, err := h.Repo.GetMentorByID(mentorID)

	if err != nil {
		DefaultErrorHandler(err, w)
		return
	}

	SendGeneral(mentor, w)
}

func (h *AppHandler) GetMentors(w http.ResponseWriter, r *http.Request) {
	userToken := requestHandler.GetToken(r)
	log.Println(userToken)
	userID, err := h.Repo.GetIDByToken(userToken)
	log.Println(userID, err)
	if err != nil {
		DefaultErrorHandler(err, w)
		return
	}
	mentors, err := h.Repo.GetMentors(userID)
	log.Println(mentors, err)
	if err != nil {
		DefaultErrorHandler(err, w)
		return
	}

	SendGeneral(mentors, w)
}

func (h *AppHandler) AddMentorInterests(w http.ResponseWriter, r *http.Request) {
	mentorToken := requestHandler.GetToken(r)
	mentorID, err := h.Repo.GetMentorIDByToken(mentorToken)
	if err != nil {
		DefaultErrorHandler(err, w)
		return
	}
	interests, err := requestHandler.GetInterestMentor(r)
	if err != nil {
		DefaultErrorHandler(err, w)
		return
	}
	err = h.Repo.AddInterestMentor(interests.InterestID, mentorID)
	log.Println(err)
	if err != nil {
		DefaultErrorHandler(err, w)
		return
	}

	SendGeneral(interests, w)
}

func (h *AppHandler) AddMentor(w http.ResponseWriter, r *http.Request) {
	userToken := requestHandler.GetToken(r)
	userID, err := h.Repo.GetMentorIDByToken(userToken)
	if err != nil {
		DefaultErrorHandler(err, w)
		return
	}
	mentorID, err := requestHandler.GetMentorID(r)
	if err != nil {
		DefaultErrorHandler(err, w)
		return
	}
	err = h.Repo.AddMentor(userID, mentorID)
	if err != nil {
		DefaultErrorHandler(err, w)
		return
	}

	SendGeneral(mentorID, w)
}
