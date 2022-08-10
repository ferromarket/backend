package middlewares

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/ferromarket/backend/utils"

	"github.com/julienschmidt/httprouter"
)

func Authenticate(handle httprouter.Handle) httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		type ErrorMessage struct {
			ErrorMessage string `json:"error_message"`
		}

		reqToken := request.Header.Get("Authorization")
		splitToken := strings.Split(reqToken, "Bearer ")
		if len(splitToken) < 2 {
			writer.Header().Set("Content-Type", "application/json")
			writer.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(writer).Encode(ErrorMessage{ErrorMessage: errors.New("no token received").Error()})
			return
		}
		tokenString := splitToken[1]

		err := utils.ValidateToken(tokenString)
		if err != nil {
			writer.Header().Set("Content-Type", "application/json")
			writer.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(writer).Encode(ErrorMessage{ErrorMessage: err.Error()})
			return
		}

		handle(writer, request, params)
	}
}
