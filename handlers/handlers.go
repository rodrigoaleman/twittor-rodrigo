package handlers

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
	"twittor-rodrigo/middleware"
	"twittor-rodrigo/routers"
)

// Manejadores seteo mi puerto, el Handler y pongo a escuchar al Servidor
func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middleware.ChequeoBD(routers.Registro)).Methods("POST")
	router.HandleFunc("/login", middleware.ChequeoBD(routers.Login)).Methods("POST")
	router.HandleFunc("/verperfil", middleware.ChequeoBD(middleware.ValidoJWT(routers.VerPerfil))).Methods("GET")
	router.HandleFunc("/modificarperfil", middleware.ChequeoBD(middleware.ValidoJWT(routers.ModificarPerfil))).Methods("PUT")
	router.HandleFunc("/tweet", middleware.ChequeoBD(middleware.ValidoJWT(routers.GraboTweet))).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
