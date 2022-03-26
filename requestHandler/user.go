package requestHandler

import (
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
