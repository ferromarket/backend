package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/ferromarket/backend/models"
	"github.com/julienschmidt/httprouter"
)

type Ferreterias struct {
    Ferreteria []models.Ferreteria `json:"ferreterias"`
}

func PostFerreteria(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	db := dbConnect()
	db.AutoMigrate(&models.Ferreteria{})

	db.Create(&models.Ferreteria{Name: "Chris's hardware"})

	dbClose(db)
}

func ListFerreterias(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	ferreteriaList := Ferreterias{}
	db := dbConnect()
	var ferreterias []models.Ferreteria

	db.Order("ID asc").Find(&ferreterias)

	ferreteriaList.Ferreteria = ferreterias

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
