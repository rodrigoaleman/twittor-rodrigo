package middleware

import (
	"net/http"
	"twittor-rodrigo/bd"
)

// ChequeoBD es el middleware que me perminte conocer el estado de la BD
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConnection() == false {
			http.Error(w, "Conexion perdida con la BD", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
