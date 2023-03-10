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
	router.HandleFunc("/tweet", middleware.ChequeoBD(middleware.ValidoJWT(routers.LeoTweets))).Methods("GET")
	router.HandleFunc("/eliminartweet", middleware.ChequeoBD(middleware.ValidoJWT(routers.EliminarTweet))).Methods("DELETE")

	router.HandleFunc("/subiravatar", middleware.ChequeoBD(middleware.ValidoJWT(routers.SubirAvatar))).Methods("POST")
	router.HandleFunc("/obteneravatar", middleware.ChequeoBD(middleware.ValidoJWT(routers.ObtenerAvatar))).Methods("GET")
	router.HandleFunc("/subirbanner", middleware.ChequeoBD(middleware.ValidoJWT(routers.SubirBanner))).Methods("POST")
	router.HandleFunc("/obtenerbanner", middleware.ChequeoBD(middleware.ValidoJWT(routers.ObtenerBanner))).Methods("GET")

	router.HandleFunc("/altarelacion", middleware.ChequeoBD(middleware.ValidoJWT(routers.AltaRelacion))).Methods("POST")
	router.HandleFunc("/bajarelacion", middleware.ChequeoBD(middleware.ValidoJWT(routers.BajaRelacion))).Methods("DELETE")
	router.HandleFunc("/consultarelacion", middleware.ChequeoBD(middleware.ValidoJWT(routers.ConsultaRelacion))).Methods("GET")

	router.HandleFunc("/listausuarios", middleware.ChequeoBD(middleware.ValidoJWT(routers.ListaUsuarios))).Methods("GET")
	router.HandleFunc("/leotweetsseguidores", middleware.ChequeoBD(middleware.ValidoJWT(routers.LeoTweetSeguidores))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
