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

func PostFerreteria(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()

	decoder := json.NewDecoder(request.Body)

	var ferreteria models.Ferreteria

	err := decoder.Decode(&ferreteria)
	if (err != nil) {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(utils.ErrorMessage{ErrorMessage: err.Error()})
	}

	result := postFerreteria(ferreteria, gdb)
	if (result.Error != nil) {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(utils.ErrorMessage{ErrorMessage: result.Error.Error()})
	} else {
		writer.WriteHeader(http.StatusOK)
	}

	database.Close(gdb)
}

func postFerreteria(ferreteria models.Ferreteria, gdb *gorm.DB) *gorm.DB {
	return gdb.Create(&ferreteria)
}

func ListFerreterias(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()

	var ferreterias []models.Ferreteria

	result := listFerreterias(&ferreterias, gdb)
	if (result.Error != nil) {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(utils.ErrorMessage{ErrorMessage: result.Error.Error()})
	} else {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode(ferreterias)
	}

	database.Close(gdb)
}

func listFerreterias(ferreterias *[]models.Ferreteria, gdb *gorm.DB) *gorm.DB {
	return gdb.Model(&models.Ferreteria{}).Order("ID asc").Preload("Horarios.Dia").Preload("Horarios.Abrir").Preload("Horarios.Cerrar").Preload("Comuna.Ciudad.Region.Pais").Joins("LEFT JOIN ferreteria_horario fh ON ferreteria.id = fh.ferreteria_id").Find(&ferreterias)
}

func GetFerreteria(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()

	var ferreteria models.Ferreteria

	result := getFerreteria(&ferreteria, params.ByName("id"), gdb)
	if (result.Error != nil) {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(utils.ErrorMessage{ErrorMessage: result.Error.Error()})
	} else if (result.RowsAffected == 0) {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)
		json.NewEncoder(writer).Encode(utils.ErrorMessage{ErrorMessage: "No existe ferreteria con id " + params.ByName("id") + "!"})
	} else {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode(ferreteria)
	}

	database.Close(gdb)
}

func getFerreteria(ferreteria *models.Ferreteria, id string, gdb *gorm.DB) *gorm.DB {
	return gdb.Model(&models.Ferreteria{}).Order("ID asc").Preload("Horarios.Dia").Preload("Horarios.Abrir").Preload("Horarios.Cerrar").Preload("Comuna.Ciudad.Region.Pais").Joins("LEFT JOIN ferreteria_horario fh ON ferreteria.id = fh.ferreteria_id").Find(&ferreteria, id)
}

func PutFerreteria(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()

	var ferreteria models.Ferreteria

	decoder := json.NewDecoder(request.Body)

	err := decoder.Decode(&ferreteria)
	if (err != nil) {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(utils.ErrorMessage{ErrorMessage: err.Error()})
	}

	ferreteria.ID, _ = strconv.ParseUint(params.ByName("id"), 10, 64)

	result := putFerreteria(&ferreteria, gdb)
	if (result.Error != nil) {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(utils.ErrorMessage{ErrorMessage: result.Error.Error()})
	} else if (result.RowsAffected == 0) {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)
		json.NewEncoder(writer).Encode(utils.ErrorMessage{ErrorMessage: "No existe ferreteria con id " + params.ByName("id") + "!"})
	} else {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode(ferreteria)
	}

	database.Close(gdb)
}

func putFerreteria(ferreteria *models.Ferreteria, gdb *gorm.DB) *gorm.DB {
	return gdb.Save(&ferreteria)
}

func PatchFerreteria(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()

	var ferreteria models.Ferreteria

	decoder := json.NewDecoder(request.Body)

	err := decoder.Decode(&ferreteria)
	if (err != nil) {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(utils.ErrorMessage{ErrorMessage: err.Error()})
	}

	ferreteria.ID, _ = strconv.ParseUint(params.ByName("id"), 10, 64)

	result := patchFerreteria(&ferreteria, gdb)
	if (result.Error != nil) {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(utils.ErrorMessage{ErrorMessage: result.Error.Error()})
	} else if (result.RowsAffected == 0) {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)
		json.NewEncoder(writer).Encode(utils.ErrorMessage{ErrorMessage: "No existe ferreteria con id " + params.ByName("id") + "!"})
	} else {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode(ferreteria)
	}

	database.Close(gdb)
}

func patchFerreteria(ferreteria *models.Ferreteria, gdb *gorm.DB) *gorm.DB {
	return gdb.Updates(&ferreteria)
}

func DeleteFerreteria(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()

	var ferreteria models.Ferreteria
	ferreteria.ID, _ = strconv.ParseUint(params.ByName("id"), 10, 64)

	result := deleteFerreteria(&ferreteria, false, gdb)
	if (result.Error != nil) {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(utils.ErrorMessage{ErrorMessage: result.Error.Error()})
	} else if (result.RowsAffected == 0) {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)
		json.NewEncoder(writer).Encode(utils.ErrorMessage{ErrorMessage: "No existe ferreteria con id " + params.ByName("id") + "!"})
	} else {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode(ferreteria)
	}

	database.Close(gdb)
}

func deleteFerreteria(ferreteria *models.Ferreteria, hard bool, gdb *gorm.DB) *gorm.DB {
	if (hard) {
		// Delete the record
		return gdb.Unscoped().Delete(&ferreteria)
	} else {
		// Update the "deleted_at" column
		return gdb.Delete(&ferreteria)
	}
}
