package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/ferromarket/backend/models"
	"github.com/julienschmidt/httprouter"
	"gorm.io/gorm"
)

type Ferreterias struct {
    Ferreterias []models.Ferreteria `json:"ferreterias"`
}

func PostFerreteria(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	db := dbConnect()
	db.AutoMigrate(&models.Ferreteria{})

	postFerreteria(models.Ferreteria{Name: "Chris's hardware"}, db)

	dbClose(db)
}

func postFerreteria(ferreteria models.Ferreteria, db *gorm.DB) {
	db.Create(&ferreteria)
}

func ListFerreterias(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	db := dbConnect()

	ferreteriaList := Ferreterias{}
	var ferreterias []models.Ferreteria

	db.Order("ID asc").Find(&ferreterias)

	ferreteriaList.Ferreterias = ferreterias

    writer.Header().Set("Content-Type", "application/json")
    writer.WriteHeader(http.StatusOK)
    json.NewEncoder(writer).Encode(ferreteriaList)

	dbClose(db)
}

func GetFerreteria(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

}

func PutFerreteria(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

}

func PatchFerreteria(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

}

func DeleteFerreteria(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

}
