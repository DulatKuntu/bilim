package requestHandler

import (
	"net/http"
)

// GetSignUp used to get signup struct from request, also for validation
func GetToken(r *http.Request) string {
	reqToken := r.Header.Get("Authorization")
	return reqToken
}
