package handlers

import (
	"github.com/gorilla/mux"
	"infctas/internal/core/ports"
	"net/http"
	"strconv"
)

type ResponsableHandler struct {
	ports.ResponsableAPI
}

func NewResponsableHandler(api ports.ResponsableAPI) *ResponsableHandler {
	return &ResponsableHandler{
		ResponsableAPI: api,
	}
}

func (h *ResponsableHandler) GetResponsablesByCoordinadorHandler(w http.ResponseWriter, r *http.Request) {
	// Implementar la lógica para manejar la solicitud de obtener un responsable por coordinador
	// Extraer el ID del coordinador de la solicitud
	vars := mux.Vars(r)
	idCoordinador, err := strconv.Atoi(vars["coordinador"])
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid coordinator ID")
		return
	}
	responsable, err := h.ResponsableAPI.FindResponsableByCoordinador(r.Context(), idCoordinador)
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Error retrieving responsible person")
		return
	}
	// Devolver el responsable como respuesta
	respondJSON(w, http.StatusOK, responsable)
}