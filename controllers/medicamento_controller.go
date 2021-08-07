package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/cristhiandavid96/ServerGolang/commons"
	"github.com/cristhiandavid96/ServerGolang/models"
	"github.com/gorilla/mux"
)

func GetAllMedicamento(writer http.ResponseWriter, request *http.Request) {
	medicamentos := []models.Medicamento{}
	db := commons.GetConnection()
	defer db.Close()
	db.Find(&medicamentos)
	json, _ := json.Marshal(medicamentos)
	commons.SendReponse(writer, http.StatusOK, json)
}

func GetMedicamento(writer http.ResponseWriter, request *http.Request) {
	medicamento := models.Medicamento{}

	id := mux.Vars(request)["id"]

	db := commons.GetConnection()
	defer db.Close()

	db.Find(&medicamento, id)

	if medicamento.ID > 0 {
		json, _ := json.Marshal(medicamento)
		commons.SendReponse(writer, http.StatusOK, json)
	} else {
		commons.SendError(writer, http.StatusNotFound)
	}

}

func SaveMedicamento(writer http.ResponseWriter, request *http.Request) {
	medicamento := models.Medicamento{}

	db := commons.GetConnection()
	defer db.Close()

	error := json.NewDecoder(request.Body).Decode(&medicamento)

	if error != nil {
		log.Fatal(error)
		commons.SendError(writer, http.StatusBadRequest)
		return
	}

	error = db.Save(&medicamento).Error

	if error != nil {
		log.Fatal(error)
		commons.SendError(writer, http.StatusInternalServerError)
		return
	}

	json, _ := json.Marshal(medicamento)

	commons.SendReponse(writer, http.StatusCreated, json)
}

func DeleteMedicamento(writer http.ResponseWriter, request *http.Request) {
	medicamento := models.Medicamento{}

	db := commons.GetConnection()
	defer db.Close()

	id := mux.Vars(request)["id"]

	db.Find(&medicamento, id)

	if medicamento.ID > 0 {
		db.Delete(medicamento)
		commons.SendReponse(writer, http.StatusOK, []byte(`{}`))
	} else {
		commons.SendError(writer, http.StatusNotFound)
	}
}
