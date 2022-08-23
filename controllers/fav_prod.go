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

func PostFavProd(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()

	decoder := json.NewDecoder(request.Body)

	var favProd models.FavProd

	err := decoder.Decode(&favProd)
	if err != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(utils.ErrorMessage{ErrorMessage: err.Error()})
	}

	err = postFavProd(favProd, gdb)
	if err != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(utils.ErrorMessage{ErrorMessage: err.Error()})
	} else {
		writer.WriteHeader(http.StatusOK)
	}

	database.Close(gdb)
}

func postFavProd(favProd models.FavProd, gdb *gorm.DB) error {
	return gdb.Create(&favProd).Error
}

func ListFavProd(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()

	type FavProd struct {
		FavProd []models.FavProd `json:"FavProd"`
	}

	favProdList := FavProd{}
	var favProd []models.FavProd

	gdb.Model(&models.FavProd{}).Order("ID asc").Preload("Usuario").Find(&favProd)
//gdb.Model(&models.FavProd{}).Order("ID asc").Preload("Usuario").Preload("Producto").Find(&favProd)
	favProdList.FavProd = favProd

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(favProdList)

	database.Close(gdb)
}

func GetFavProd(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()

	var favProd models.FavProd

	gdb.Model(&models.FavProd{}).Order("ID asc").Preload("Rol").Find(&favProd, params.ByName("id"))

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(favProd)

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
