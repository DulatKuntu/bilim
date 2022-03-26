package requestHandler

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/mail"

	"github.com/DulatKuntu/bilim/model"
)

// GetSignUp used to get signup struct from request, also for validation
func GetSignUp(r *http.Request) (*model.User, error) {
	var s model.User
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &s)

	if err != nil {
		return nil, errors.New("bad request")
	}
	_, err = mail.ParseAddress(s.Email) //checking for valid email

	if err != nil {
		return nil, errors.New("bad request")
	}

	return &s, nil
}
