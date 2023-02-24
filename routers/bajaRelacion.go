package routers

import (
	"net/http"
	"twittor-rodrigo/bd"
	"twittor-rodrigo/models"
)

func BajaRelacion(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro id", http.StatusBadRequest)
		return
	}

	var t models.Relacion

	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	status, err := bd.BorroRelacion(t)
	if err != nil {
		http.Error(w, "Ocurrio un error al borrar la relacion "+err.Error(), http.StatusBadRequest)
		return
	}

	if status == false {
		http.Error(w, "No se a logrado borrar la relacion "+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
