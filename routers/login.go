package routers

import (
	"encoding/json"
	"net/http"
	"time"
	"twittor-rodrigo/bd"
	"twittor-rodrigo/jwt"
	"twittor-rodrigo/models"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Usuario y/o contrasena invalidos "+err.Error(), http.StatusBadRequest)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El email del usuario es requerido ", http.StatusBadRequest)
		return
	}

	documento, existe := bd.Login(t.Email, t.Password)
	if existe == false {
		http.Error(w, "Usuario y/o contrasena invalidos ", http.StatusBadRequest)
		return
	}
	jwtKey, err := jwt.GeneroJWT(documento)
	if err != nil {
		http.Error(w, "Usuario y/o contrasena invalidos "+err.Error(), http.StatusBadRequest)
		return
	}

	resp := models.RespuestaLogin{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
}
