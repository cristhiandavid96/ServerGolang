package main

import (
	"log"
	"net/http"

	"github.com/cristhiandavid96/ServerGolang/commons"
	"github.com/cristhiandavid96/ServerGolang/routes"
	"github.com/gorilla/mux"
)

func main() {
	commons.Migrate()

	router := mux.NewRouter()
	routes.SetPersonaRoutes(router)
	commons.EnableCORS(router)

	server := http.Server{
		Addr:    ":9000",
		Handler: router,
	}

	log.Println("Servidor ejecutandose sobre el puerto 9000")
	log.Println(server.ListenAndServe())
}
