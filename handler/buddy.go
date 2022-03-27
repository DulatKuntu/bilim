package handler

import (
	"net/http"

	"github.com/DulatKuntu/bilim/requestHandler"
)

func (h *AppHandler) PostBuddy(w http.ResponseWriter, r *http.Request) {
	userToken := requestHandler.GetToken(r)
	userID, err := h.Repo.GetIDByToken(userToken)
	if err != nil {
		DefaultErrorHandler(err, w)
		return
	}
	PostData, err := requestHandler.GetPostData(r)
	if err != nil {
		DefaultErrorHandler(err, w)
		return
	}

	post := h.Repo.CreateBuddy(PostData, userID)

	SendGeneral(post, w)

}
