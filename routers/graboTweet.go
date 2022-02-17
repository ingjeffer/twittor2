package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/ingjeffer/twittor2/bd"
	"github.com/ingjeffer/twittor2/models"
)

func GraboTweet(w http.ResponseWriter, r *http.Request) {
	var mensaje models.Tweet

	err := json.NewDecoder(r.Body).Decode(&mensaje)

	registro := models.GraboTweet{
		UserId:  IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}

	_, status, err := bd.InsertoTweet(registro)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar insertar el registro "+err.Error(), http.StatusBadRequest)
		return
	}

	if !status {
		http.Error(w, "No se ha podido insertarr el registro", http.StatusBadRequest)
		return
	}

	// json.NewEncoder(w).
	w.WriteHeader(http.StatusCreated)

}
