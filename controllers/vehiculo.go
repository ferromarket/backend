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

func PostVehiculo(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()

	decoder := json.NewDecoder(request.Body)

	var vehiculo models.Vehiculo

	err := decoder.Decode(&vehiculo)
	if err != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(utils.ErrorMessage{ErrorMessage: err.Error()})
	}

	result := postVehiculo(vehiculo, gdb)
	if result.Error != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(utils.ErrorMessage{ErrorMessage: result.Error.Error()})
	} else {
		writer.WriteHeader(http.StatusOK)
	}

	database.Close(gdb)
}

func postVehiculo(vehiculo models.Vehiculo, gdb *gorm.DB) *gorm.DB {
	return gdb.Create(&vehiculo)
}

func ListVehiculos(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()

	var Vehiculos []models.Vehiculo

	result := listVehiculos(&Vehiculos, gdb)
	if result.Error != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(utils.ErrorMessage{ErrorMessage: result.Error.Error()})
	} else {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode(Vehiculos)
	}

	database.Close(gdb)
}

func listVehiculos(Vehiculos *[]models.Vehiculo, gdb *gorm.DB) *gorm.DB {
	return gdb.Model(&models.Vehiculo{}).Order("ID asc").Preload("Repartidor").Find(&Vehiculos)
}

func GetVehiculo(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()

	var vehiculo models.Vehiculo

	result := getVehiculo(&vehiculo, params.ByName("id"), gdb)
	if result.Error != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(utils.ErrorMessage{ErrorMessage: result.Error.Error()})
	} else if result.RowsAffected == 0 {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)
		json.NewEncoder(writer).Encode(utils.ErrorMessage{ErrorMessage: "No existe vehiculo con id " + params.ByName("id") + "!"})
	} else {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode(vehiculo)
	}

	database.Close(gdb)
}

func getVehiculo(vehiculo *models.Vehiculo, id string, gdb *gorm.DB) *gorm.DB {
	return gdb.Model(&models.Vehiculo{}).Order("ID asc").Preload("Repartidor").Find(&vehiculo, id)
}

func PutVehiculo(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()

	var vehiculo models.Vehiculo

	decoder := json.NewDecoder(request.Body)

	err := decoder.Decode(&vehiculo)
	if err != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(utils.ErrorMessage{ErrorMessage: err.Error()})
	}

	vehiculo.ID, _ = strconv.ParseUint(params.ByName("id"), 10, 64)

	result := putVehiculo(&vehiculo, gdb)
	if result.Error != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(utils.ErrorMessage{ErrorMessage: result.Error.Error()})
	} else if result.RowsAffected == 0 {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)
		json.NewEncoder(writer).Encode(utils.ErrorMessage{ErrorMessage: "No existe vehiculo con id " + params.ByName("id") + "!"})
	} else {
		writer.WriteHeader(http.StatusOK)
	}

	database.Close(gdb)
}

func putVehiculo(vehiculo *models.Vehiculo, gdb *gorm.DB) *gorm.DB {
	return gdb.Save(&vehiculo)
}

func PatchVehiculo(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()

	var vehiculo models.Vehiculo

	decoder := json.NewDecoder(request.Body)

	err := decoder.Decode(&vehiculo)
	if err != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(utils.ErrorMessage{ErrorMessage: err.Error()})
	}

	vehiculo.ID, _ = strconv.ParseUint(params.ByName("id"), 10, 64)

	result := patchVehiculo(&vehiculo, gdb)
	if result.Error != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(utils.ErrorMessage{ErrorMessage: result.Error.Error()})
	} else if result.RowsAffected == 0 {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)
		json.NewEncoder(writer).Encode(utils.ErrorMessage{ErrorMessage: "No existe vehiculo con id " + params.ByName("id") + "!"})
	} else {
		writer.WriteHeader(http.StatusOK)
	}

	database.Close(gdb)
}

func patchVehiculo(vehiculo *models.Vehiculo, gdb *gorm.DB) *gorm.DB {
	return gdb.Updates(&vehiculo)
}

func DeleteVehiculo(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()

	var vehiculo models.Vehiculo
	vehiculo.ID, _ = strconv.ParseUint(params.ByName("id"), 10, 64)

	result := deleteVehiculo(&vehiculo, false, gdb)
	if result.Error != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(utils.ErrorMessage{ErrorMessage: result.Error.Error()})
	} else if result.RowsAffected == 0 {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)
		json.NewEncoder(writer).Encode(utils.ErrorMessage{ErrorMessage: "No existe vehiculo con id " + params.ByName("id") + "!"})
	} else {
		writer.WriteHeader(http.StatusOK)
	}

	database.Close(gdb)
}

func deleteVehiculo(vehiculo *models.Vehiculo, hard bool, gdb *gorm.DB) *gorm.DB {
	if hard {
		// Delete the record
		return gdb.Unscoped().Delete(&vehiculo)
	} else {
		// Update the "deleted_at" column
		return gdb.Delete(&vehiculo)
	}
}
