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
	facturas := []models.Factura{}
	db := commons.GetConnection()
	defer db.Close()

	db.Find(&facturas)
	json, _ := json.Marshal(facturas)
	commons.SendReponse(writer, http.StatusOK, json)
}

func Get(writer http.ResponseWriter, request *http.Request) {
	factura := models.Factura{}

	id := mux.Vars(request)["id"]

	db := commons.GetConnection()
	defer db.Close()

	db.Find(&factura, id)

	if factura.ID > 0 {
		json, _ := json.Marshal(factura)
		commons.SendReponse(writer, http.StatusOK, json)
	} else {
		commons.SendError(writer, http.StatusNotFound)
	}

}

func Save(writer http.ResponseWriter, request *http.Request) {
	factura := models.Factura{}

	db := commons.GetConnection()
	defer db.Close()

	error := json.NewDecoder(request.Body).Decode(&factura)

	if error != nil {
		log.Fatal(error)
		commons.SendError(writer, http.StatusBadRequest)
		return
	}

	error = db.Save(&factura).Error

	if error != nil {
		log.Fatal(error)
		commons.SendError(writer, http.StatusInternalServerError)
		return
	}

	json, _ := json.Marshal(factura)

	commons.SendReponse(writer, http.StatusCreated, json)
}

func Delete(writer http.ResponseWriter, request *http.Request) {
	factura := models.Factura{}

	db := commons.GetConnection()
	defer db.Close()

	id := mux.Vars(request)["id"]

	db.Find(&factura, id)

	if factura.ID > 0 {
		db.Delete(factura)
		commons.SendReponse(writer, http.StatusOK, []byte(`{}`))
	} else {
		commons.SendError(writer, http.StatusNotFound)
	}
}
