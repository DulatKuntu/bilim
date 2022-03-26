package requestHandler

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/DulatKuntu/bilim/model"
)

func getPostData(r *http.Request) (*model.BuddyPost, error) {
	var s model.BuddyPost

	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &s)

	if err != nil {
		return nil, errors.New("bad request")
	}

	return &s, nil
}
