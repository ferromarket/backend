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

func ListCategorias(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()
	defer database.Close(gdb)

	var categorias []models.Categoria

	result := listCategorias(&categorias, gdb)
	if result.Error != nil {
		utils.JSONErrorOutput(writer, http.StatusBadRequest, result.Error.Error())
		return
	} else {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode(categorias)
		return
	}
}

func listCategorias(categorias *[]models.Categoria, gdb *gorm.DB) *gorm.DB {
	return gdb.Model(&models.Categoria{}).Order("ID asc").Find(&categorias)
}

func GetCategoria(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()
	defer database.Close(gdb)

	var categoria models.Categoria

	result := getCategoria(&categoria, params.ByName("id"), gdb)
	if result.Error != nil {
		utils.JSONErrorOutput(writer, http.StatusBadRequest, result.Error.Error())
		return
	} else if result.RowsAffected == 0 {
		utils.JSONErrorOutput(writer, http.StatusNotFound, "No existe categoría con id "+params.ByName("id")+"!")
		return
	} else {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode(categoria)
		return
	}
}

func getCategoria(categoria *models.Categoria, id string, gdb *gorm.DB) *gorm.DB {
	return gdb.Model(&models.Categoria{}).Order("ID asc").Find(&categoria, id)
}
