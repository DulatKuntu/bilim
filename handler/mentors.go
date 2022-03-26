package handler

import (
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
	userID, err := h.Repo.GetIDByToken(userToken)
	if err != nil {
		DefaultErrorHandler(err, w)
		return
	}
	mentors, err := h.Repo.GetMentors(userID)

	if err != nil {
		DefaultErrorHandler(err, w)
		return
	}

	SendGeneral(mentors, w)
}
