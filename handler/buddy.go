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

	post, err := h.Repo.CreatePost(PostData)

	SendGeneral(post, w)

}
