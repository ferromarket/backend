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
	gdb := database.Connect()
	var producto models.Producto
	result := getProducto(&producto, params.ByName("id"), gdb)
	if result.Error != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(utils.ErrorMessage{ErrorMessage: result.Error.Error()})
	} else if result.RowsAffected == 0 {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)
		json.NewEncoder(writer).Encode(utils.ErrorMessage{ErrorMessage: "No existe producto con id " + params.ByName("id") + "!"})
	} else {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode(producto)
	}

	database.Close(gdb)
}
func getProducto(Producto *models.Producto, id string, gdb *gorm.DB) *gorm.DB {
	return gdb.Model(&models.Producto{}).Order("ID asc").Find(&Producto, id)
}

func PutProducto(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()

	var producto models.Producto

	decoder := json.NewDecoder(request.Body)

	err := decoder.Decode(&producto)
	if err != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(utils.ErrorMessage{ErrorMessage: err.Error()})
	}

	producto.ID, _ = strconv.ParseUint(params.ByName("id"), 10, 64)

	result := putProducto(&producto, gdb)
	if result.Error != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(utils.ErrorMessage{ErrorMessage: result.Error.Error()})
	} else if result.RowsAffected == 0 {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)
		json.NewEncoder(writer).Encode(utils.ErrorMessage{ErrorMessage: "No existe producto con id " + params.ByName("id") + "!"})
	} else {
		writer.WriteHeader(http.StatusOK)
	}

	database.Close(gdb)
}

func putProducto(producto *models.Producto, gdb *gorm.DB) *gorm.DB {
	return gdb.Save(&producto)
}

func PatchProducto(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()

	var producto models.Producto

	decoder := json.NewDecoder(request.Body)

	err := decoder.Decode(&producto)
	if err != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(utils.ErrorMessage{ErrorMessage: err.Error()})
	}

	producto.ID, _ = strconv.ParseUint(params.ByName("id"), 10, 64)

	result := patchProducto(&producto, gdb)
	if result.Error != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(utils.ErrorMessage{ErrorMessage: result.Error.Error()})
	} else if result.RowsAffected == 0 {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)
		json.NewEncoder(writer).Encode(utils.ErrorMessage{ErrorMessage: "No existe producto con id " + params.ByName("id") + "!"})
	} else {
		writer.WriteHeader(http.StatusOK)
	}

	database.Close(gdb)
}

func patchProducto(producto *models.Producto, gdb *gorm.DB) *gorm.DB {
	return gdb.Updates(&producto)
}

func DeleteProducto(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()
	var producto models.Producto
	producto.ID, _ = strconv.ParseUint(params.ByName("ID"), 10, 64)

	result := deleteProducto(&producto, false, gdb)
	if result.Error != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(utils.ErrorMessage{ErrorMessage: result.Error.Error()})
	} else if result.RowsAffected == 0 {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)
		json.NewEncoder(writer).Encode(utils.ErrorMessage{ErrorMessage: "No existe producto con id " + params.ByName("id") + "!"})
	} else {
		writer.WriteHeader(http.StatusOK)
	}

	database.Close(gdb)

}

func deleteProducto(Producto *models.Producto, hard bool, gdb *gorm.DB) *gorm.DB {
	if hard {
		// Delete the record
		return gdb.Unscoped().Delete(&Producto)
	} else {
		// Update the "deleted_at" column
		return gdb.Delete(&Producto)
	}
}
