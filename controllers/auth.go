package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/ferromarket/backend/database"
	"github.com/ferromarket/backend/models"
	"github.com/ferromarket/backend/utils"
	"github.com/julienschmidt/httprouter"
)

func AuthenticateUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	type Message struct {
		Status string `json:"status"`
	}

	message := Message{Status: "authorized"}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(message)
}

func Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()
	defer database.Close(gdb)

	decoder := json.NewDecoder(request.Body)

	var receivedUser models.Usuario
	var user models.Usuario

	err := decoder.Decode(&receivedUser)
	if (err != nil) {
		utils.JSONErrorOutput(writer, http.StatusBadRequest, err.Error())
		return
	}

	gdb.Model(models.Usuario{}).Where(&models.Usuario{RUT: receivedUser.RUT}).Find(&user)

	err = user.CheckPassword(receivedUser.Contrasena)
	if err != nil {
		utils.JSONErrorOutput(writer, http.StatusBadRequest, errors.New("contraseña incorrecta").Error())
		return
	}

	type Token struct {
		Token string `json:"token"`
	}

	tokenString, err := utils.GenerateJWT(strconv.FormatUint(user.ID, 10), user.RUT, user.Email)
	if err != nil {
		utils.JSONErrorOutput(writer, http.StatusBadRequest, err.Error())
		return
	}

	token := Token{Token: tokenString}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(token)
}
