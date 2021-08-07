package routes

import (
	"github.com/cristhiandavid96/ServerGolang/controllers"
	"github.com/gorilla/mux"
)



func SetPromocionesRoutes(router *mux.Router) {
	subRoute := router.PathPrefix("/promocion/api").Subrouter()
	subRoute.HandleFunc("/all", controllers.GetAll).Methods("GET")
	subRoute.HandleFunc("/save", controllers.Save).Methods("POST")
	subRoute.HandleFunc("/delete/{id}", controllers.Delete).Methods("POST")
	subRoute.HandleFunc("/find/{id}", controllers.Get).Methods("GET")
}
