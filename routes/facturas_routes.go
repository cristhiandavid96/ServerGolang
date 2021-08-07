package routes

import (
	"github.com/cristhiandavid96/ServerGolang/controllers"
	"github.com/gorilla/mux"
)

func SetFacturasRoutes(router *mux.Router) {
	subRoute := router.PathPrefix("/factura/api").Subrouter()
	subRoute.HandleFunc("/all", controllers.GetAllFactura).Methods("GET")
	subRoute.HandleFunc("/save", controllers.SaveFactura).Methods("POST")
	subRoute.HandleFunc("/delete/{id}", controllers.DeleteFactura).Methods("POST")
	subRoute.HandleFunc("/find/{id}", controllers.GetFactura).Methods("GET")
}
