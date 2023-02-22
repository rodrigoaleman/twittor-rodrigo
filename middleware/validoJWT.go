package middleware

import (
	"net/http"
	"twittor-rodrigo/routers"
)

// ValidoJWT permite validar el JWT que nos viene en la peticion
func ValidoJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := routers.ProcesoToken(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "Error en el token !", http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	}
}
