package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/ferromarket/backend/models"
	"github.com/julienschmidt/httprouter"
	"gorm.io/gorm"
)

type Ferreterias struct {
    Ferreterias []models.Ferreteria `json:"Ferreterias"`
}

func PostFerreteria(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	db := DBConnect()

	postFerreteria(models.Ferreteria{Nombre: params.ByName("nombre")}, db)

	DBClose(db)
}

func postFerreteria(ferreteria models.Ferreteria, db *gorm.DB) {
	db.Create(&ferreteria)
}

func ListFerreterias(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	db := DBConnect()

	ferreteriaList := Ferreterias{}
	var ferreterias []models.Ferreteria

	db.Model(&models.Ferreteria{}).Order("ID asc").Preload("Horarios.Dia").Preload("Horarios.Abrir").Preload("Horarios.Cerrar").Preload("Comuna.Ciudad.Region.Pais").Joins("JOIN ferreteria_horario fh ON ferreteria.id = fh.ferreteria_id").Find(&ferreterias)

	ferreteriaList.Ferreterias = ferreterias

    writer.Header().Set("Content-Type", "application/json")
    writer.WriteHeader(http.StatusOK)
    json.NewEncoder(writer).Encode(ferreteriaList)

	DBClose(db)
}

func GetFerreteria(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

}

func PutFerreteria(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

}

func PatchFerreteria(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

}

func DeleteFerreteria(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

}
