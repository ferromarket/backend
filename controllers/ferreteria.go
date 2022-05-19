package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/ferromarket/backend/database"
	"github.com/ferromarket/backend/models"
	"github.com/julienschmidt/httprouter"
	"gorm.io/gorm"
)

type Ferreterias struct {
    Ferreterias []models.Ferreteria `json:"Ferreterias"`
}

func PostFerreteria(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()

	postFerreteria(models.Ferreteria{Nombre: params.ByName("nombre")}, gdb)

	database.Close(gdb)
}

func postFerreteria(ferreteria models.Ferreteria, gdb *gorm.DB) {
	gdb.Create(&ferreteria)
}

func ListFerreterias(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()

	ferreteriaList := Ferreterias{}
	var ferreterias []models.Ferreteria

	gdb.Model(&models.Ferreteria{}).Order("ID asc").Preload("Horarios.Dia").Preload("Horarios.Abrir").Preload("Horarios.Cerrar").Preload("Comuna.Ciudad.Region.Pais").Joins("JOIN ferreteria_horario fh ON ferreteria.id = fh.ferreteria_id").Find(&ferreterias)

	ferreteriaList.Ferreterias = ferreterias

    writer.Header().Set("Content-Type", "application/json")
    writer.WriteHeader(http.StatusOK)
    json.NewEncoder(writer).Encode(ferreteriaList)

	database.Close(gdb)
}

func GetFerreteria(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

}

func PutFerreteria(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

}

func PatchFerreteria(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

}

func DeleteFerreteria(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

}
