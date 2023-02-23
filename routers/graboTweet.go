package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"twittor-rodrigo/bd"
	"twittor-rodrigo/models"
)

func GraboTweet(w http.ResponseWriter, r *http.Request) {

	var mensaje models.Tweet
	err := json.NewDecoder(r.Body).Decode(&mensaje)
	if err != nil {
		return
	}

	registro := models.GraboTweet{
		UserID:  IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}

	_, status, err := bd.InsertoTweet(registro)
	if err != nil {
		http.Error(w, "Error al insertar tweet "+err.Error(), http.StatusBadRequest)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado insertar el tweet", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
