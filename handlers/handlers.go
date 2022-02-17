package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/ingjeffer/twittor2/middlew"
	"github.com/ingjeffer/twittor2/routers"
	"github.com/rs/cors"
)

/* Manejadores seteo mi puerto, el hanlder y levanta el servidor */
func Manejadores() {
	router := mux.NewRouter()

	// router.HandleFunc("/registro", routers.Registro).Methods("POST")
	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")
	router.HandleFunc("/login", middlew.ChequeoBD(routers.Login)).Methods("POST")

	router.HandleFunc("/verperfil", middlew.ChequeoBD(middlew.ValidoJWT(routers.VerPerfil))).Methods("GET")
	router.HandleFunc("/modificar-perfil", middlew.ChequeoBD(middlew.ValidoJWT(routers.ModificoRegistro))).Methods("PUT")

	router.HandleFunc("/tweet", middlew.ChequeoBD(middlew.ValidoJWT(routers.GraboTweet))).Methods("POST")
	router.HandleFunc("/tweet", middlew.ChequeoBD(middlew.ValidoJWT(routers.LeoTweets))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8085"
	}
	handler := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
