package handler

import (
	"errors"
	"net/http"

	"github.com/DulatKuntu/bilim/requestHandler"
	//"github.com/DulatKuntu/bilim/utils"
)

// SingUp used to Sing UP
func (h *AppHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/all")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	signupData, err := requestHandler.GetSignUp(r)
	if err != nil {
		DefaultErrorHandler(err, w)
		return
	}
	_, err = h.Repo.GetUserByUsername(signupData.Username)

	if err == nil {
		DefaultErrorHandler(errors.New("username is already taken"), w)
		return
	}
	user, err := h.Repo.GetUserByEmail(signupData.Email)
	if err == nil { //case user already signedIn with this email before
		SendGeneral(user, w)
		return
	}
	user, err = h.Repo.CreateUser(signupData)
	if err != nil {
		DefaultErrorHandler(err, w)
		return
	}
	SendGeneral(user, w)
}

func (h *AppHandler) SignUpMentor(w http.ResponseWriter, r *http.Request) {

	h.Repo.Sign()
}
