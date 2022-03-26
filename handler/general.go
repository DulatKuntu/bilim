package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/DulatKuntu/bilim/model"
	"go.mongodb.org/mongo-driver/mongo"
)

// ServerError creator
func ServerError(StatusCode int) *model.DefaultResponse {
	response := &model.DefaultResponse{}
	response.StatusCode = StatusCode
	return response
}

// ServerErrorMessage creator
func ServerErrorMessage(StatusCode int, Message string) *model.DefaultResponse {
	response := &model.DefaultResponse{}
	response.StatusCode = StatusCode
	response.Data = Message
	return response
}

// DefaultErrorHandler error only with status code
func DefaultErrorHandler(Err error, w http.ResponseWriter) {
	if Err == mongo.ErrNoDocuments {
		Err = errors.New("not found")
	}

	fullError := Err.Error()

	parts := strings.Split(fullError, "|")
	mainMessage := strings.TrimSpace(parts[0])

	switch mainMessage {

	case "bad request":
		json.NewEncoder(w).Encode(ServerErrorMessage(2001, fullError))
	case "not found":
		json.NewEncoder(w).Encode(ServerErrorMessage(2002, fullError))
	case "incorrect boyman":
		json.NewEncoder(w).Encode(ServerErrorMessage(2003, fullError))
	case "incorrect token":
		json.NewEncoder(w).Encode(ServerErrorMessage(2004, fullError))
	case "username is already taken":
		json.NewEncoder(w).Encode(ServerErrorMessage(2005, fullError))

	case "file system":
		json.NewEncoder(w).Encode(ServerErrorMessage(8001, fullError))
	default:
		json.NewEncoder(w).Encode(ServerErrorMessage(8000, fullError))
	}
}

// SendGeneral sends general data
func SendGeneral(data interface{}, w http.ResponseWriter) {
	gr := model.SuccessResponse()
	gr.Data = data

	json.NewEncoder(w).Encode(gr)
}

// SendStatus sends response with status and empty data
func SendStatus(statusCode int, w http.ResponseWriter) {
	gr := &model.DefaultResponse{StatusCode: statusCode}

	json.NewEncoder(w).Encode(gr)
}
