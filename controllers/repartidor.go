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

func PostRepartidor(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()
	defer database.Close(gdb)

	decoder := json.NewDecoder(request.Body)

	var repartidor models.Repartidor

	err := decoder.Decode(&repartidor)
	if err != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(utils.ErrorMessage{ErrorMessage: err.Error()})
		return
	}

	result := postRepartidor(repartidor, gdb)
	if result.Error != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(utils.ErrorMessage{ErrorMessage: result.Error.Error()})
		return
	} else {
		writer.WriteHeader(http.StatusOK)
		return
	}
}

func postRepartidor(repartidor models.Repartidor, gdb *gorm.DB) *gorm.DB {
	return gdb.Create(&repartidor)
}

func ListRepartidores(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()
	defer database.Close(gdb)
	var repartidores []models.Repartidor

	result := listRepartidores(&repartidores, gdb)
	if result.Error != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(utils.ErrorMessage{ErrorMessage: result.Error.Error()})
		return
	} else {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode(repartidores)
		return
	}
}

func listRepartidores(repartidores *[]models.Repartidor, gdb *gorm.DB) *gorm.DB {
	return gdb.Model(&models.Repartidor{}).Order("ID asc").Find(&repartidores)
}

func GetRepartidor(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()
	defer database.Close(gdb)
	var repartidor models.Repartidor

	result := getRepartidor(&repartidor, params.ByName("id"), gdb)
	if result.Error != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(utils.ErrorMessage{ErrorMessage: result.Error.Error()})
		return
	} else if result.RowsAffected == 0 {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)
		json.NewEncoder(writer).Encode(utils.ErrorMessage{ErrorMessage: "No existe repartidor con id " + params.ByName("id") + "!"})
		return
	} else {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode(repartidor)
		return
	}
}

func getRepartidor(repartidor *models.Repartidor, id string, gdb *gorm.DB) *gorm.DB {
	return gdb.Model(&models.Repartidor{}).Order("ID asc").Find(&repartidor, id)
}

func PutRepartidor(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()
	defer database.Close(gdb)
	var repartidor models.Repartidor

	decoder := json.NewDecoder(request.Body)

	err := decoder.Decode(&repartidor)
	if err != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(utils.ErrorMessage{ErrorMessage: err.Error()})
		return
	}

	repartidor.ID, _ = strconv.ParseUint(params.ByName("id"), 10, 64)

	result := putRepartidor(&repartidor, gdb)
	if result.Error != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(utils.ErrorMessage{ErrorMessage: result.Error.Error()})
		return
	} else if result.RowsAffected == 0 {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)
		json.NewEncoder(writer).Encode(utils.ErrorMessage{ErrorMessage: "No existe repartidor con id " + params.ByName("id") + "!"})
		return
	} else {
		writer.WriteHeader(http.StatusOK)
		return
	}

}

func putRepartidor(repartidor *models.Repartidor, gdb *gorm.DB) *gorm.DB {
	return gdb.Save(&repartidor)
}

func PatchRepartidor(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()
	defer database.Close(gdb)

	var repartidor models.Repartidor

	decoder := json.NewDecoder(request.Body)

	err := decoder.Decode(&repartidor)
	if err != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(utils.ErrorMessage{ErrorMessage: err.Error()})
		return
	}

	repartidor.ID, _ = strconv.ParseUint(params.ByName("id"), 10, 64)

	result := patchRepartidor(&repartidor, gdb)
	if result.Error != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(utils.ErrorMessage{ErrorMessage: result.Error.Error()})
		return
	} else if result.RowsAffected == 0 {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)
		json.NewEncoder(writer).Encode(utils.ErrorMessage{ErrorMessage: "No existe repartidor con id " + params.ByName("id") + "!"})
		return
	} else {
		writer.WriteHeader(http.StatusOK)
		return
	}

}

func patchRepartidor(repartidor *models.Repartidor, gdb *gorm.DB) *gorm.DB {
	return gdb.Updates(&repartidor)
}

func DeleteRepartidor(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()
	database.Close(gdb)
	var repartidor models.Repartidor
	repartidor.ID, _ = strconv.ParseUint(params.ByName("id"), 10, 64)

	result := deleteRepartidor(&repartidor, false, gdb)
	if result.Error != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(utils.ErrorMessage{ErrorMessage: result.Error.Error()})
		return
	} else if result.RowsAffected == 0 {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)
		json.NewEncoder(writer).Encode(utils.ErrorMessage{ErrorMessage: "No existe repartidor con id " + params.ByName("id") + "!"})
		return
	} else {
		writer.WriteHeader(http.StatusOK)
		return
	}
}

func deleteRepartidor(repartidor *models.Repartidor, hard bool, gdb *gorm.DB) *gorm.DB {
	if hard {
		// Delete the record
		return gdb.Unscoped().Delete(&repartidor)
	} else {
		// Update the "deleted_at" column
		return gdb.Delete(&repartidor)
	}
}
