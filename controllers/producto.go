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

type Productos struct {
	Productos []models.Producto `json:"Producto"`
}

func PostProducto(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()

	decoder := json.NewDecoder(request.Body)

	var producto models.Producto

	err := decoder.Decode(&producto)
	if err != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(utils.ErrorMessage{ErrorMessage: err.Error()})
	}

	err = postProducto(producto, gdb)
	if err != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(utils.ErrorMessage{ErrorMessage: err.Error()})
	} else {
		writer.WriteHeader(http.StatusOK)
	}
	database.Close(gdb)
}

func postProducto(producto models.Producto, gdb *gorm.DB) error {
	return gdb.Create(&producto).Error
}

func ListProductos(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()

	var productos []models.Producto

	result := listProductos(&productos, gdb)

	if result.Error != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(utils.ErrorMessage{ErrorMessage: result.Error.Error()})
	} else {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode(productos)
	}
	database.Close(gdb)
}

func listProductos(productos *[]models.Producto, gdb *gorm.DB) *gorm.DB {
	result := gdb.Model(&models.Producto{}).Order("ID asc").Preload("Categoria").Find(&productos)
	if result.Error != nil {
		return result
	}

	for indice := range *productos {
		(*productos)[indice].Categoria.Categoria = new(models.Categoria)

		categoria := (*productos)[indice].Categoria.Categoria
		categoriaID := *((*productos)[indice].Categoria.CategoriaID)

		result = getCategorias(categoriaID, categoria, gdb)
	}

	return result
}

func getCategorias(categoriaID uint64, categoria *models.Categoria, gdb *gorm.DB) *gorm.DB {
	var result *gorm.DB
	for ok := true; ok; ok = !(categoriaID == 0) {
		var nuevaCategoria models.Categoria
		result = gdb.Model(&models.Categoria{}).Preload("Categoria").Find(&nuevaCategoria, categoriaID)
		if result.Error != nil {
			return result
		}

		// Creacion de lista enlazada para las categorias padre
		*categoria = nuevaCategoria
		if nuevaCategoria.CategoriaID != nil {
			categoria.Categoria = new(models.Categoria)
			categoria = categoria.Categoria
			categoriaID = *nuevaCategoria.CategoriaID
		} else {
			categoriaID = 0
		}
	}
	return result
}

func GetProducto(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

}

func PutProducto(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

}

func PatchProducto(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

}

func DeleteProducto(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

}
