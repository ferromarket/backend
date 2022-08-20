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

func ListRegiones(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()
	defer database.Close(gdb)

	var regiones []models.Region

	result := listRegiones(&regiones, params.ByName("pais"), gdb)
	if (result.Error != nil) {
		utils.JSONErrorOutput(writer, http.StatusBadRequest, result.Error.Error())
		return
	} else {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode(regiones)
		return
	}
}

func listRegiones(regiones *[]models.Region, pais string, gdb *gorm.DB) *gorm.DB {
	if pais != "" {
		return gdb.Model(&models.Region{}).Where("pais_id = ?", pais).Order("ID asc").Find(&regiones)
	} else {
		return gdb.Model(&models.Region{}).Order("ID asc").Find(&regiones)
	}
}

func GetRegion(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()
	defer database.Close(gdb)

	var region models.Region

	result := getRegion(&region, params.ByName("id"), gdb)
	if (result.Error != nil) {
		utils.JSONErrorOutput(writer, http.StatusBadRequest, result.Error.Error())
		return
	} else if (result.RowsAffected == 0) {
		utils.JSONErrorOutput(writer, http.StatusNotFound, "No existe región con id " + params.ByName("id") + "!")
		return
	} else {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode(region)
		return
	}
}

func getRegion(region *models.Region, id string, gdb *gorm.DB) *gorm.DB {
	return gdb.Model(&models.Region{}).Order("ID asc").Find(&region, id)
}
