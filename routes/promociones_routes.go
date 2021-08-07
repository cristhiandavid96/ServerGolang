package routes

import (
	"github.com/cristhiandavid96/ServerGolang/controllers"
	"github.com/gorilla/mux"
)



func SetPromocionesRoutes(router *mux.Router) {
	subRoute := router.PathPrefix("/promocion/api").Subrouter()
	subRoute.HandleFunc("/all", controllers.GetAllPromocion).Methods("GET")
	subRoute.HandleFunc("/save", controllers.SavePromocion).Methods("POST")
	subRoute.HandleFunc("/delete/{id}", controllers.DeletePromocion).Methods("POST")
	subRoute.HandleFunc("/find/{id}", controllers.GetPromocion).Methods("GET")
}
