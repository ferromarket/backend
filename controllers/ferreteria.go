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

type Ferreterias struct {
    Ferreterias []models.Ferreteria `json:"Ferreterias"`
}

func PostFerreteria(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()

	decoder := json.NewDecoder(request.Body)

	var ferreteria models.Ferreteria

	err := decoder.Decode(&ferreteria)
	if (err != nil) {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(utils.ErrorMessage{ErrorMessage: err.Error()})
	}

	err = postFerreteria(ferreteria, gdb)
	if (err != nil) {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(utils.ErrorMessage{ErrorMessage: err.Error()})
	} else {
		writer.WriteHeader(http.StatusOK)
	}

	database.Close(gdb)
}

func postFerreteria(ferreteria models.Ferreteria, gdb *gorm.DB) error {
	return gdb.Create(&ferreteria).Error
	/*sql := gdb.ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.Create(&ferreteria)
	})

	fmt.Println(sql)*/
}

func ListFerreterias(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()

	ferreteriaList := Ferreterias{}
	var ferreterias []models.Ferreteria

	gdb.Model(&models.Ferreteria{}).Order("ID asc").Preload("Horarios.Dia").Preload("Horarios.Abrir").Preload("Horarios.Cerrar").Preload("Comuna.Ciudad.Region.Pais").Joins("LEFT JOIN ferreteria_horario fh ON ferreteria.id = fh.ferreteria_id").Find(&ferreterias)

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
