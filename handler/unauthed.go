package handler

import (
	"errors"
	"log"
	"net/http"

	"github.com/DulatKuntu/bilim/requestHandler"
	"github.com/DulatKuntu/bilim/utils"
	//"github.com/DulatKuntu/bilim/utils"
)

// SingUp used to Sing UP
func (h *AppHandler) SignUp(w http.ResponseWriter, r *http.Request) {
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
	signupData, err := requestHandler.GetSignUpMentor(r)
	if err != nil {
		DefaultErrorHandler(err, w)
		return
	}

	_, err = h.Repo.GetUserByUsername(signupData.Username)
	if err == nil {
		DefaultErrorHandler(errors.New("username is already taken"), w)
		return
	}
	_, err = h.Repo.GetMentorByUsername(signupData.Username)
	if err == nil {
		DefaultErrorHandler(errors.New("username is already taken"), w)
		return
	}
	user, err := h.Repo.GetMentorByEmail(signupData.Email)

	if err == nil { //case user already signedIn with this email before
		SendGeneral(user, w)
		return
	}
	user, err = h.Repo.CreateMentor(signupData)
	if err != nil {
		DefaultErrorHandler(err, w)
		return
	}
	SendGeneral(user, w)
}

func (h *AppHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	loginData, err := requestHandler.GetLogin(r)
	if err != nil {
		DefaultErrorHandler(err, w)
		return
	}

	token := utils.GenerateToken(loginData.Username)
	log.Println(loginData.Username, loginData.Password)
	user, err := h.Repo.CheckPassword(loginData.Username, loginData.Password)

	if err != nil {
		DefaultErrorHandler(err, w)
		return
	}

	err = h.Repo.InsertToken(user.ID, token)

	if err != nil {
		DefaultErrorHandler(err, w)
		return
	}
	SendGeneral(user, w)
}

func (h *AppHandler) SignInMentor(w http.ResponseWriter, r *http.Request) {
	loginData, err := requestHandler.GetLoginMentor(r)
	if err != nil {
		DefaultErrorHandler(err, w)
		return
	}

	token := utils.GenerateToken(loginData.Username)
	log.Println(loginData.Username, loginData.Password)
	mentor, err := h.Repo.CheckPasswordMentor(loginData.Username, loginData.Password)

	mentor.Token = token

	if err != nil {
		DefaultErrorHandler(err, w)
		return
	}

	err = h.Repo.InsertToken(mentor.ID, token)

	if err != nil {
		DefaultErrorHandler(err, w)
		return
	}
	SendGeneral(mentor, w)
}
