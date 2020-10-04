package middlew

import (
	"net/http"

	"github.com/andergarcia1617/Twitter/bd"
)

/*ChequeoBD es el middlew que me permite conocer el estado de la base de datos */
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConnection() == 0 {
			http.Error(w, "Conexión perdida con la Base de Datos", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
