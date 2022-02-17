package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ingjeffer/twittor2/bd"
)

func LeoTweets(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar un id ", http.StatusBadRequest)
		return
	}

	paginaStr := r.URL.Query().Get("pagina")
	if len(paginaStr) < 1 {
		http.Error(w, "Debe enviar un pagina ", http.StatusBadRequest)
		return
	}

	pagina, err := strconv.Atoi(paginaStr)
	if err != nil {
		http.Error(w, "Debe enviar un pagina  con valor mayor a 0", http.StatusBadRequest)
		return
	}

	pag := int64(pagina)
	respuesta, correcto := bd.LeoTweets(ID, pag)
	if correcto == false {
		http.Error(w, "Error al leer Tweets ", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)
}
