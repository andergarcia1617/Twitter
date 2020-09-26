package handles

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/*Manejadores setear el puerto, el Handler y pongo a escuchar a mi servidor */
func Manejadores() {
	router := mux.NewRouter()
	PORT := os.Getenv("PORT")
	if PORT == " " {
		PORT = "8080 "
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
