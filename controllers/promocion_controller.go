package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/cristhiandavid96/ServerGolang/commons"
	"github.com/cristhiandavid96/ServerGolang/models"
	"github.com/gorilla/mux"
)

func GetAll(writer http.ResponseWriter, request *http.Request) {
	promociones := []models.Promocion{}
	db := commons.GetConnection()
	defer db.Close()

	db.Find(&promociones)
	json, _ := json.Marshal(promociones)
	commons.SendReponse(writer, http.StatusOK, json)
}

func Get(writer http.ResponseWriter, request *http.Request) {
	promocion := models.Promocion{}

	id := mux.Vars(request)["id"]

	db := commons.GetConnection()
	defer db.Close()

	db.Find(&promocion, id)

	if promocion.ID > 0 {
		json, _ := json.Marshal(promocion)
		commons.SendReponse(writer, http.StatusOK, json)
	} else {
		commons.SendError(writer, http.StatusNotFound)
	}

}

func Save(writer http.ResponseWriter, request *http.Request) {
	promocion := models.Promocion{}

	db := commons.GetConnection()
	defer db.Close()

	error := json.NewDecoder(request.Body).Decode(&promocion)

	if error != nil {
		log.Fatal(error)
		commons.SendError(writer, http.StatusBadRequest)
		return
	}

	error = db.Save(&promocion).Error

	if error != nil {
		log.Fatal(error)
		commons.SendError(writer, http.StatusInternalServerError)
		return
	}

	json, _ := json.Marshal(promocion)

	commons.SendReponse(writer, http.StatusCreated, json)
}

func Delete(writer http.ResponseWriter, request *http.Request) {
	promocion := models.Promocion{}

	db := commons.GetConnection()
	defer db.Close()

	id := mux.Vars(request)["id"]

	db.Find(&promocion, id)

	if promocion.ID > 0 {
		db.Delete(promocion)
		commons.SendReponse(writer, http.StatusOK, []byte(`{}`))
	} else {
		commons.SendError(writer, http.StatusNotFound)
	}
}
