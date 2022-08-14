package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/ferromarket/backend/database"
	"github.com/ferromarket/backend/models"
	"github.com/ferromarket/backend/utils"
	"github.com/julienschmidt/httprouter"
	"gorm.io/gorm"
)

func ListCiudades(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()
	defer database.Close(gdb)

	var ciudades []models.Ciudad

	result := listCiudades(&ciudades, params.ByName("region"), gdb)
	if (result.Error != nil) {
		utils.JSONErrorOutput(writer, http.StatusBadRequest, result.Error.Error())
		return
	} else {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode(ciudades)
		return
	}
}

func listCiudades(ciudades *[]models.Ciudad, region string, gdb *gorm.DB) *gorm.DB {
	if region != "" {
		return gdb.Model(&models.Ciudad{}).Where("region_id = ?", region).Order("ID asc").Find(&ciudades)
	} else {
		return gdb.Model(&models.Ciudad{}).Order("ID asc").Find(&ciudades)
	}
}

func GetCiudad(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()
	defer database.Close(gdb)

	var ciudad models.Ciudad

	result := getCiudad(&ciudad, params.ByName("id"), gdb)
	if (result.Error != nil) {
		utils.JSONErrorOutput(writer, http.StatusBadRequest, result.Error.Error())
		return
	} else if (result.RowsAffected == 0) {
		utils.JSONErrorOutput(writer, http.StatusNotFound, "No existe ciudad con id " + params.ByName("id") + "!")
		return
	} else {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode(ciudad)
		return
	}
}

func getCiudad(ciudad *models.Ciudad, id string, gdb *gorm.DB) *gorm.DB {
	return gdb.Model(&models.Ciudad{}).Order("ID asc").Find(&ciudad, id)
}
