package handler

import (
	"net/http"

	"github.com/DulatKuntu/bilim/requestHandler"
)

func (h *AppHandler) PostBuddy(w http.ResponseWriter, r *http.Request) {
	PostData, err := requestHandler.GetPostData(r)
	if err != nil {
		DefaultErrorHandler(err, w)
		return
	}

	post := h.Repo.CreateBuddy(PostData)

	SendGeneral(post, w)

}
