package routers

import (
	"encoding/json"
	"net/http"

	"github.com/ingjeffer/twittor2/bd"
	"github.com/ingjeffer/twittor2/models"
)

func ModificoRegistro(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Datos incorrectos "+err.Error(), http.StatusBadRequest)
		return
	}

	var status bool
	status, err = bd.ModificoRegistro(t, IDUsuario)
	if err != nil {
		http.Error(w, "Ocurrió un erro al intentar modificar el registro "+err.Error(), http.StatusBadRequest)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado modificar el registro ", http.StatusBadRequest)
		return
	}

	// var resObj models.ResponseGenerico
	// resObj = {
	// 	Ok:      true,
	// 	Message: "El registro se modificó exitosamente",
	// }

	response, _ := json.Marshal(models.ResponseGenerico{
		Ok:      true,
		Message: "El registro se modificó exitosamente",
	})

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
