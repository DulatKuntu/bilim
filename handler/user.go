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
