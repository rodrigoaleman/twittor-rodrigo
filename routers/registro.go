package routers

import (
	"encoding/json"
	"net/http"
	"twittor-rodrigo/bd"
	"twittor-rodrigo/models"
)

// Registro es la funcion para crear en la BD el registro de usuarios
func Registro(w http.ResponseWriter, r *http.Request) {

	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El email de usuario es requerido", 400)
		return
	}

	if len(t.Password) < 6 {
		http.Error(w, "Debe especificar una contrasenna de al menos 6 caracteres", 400)
		return
	}

	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)
	if encontrado == true {
		http.Error(w, "Ya existe un usuario registrado con ese email", http.StatusBadRequest)
		return
	}

	_, status, err := bd.InsertoRegistro(t)
	if err != nil {
		http.Error(w, "Ocurrio un erro al intentar registrar el usuario "+err.Error(), http.StatusBadRequest)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado insertar el usuario"+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
