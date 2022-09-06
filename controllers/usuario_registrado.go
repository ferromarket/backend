package controllers

import (
	"encoding/base64"
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
	defer database.Close(gdb)

	decoder := json.NewDecoder(request.Body)

	var usuario models.Usuario

	err := decoder.Decode(&usuario)
	if err != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(utils.ErrorMessage{ErrorMessage: err.Error()})
		return
	}

	rawDecodedText, err := base64.StdEncoding.DecodeString(usuario.Contrasena)
    if err != nil {
        utils.JSONErrorOutput(writer, http.StatusBadRequest, err.Error())
    }
	usuario.Contrasena = string(rawDecodedText)

	usuario.HashPassword(usuario.Contrasena)

	err = postUsuario(usuario, gdb)
	if err != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(utils.ErrorMessage{ErrorMessage: err.Error()})
	} else {
		writer.WriteHeader(http.StatusOK)
	}
}

func postUsuario(usuario models.Usuario, gdb *gorm.DB) error {
	return gdb.Create(&usuario).Error
}

func ListUsuarios(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()
	defer database.Close(gdb)

	type Usuarios struct {
		Usuarios []models.Usuario `json:"Usuario"`
	}

	var usuarios []models.Usuario

	gdb.Model(&models.Usuario{}).Order("ID asc").Find(&usuarios)

	for i := range usuarios {
		usuarios[i].Contrasena = ""
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(usuarios)
}

func GetUsuario(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()
	defer database.Close(gdb)

	var usuario models.Usuario

	gdb.Model(&models.Usuario{}).Order("ID asc").Find(&usuario, params.ByName("id"))

	usuario.Contrasena = ""

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(usuario)
}

// insertar o update. Necesita el objeto completo. Todos los atributos
func PutUsuario(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()
	defer database.Close(gdb)

	decoder := json.NewDecoder(request.Body)

	var usuario models.Usuario

	err := decoder.Decode(&usuario)
	if err != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(utils.ErrorMessage{ErrorMessage: err.Error()})
		return
	}

	if usuario.Contrasena != "" {
		rawDecodedText, err := base64.StdEncoding.DecodeString(usuario.Contrasena)
		if err != nil {
			utils.JSONErrorOutput(writer, http.StatusBadRequest, err.Error())
		}
		usuario.Contrasena = string(rawDecodedText)
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
}

func putUsuario(usuario models.Usuario, gdb *gorm.DB) error {
	return gdb.Save(&usuario).Error
}

// solo update solo para un atributo
func PatchUsuario(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// hashClave()
	gdb := database.Connect()
	defer database.Close(gdb)

	decoder := json.NewDecoder(request.Body)

	var usuario models.Usuario

	err := decoder.Decode(&usuario)
	if err != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(utils.ErrorMessage{ErrorMessage: err.Error()})
		return
	}

	if usuario.Contrasena != "" {
		rawDecodedText, err := base64.StdEncoding.DecodeString(usuario.Contrasena)
		if err != nil {
			utils.JSONErrorOutput(writer, http.StatusBadRequest, err.Error())
		}
		usuario.Contrasena = string(rawDecodedText)
	}

	usuario.ID, _ = strconv.ParseUint(params.ByName("id"), 10, 64)

	var usuario_verif models.Usuario

	result := gdb.Find(&usuario_verif, usuario.ID)

	if result.RowsAffected == 0 {
		writer.WriteHeader(http.StatusNotFound)
	} else {
		err = patchUsuario(usuario, gdb)
		if err != nil {
			writer.Header().Set("Content-Type", "application/json")
			writer.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(writer).Encode(utils.ErrorMessage{ErrorMessage: err.Error()})
			return
		} else {
			writer.WriteHeader(http.StatusOK)
		}
	}
}

func patchUsuario(usuario models.Usuario, gdb *gorm.DB) error {
	return gdb.Updates(&usuario).Error
}

func DeleteUsuario(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()
	defer database.Close(gdb)

	var usuario models.Usuario

	usuario.ID, _ = strconv.ParseUint(params.ByName("id"), 10, 64)
	result := deleteUsuario(&usuario, false, gdb)

	if result.Error != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(utils.ErrorMessage{ErrorMessage: result.Error.Error()})
		return
	} else if result.RowsAffected == 0 {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)
		json.NewEncoder(writer).Encode(utils.ErrorMessage{ErrorMessage: "No existe el Usuario con ID: " + params.ByName("id") + "!"})
	} else {
		writer.WriteHeader(http.StatusOK)
	}
}

func deleteUsuario(usuario *models.Usuario, hard bool, gdb *gorm.DB) *gorm.DB {
	if hard {
		// Delete the record
		return gdb.Unscoped().Delete(&usuario)
	} else {
		// Update the "deleted_at" column
		return gdb.Delete(&usuario)
	}
}
