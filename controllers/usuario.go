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

type Usuarios struct {
	Usuarios []models.Usuario `json:"Usuario"`
}

func PostUsuario(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()

	decoder := json.NewDecoder(request.Body)

	var usuario models.Usuario

	err := decoder.Decode(&usuario)
	if err != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(utils.ErrorMessage{ErrorMessage: err.Error()})
	}

	err = postUsuario(usuario, gdb)
	if err != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(utils.ErrorMessage{ErrorMessage: err.Error()})
	} else {
		writer.WriteHeader(http.StatusOK)
	}

	database.Close(gdb)
}

func postUsuario(usuario models.Usuario, gdb *gorm.DB) error {
	return gdb.Create(&usuario).Error
	/*sql := gdb.ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.Create(&ferreteria)
	})

	fmt.Println(sql)*/
}

func putUsuario(usuario models.Usuario, gdb *gorm.DB) error{
	return gdb.Save(&usuario).Error
}

func patchUsuario(usuario models.Usuario, gdb *gorm.DB) error{
	return gdb.Updates(&usuario).Error
}


func ListUsuarios(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()

	usuarioList := Usuarios{}
	var usuarios []models.Usuario

	gdb.Model(&models.Usuario{}).Order("ID asc").Preload("Rol").Joins("LEFT JOIN usuario_rol ur ON usuario.id = ur.usuario_id").Find(&usuarios)

	usuarioList.Usuarios = usuarios

	for i := range usuarioList.Usuarios {
		usuarioList.Usuarios[i].Contrasena = nil
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(usuarioList)

	database.Close(gdb)
}

func GetUsuario(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()

	var usuario models.Usuario

	gdb.Model(&models.Usuario{}).Order("ID asc").Preload("Rol").Find(&usuario, params.ByName("id"))

	usuario.Contrasena = nil

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(usuario)

	database.Close(gdb)
}

// insertar o update. Necesita el objeto completo. Todos los atributos
func PutUsuario(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()

	decoder := json.NewDecoder(request.Body)

	var usuario models.Usuario

	err := decoder.Decode(&usuario)
	if err != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(utils.ErrorMessage{ErrorMessage: err.Error()})
	}

	usuario.ID, _ = strconv.ParseUint(params.ByName("id"), 10, 64)

	err = putUsuario(usuario, gdb)
	if err != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(utils.ErrorMessage{ErrorMessage: err.Error()})
	} else {
		writer.WriteHeader(http.StatusOK)
	}

	database.Close(gdb)
}

// func hashClave (){

// }

// solo update solo para un atributo
func PatchUsuario(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// hashClave()
	gdb := database.Connect()

	decoder := json.NewDecoder(request.Body)

	var usuario models.Usuario

	err := decoder.Decode(&usuario)
	if err != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(utils.ErrorMessage{ErrorMessage: err.Error()})
	}

	usuario.ID, _ = strconv.ParseUint(params.ByName("id"), 10, 64)

	var usuario_verif models.Usuario


	result := gdb.Find(&usuario_verif, usuario.ID)

	if result.RowsAffected == 0 {
		writer.WriteHeader(http.StatusNotFound)
	}else{
		err = patchUsuario(usuario, gdb)
		if err != nil {
			writer.Header().Set("Content-Type", "application/json")
			writer.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(writer).Encode(utils.ErrorMessage{ErrorMessage: err.Error()})
		} else {
			writer.WriteHeader(http.StatusOK)
		}
	}

	database.Close(gdb)

}

func DeleteUsuario(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	gdb := database.Connect()

	var usuario models.Usuario
	var err error

	if usuario.ID, err = strconv.ParseUint(params.ByName("id"), 10, 64); err != nil {
		panic(err)
	} else {
		gdb.Delete(&usuario)
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(usuario)

	database.Close(gdb)
}


