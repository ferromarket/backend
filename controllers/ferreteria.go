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
	defer database.Close(gdb)

	decoder := json.NewDecoder(request.Body)

	var ferreteria models.Ferreteria

	err := decoder.Decode(&ferreteria)
	if (err != nil) {
		utils.JSONErrorOutput(writer, http.StatusBadRequest, err.Error())
		return
	}

	err = ferreteria.Validate()
	if (err != nil) {
		utils.JSONErrorOutput(writer, http.StatusBadRequest, err.Error())
		return
	}

	result := postFerreteria(ferreteria, gdb)
	if (result.Error != nil) {
		utils.JSONErrorOutput(writer, http.StatusBadRequest, result.Error.Error())
		return
	} else {
		writer.WriteHeader(http.StatusNoContent)
		return
	}
}

func postFerreteria(ferreteria models.Ferreteria, gdb *gorm.DB) *gorm.DB {
	return gdb.Create(&ferreteria)
}

func ListFerreterias(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()
	defer database.Close(gdb)

	var ferreterias []models.Ferreteria

	result := listFerreterias(&ferreterias, gdb)
	if (result.Error != nil) {
		utils.JSONErrorOutput(writer, http.StatusBadRequest, result.Error.Error())
		return
	} else {
		for ferreteria := range(ferreterias) {
			result := gdb.Model(&models.FerreteriaHorario{}).Order("ID asc").Preload("Abrir").Preload("Cerrar").Where("ferreteria_id = ?", ferreterias[ferreteria].ID).Find(&ferreterias[ferreteria].Horarios)
			if (result.Error != nil) {
				utils.JSONErrorOutput(writer, http.StatusBadRequest, result.Error.Error())
				return
			} 
		}
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode(ferreterias)
		return
	}
}

func listFerreterias(ferreterias *[]models.Ferreteria, gdb *gorm.DB) *gorm.DB {
	return gdb.Model(&models.Ferreteria{}).Order("ID asc").Preload("Comuna.Ciudad.Region.Pais").Find(&ferreterias)
}

func GetFerreteria(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()
	defer database.Close(gdb)

	var ferreteria models.Ferreteria

	result := getFerreteria(&ferreteria, params.ByName("id"), gdb)
	if (result.Error != nil) {
		utils.JSONErrorOutput(writer, http.StatusBadRequest, result.Error.Error())
		return
	} else if (result.RowsAffected == 0) {
		utils.JSONErrorOutput(writer, http.StatusNotFound, "No existe ferreteria con id " + params.ByName("id") + "!")
		return
	} else {
		result := gdb.Model(&models.FerreteriaHorario{}).Order("ID asc").Where("ferreteria_id = ?", params.ByName("id")).Find(&ferreteria.Horarios)
		if (result.Error != nil) {
			utils.JSONErrorOutput(writer, http.StatusBadRequest, result.Error.Error())
			return
		} else {
			writer.Header().Set("Content-Type", "application/json")
			writer.WriteHeader(http.StatusOK)
			json.NewEncoder(writer).Encode(ferreteria)
			return
		}
	}
}

func getFerreteria(ferreteria *models.Ferreteria, id string, gdb *gorm.DB) *gorm.DB {
	return gdb.Model(&models.Ferreteria{}).Order("ID asc").Preload("Comuna.Ciudad.Region.Pais").Find(&ferreteria, id)
}

func PutFerreteria(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()
	defer database.Close(gdb)

	var ferreteria models.Ferreteria

	decoder := json.NewDecoder(request.Body)

	err := decoder.Decode(&ferreteria)
	if (err != nil) {
		utils.JSONErrorOutput(writer, http.StatusBadRequest, err.Error())
		return
	}

	err = ferreteria.Validate()
	if (err != nil) {
		utils.JSONErrorOutput(writer, http.StatusBadRequest, err.Error())
		return
	}

	ferreteria.ID, _ = strconv.ParseUint(params.ByName("id"), 10, 64)

	result := putFerreteria(&ferreteria, gdb)
	if (result.Error != nil) {
		utils.JSONErrorOutput(writer, http.StatusBadRequest, result.Error.Error())
		return
	} else if (result.RowsAffected == 0) {
		utils.JSONErrorOutput(writer, http.StatusNotFound, "No existe ferreteria con id " + params.ByName("id") + "!")
		return
	} else {
		writer.WriteHeader(http.StatusNoContent)
		return
	}
}

func putFerreteria(ferreteria *models.Ferreteria, gdb *gorm.DB) *gorm.DB {
	return gdb.Save(&ferreteria)
}

func PatchFerreteria(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()
	defer database.Close(gdb)

	var ferreteria models.Ferreteria

	decoder := json.NewDecoder(request.Body)

	err := decoder.Decode(&ferreteria)
	if (err != nil) {
		utils.JSONErrorOutput(writer, http.StatusBadRequest, err.Error())
		return
	}

	err = ferreteria.Validate()
	if (err != nil) {
		utils.JSONErrorOutput(writer, http.StatusBadRequest, err.Error())
		return
	}

	ferreteria.ID, _ = strconv.ParseUint(params.ByName("id"), 10, 64)

	result := patchFerreteria(&ferreteria, gdb)
	if (result.Error != nil) {
		utils.JSONErrorOutput(writer, http.StatusBadRequest, result.Error.Error())
		return
	} else if (result.RowsAffected == 0) {
		utils.JSONErrorOutput(writer, http.StatusNotFound, "No existe ferreteria con id " + params.ByName("id") + "!")
		return
	} else {
		for _, horario := range ferreteria.Horarios {
			gdb.Updates(&horario)
		}
		writer.WriteHeader(http.StatusNoContent)
		return
	}
}

func patchFerreteria(ferreteria *models.Ferreteria, gdb *gorm.DB) *gorm.DB {
	return gdb.Updates(&ferreteria)
}

func DeleteFerreteria(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()
	defer database.Close(gdb)

	var ferreteria models.Ferreteria
	ferreteria.ID, _ = strconv.ParseUint(params.ByName("id"), 10, 64)

	result := deleteFerreteria(&ferreteria, false, gdb)
	if (result.Error != nil) {
		utils.JSONErrorOutput(writer, http.StatusBadRequest, result.Error.Error())
		return
	} else if (result.RowsAffected == 0) {
		utils.JSONErrorOutput(writer, http.StatusNotFound, "No existe ferreteria con id " + params.ByName("id") + "!")
		return
	} else {
		writer.WriteHeader(http.StatusNoContent)
		return
	}
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
