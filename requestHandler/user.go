package requestHandler

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/DulatKuntu/bilim/model"
)

// GetSignUp used to get signup struct from request, also for validation
func GetToken(r *http.Request) string {
	reqToken := r.Header.Get("Authorization")
	return reqToken
}

func GetUser(r *http.Request) (*model.User, error) {
	var s model.User
	s.Name = r.FormValue("name")
	s.Surname = r.FormValue("surname")
	s.Bio = r.FormValue("bio")
	imageName, err := CreateProfileImage(r)
	if err != nil {
		return nil, err
	}
	s.Image = imageName
	return &s, nil
}

func GetInterest(r *http.Request) (*model.Interests, error) {
	var s model.Interests
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &s)

	if err != nil {
		return nil, errors.New("bad request")
	}
	return &s, nil
}
