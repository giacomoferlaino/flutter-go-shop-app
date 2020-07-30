package api

import (
	"encoding/json"
	"flutter_shop_app/auth"
	"fmt"
	"net/http"
)

func sendError(res http.ResponseWriter, err error) {
	res.WriteHeader(http.StatusInternalServerError)
	fmt.Fprint(res, NewErrorResponse(err.Error()).ToJSON())
	return
}

func sendSuccess(res http.ResponseWriter, data interface{}, rowsAffected int64) {
	apiResponse, err := json.Marshal(NewSuccessResponse(data, rowsAffected))
	if err != nil {
		sendError(res, err)
		return
	}
	res.WriteHeader(http.StatusOK)
	fmt.Fprintln(res, string(apiResponse))
}

func authenticate(req *http.Request, jwtManager *auth.JwtManager) (uint, error) {
	authToken := req.Header.Get("Authorization")
	return jwtManager.ParseToken(authToken)
}
