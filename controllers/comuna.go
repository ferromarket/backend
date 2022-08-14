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

func ListComunas(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()
	defer database.Close(gdb)

	var comunas []models.Comuna

	result := listComunas(&comunas, params.ByName("ciudad"), gdb)
	if (result.Error != nil) {
		utils.JSONErrorOutput(writer, http.StatusBadRequest, result.Error.Error())
		return
	} else {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode(comunas)
		return
	}
}

func listComunas(comunas *[]models.Comuna, ciudad string, gdb *gorm.DB) *gorm.DB {
	if ciudad != "" {
		return gdb.Model(&models.Comuna{}).Where("ciudad_id = ?", ciudad).Order("ID asc").Find(&comunas)
	} else {
		return gdb.Model(&models.Comuna{}).Order("ID asc").Find(&comunas)
	}
}

func GetComuna(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()
	defer database.Close(gdb)

	var comuna models.Comuna

	result := getComuna(&comuna, params.ByName("id"), gdb)
	if (result.Error != nil) {
		utils.JSONErrorOutput(writer, http.StatusBadRequest, result.Error.Error())
		return
	} else if (result.RowsAffected == 0) {
		utils.JSONErrorOutput(writer, http.StatusNotFound, "No existe comuna con id " + params.ByName("id") + "!")
		return
	} else {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode(comuna)
		return
	}
}

func getComuna(comuna *models.Comuna, id string, gdb *gorm.DB) *gorm.DB {
	return gdb.Model(&models.Comuna{}).Order("ID asc").Find(&comuna, id)
}
