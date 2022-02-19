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
	router.HandleFunc("/tweet", middlew.ChequeoBD(middlew.ValidoJWT(routers.EliminarTweet))).Methods("DELETE")

	router.HandleFunc("/subir-avatar", middlew.ChequeoBD(middlew.ValidoJWT(routers.SubirAvatar))).Methods("POST")
	router.HandleFunc("/subir-avatar", middlew.ChequeoBD(middlew.ValidoJWT(routers.ObtenerAvatar))).Methods("GET")

	router.HandleFunc("/subir-banner", middlew.ChequeoBD(middlew.ValidoJWT(routers.SubirBanner))).Methods("POST")
	router.HandleFunc("/subir-banner", middlew.ChequeoBD(middlew.ValidoJWT(routers.ObtenerBanner))).Methods("GET")

	router.HandleFunc("/relacion", middlew.ChequeoBD(middlew.ValidoJWT(routers.AltaRelacion))).Methods("POST")
	router.HandleFunc("/relacion", middlew.ChequeoBD(middlew.ValidoJWT(routers.BajaRelacion))).Methods("DELETE")
	router.HandleFunc("/relacion", middlew.ChequeoBD(middlew.ValidoJWT(routers.ConsultaRelacion))).Methods("GET")

	router.HandleFunc("/lista-usuarios", middlew.ChequeoBD(middlew.ValidoJWT(routers.ListaUsuarios))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8085"
	}
	handler := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
