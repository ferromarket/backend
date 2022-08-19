package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ferromarket/backend/database"
	"github.com/ferromarket/backend/models"
	"github.com/ferromarket/backend/utils"
	"github.com/julienschmidt/httprouter"
	"gorm.io/gorm"
)

func PostUsuario(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()

	decoder := json.NewDecoder(request.Body)

	var usuario models.Usuario

	err := decoder.Decode(&usuario)
	if err != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(utils.ErrorMessage{ErrorMessage: err.Error()})
	}

	err = postUsuario(usuario, gdb)
	if err != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(utils.ErrorMessage{ErrorMessage: err.Error()})
	} else {
		writer.WriteHeader(http.StatusOK)
	}

	database.Close(gdb)
}

func postUsuario(usuario models.Usuario, gdb *gorm.DB) error {
	return gdb.Create(&usuario).Error
}

func ListUsuarios(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()

	type Usuarios struct {
		Usuarios []models.Usuario `json:"Usuario"`
	}

	usuarioList := Usuarios{}
	var usuarios []models.Usuario

	gdb.Model(&models.Usuario{}).Order("ID asc").Preload("Rol").Joins("LEFT JOIN usuario_rol ur ON usuario.id = ur.usuario_id").Find(&usuarios)

	usuarioList.Usuarios = usuarios

	for i := range usuarioList.Usuarios {
		usuarioList.Usuarios[i].Contrasena = nil
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(usuarioList)

	database.Close(gdb)
}

func GetUsuario(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()

	var usuario models.Usuario

	gdb.Model(&models.Usuario{}).Order("ID asc").Preload("Rol").Find(&usuario, params.ByName("id"))

	usuario.Contrasena = nil

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(usuario)

	database.Close(gdb)
}

// insertar o update. Necesita el objeto completo. Todos los atributos
func PutUsuario(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()

	decoder := json.NewDecoder(request.Body)

	var usuario models.Usuario

	err := decoder.Decode(&usuario)
	if err != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(utils.ErrorMessage{ErrorMessage: err.Error()})
	}

	usuario.ID, _ = strconv.ParseUint(params.ByName("id"), 10, 64)

	err = putUsuario(usuario, gdb)
	if err != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(utils.ErrorMessage{ErrorMessage: err.Error()})
	} else {
		writer.WriteHeader(http.StatusOK)
	}

	database.Close(gdb)
}

func putUsuario(usuario models.Usuario, gdb *gorm.DB) error{
	return gdb.Save(&usuario).Error
}

// solo update solo para un atributo
func PatchUsuario(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// hashClave()
	gdb := database.Connect()

	decoder := json.NewDecoder(request.Body)

	var usuario models.Usuario

	err := decoder.Decode(&usuario)
	if err != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(utils.ErrorMessage{ErrorMessage: err.Error()})
	}

	usuario.ID, _ = strconv.ParseUint(params.ByName("id"), 10, 64)

	var usuario_verif models.Usuario

	result := gdb.Find(&usuario_verif, usuario.ID)

	if result.RowsAffected == 0 {
		writer.WriteHeader(http.StatusNotFound)
	}else{
		err = patchUsuario(usuario, gdb)
		if err != nil {
			writer.Header().Set("Content-Type", "application/json")
			writer.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(writer).Encode(utils.ErrorMessage{ErrorMessage: err.Error()})
		} else {
			writer.WriteHeader(http.StatusOK)
		}
	}

	database.Close(gdb)
}

func patchUsuario(usuario models.Usuario, gdb *gorm.DB) error{
	return gdb.Updates(&usuario).Error
}

func DeleteUsuario(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()

	var usuario models.Usuario

	usuario.ID, _ = strconv.ParseUint(params.ByName("id"), 10, 64);
	result := deleteUsuario(&usuario, false, gdb)

	if result.Error != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(utils.ErrorMessage{ErrorMessage: result.Error.Error()})
	} else if result.RowsAffected == 0 {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)
		json.NewEncoder(writer).Encode(utils.ErrorMessage{ErrorMessage: "No existe el Usuario con ID: " + params.ByName("id") + "!"})
	} else {
		writer.WriteHeader(http.StatusOK)
	}

	database.Close(gdb)
}

func deleteUsuario(usuario *models.Usuario, hard bool, gdb *gorm.DB) *gorm.DB {
	if (hard) {
		// Delete the record
		return gdb.Unscoped().Delete(&usuario)
	} else {
		// Update the "deleted_at" column
		return gdb.Delete(&usuario)
	}
}