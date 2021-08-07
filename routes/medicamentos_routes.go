package routes

import (
	"github.com/cristhiandavid96/ServerGolang/controllers"
	"github.com/gorilla/mux"
)

func SetMedicamentosRoutes(router *mux.Router) {
	subRoute := router.PathPrefix("/medicamento/api").Subrouter()
	subRoute.HandleFunc("/all", controllers.GetAllMedicamento).Methods("GET")
	subRoute.HandleFunc("/save", controllers.SaveMedicamento).Methods("POST")
	subRoute.HandleFunc("/delete/{id}", controllers.DeleteMedicamento).Methods("POST")
	subRoute.HandleFunc("/find/{id}", controllers.GetMedicamento).Methods("GET")
}