package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ingjeffer/twittor2/bd"
)

func LeoTweetsSeguidores(w http.ResponseWriter, r *http.Request) {
	paginaStr := r.URL.Query().Get("pagina")
	if len(paginaStr) < 1 {
		http.Error(w, "Debe enviar el parámetro página", http.StatusBadRequest)
		return
	}
	pagina, err := strconv.Atoi(paginaStr)
	if err != nil {
		http.Error(w, "Debe enviar el parámetro página como número", http.StatusBadRequest)
		return
	}

	respuesta, correcto := bd.LeoTweetsSeguidores(IDUsuario, pagina)
	if !correcto {
		http.Error(w, "Error al leer los tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&respuesta)
}
