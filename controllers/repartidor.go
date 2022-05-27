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

type Repartidores struct {
	Repartidores []models.Repartidor `json:"Repartidores"`
}

func PostRepartidor(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()

	decoder := json.NewDecoder(request.Body)

	var repartidor models.Repartidor

	err := decoder.Decode(&repartidor)
	if err != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(utils.ErrorMessage{ErrorMessage: err.Error()})
	}

	err = postRepartidor(repartidor, gdb)
	if err != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(utils.ErrorMessage{ErrorMessage: err.Error()})
	} else {
		writer.WriteHeader(http.StatusOK)
	}

	database.Close(gdb)
}

func postRepartidor(repartidor models.Repartidor, gdb *gorm.DB) error {
	return gdb.Create(&repartidor).Error

}

func ListRepartidores(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()

	repartidorList := Repartidores{}
	var repartidores []models.Repartidor

	//gdb.Model(&models.Ferreteria{}).Order("ID asc").Preload("Horarios.Dia").Preload("Horarios.Abrir").Preload("Horarios.Cerrar").Preload("Comuna.Ciudad.Region.Pais").Joins("LEFT JOIN ferreteria_horario fh ON ferreteria.id = fh.ferreteria_id").Find(&ferreterias)
	gdb.Model(&models.Repartidor{}).Order("ID asc").Find(&repartidores)

	repartidorList.Repartidores = repartidores

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(repartidorList)

	database.Close(gdb)
}

func GetRepartidor(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

}

func PutRepartidor(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

}

func PatchRepartidor(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

}

func DeleteRepartidor(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

}
