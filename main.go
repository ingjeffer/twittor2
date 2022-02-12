package main

import (
	"log"

	"github.com/ingjeffer/twittor2/bd"
	"github.com/ingjeffer/twittor2/handlers"
)

func main() {
	if bd.ChequeoConexion() == 0 {
		log.Fatal("Sin Conexión a la BD")
		return
	}
	handlers.Manejadores()
}
