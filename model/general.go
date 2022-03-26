package model

import (
	"github.com/DulatKuntu/bilim/utils"
)

// DefaultResponse used in every response
type DefaultResponse struct {
	StatusCode int         `json:"statusCode"`
	Data       interface{} `json:"data"`
}

// ReqID struct
type ReqID struct {
	ID int `json:"id" bson:"id"`
}

// SuccessResponse returns success
func SuccessResponse() *DefaultResponse {
	dr := &DefaultResponse{StatusCode: utils.StatusSuccess}
	return dr
}
