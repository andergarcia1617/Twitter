package handles

import (
	"log"
	"net/http"
	"os"

	"github.com/andergarcia1617/Twitter/middlew"
	"github.com/andergarcia1617/Twitter/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/*Manejadores setear el puerto, el Handler y pongo a escuchar a mi servidor */
func Manejadores() {
	router := mux.NewRouter()
	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
