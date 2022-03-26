package model

import "github.com/DulatKuntu/bilim/utils"

type SystemData struct {
	UserID int
}

// DefaultPage
type DefaultPage struct {
	TotalPages  int         `json:"totalPages" bson:"totalPages"`
	CurrentPage int         `json:"currentPage" bson:"currentPage"`
	PageSize    int         `json:"pageSize" bson:"pageSize"`
	Data        interface{} `json:"data" bson:"data"`
}

func CreateDefaultPage(CurrentPage, PageSize, TotalDocuments int, Data interface{}) *DefaultPage {
	totalPages := TotalDocuments / PageSize

	if TotalDocuments%PageSize != 0 {
		totalPages++
	}

	return &DefaultPage{
		TotalPages:  totalPages,
		CurrentPage: CurrentPage,
		PageSize:    PageSize,
		Data:        Data,
	}
}

// DefaultResponse used in every response
type DefaultResponse struct {
	StatusCode int         `json:"statusCode"`
	Data       interface{} `json:"data"`
}

// ReqID struct
type ReqID struct {
	ID int `json:"id" bson:"id"`
}

// ReqIDString struct
type ReqIDString struct {
	ID string `json:"id" bson:"id"`
}

// ReqName struct
type ReqName struct {
	Name string `json:"name" bson:"name"`
}

// ReqPage struct
type ReqPage struct {
	Page int `json:"page" bson:"page"`
}

// Languages struct
type Languages struct {
	En string `json:"en" bson:"en"`
	Es string `json:"es" bson:"es"`
	De string `json:"de" bson:"de"`
	Fr string `json:"fr" bson:"fr"`
	It string `json:"it" bson:"it"`
	Pt string `json:"pt" bson:"pt"`
}

// SuccessResponse returns success
func SuccessResponse() *DefaultResponse {
	dr := &DefaultResponse{StatusCode: utils.StatusSuccess}
	return dr
}
