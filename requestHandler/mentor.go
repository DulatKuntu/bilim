package requestHandler

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/DulatKuntu/bilim/model"
)

func GetInterestMentor(r *http.Request) (*model.Interests, error) {
	var s model.Interests
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &s)

	if err != nil {
		return nil, errors.New("bad request")
	}
	return &s, nil
}
