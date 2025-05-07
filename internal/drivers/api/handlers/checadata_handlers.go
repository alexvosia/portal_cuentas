package handlers

import (
	"github.com/gorilla/mux"
	"infctas/internal/core/ports"
	"net/http"
	"strconv"
)

type ChecaDataHandler struct {
	ports.ChecaDataAPI
}

func NewChecaDataHandler(api ports.ChecaDataAPI) *ChecaDataHandler {
	return &ChecaDataHandler{
		ChecaDataAPI: api,
	}
}

func (h *ChecaDataHandler) GetCoordinadoresHandler(w http.ResponseWriter, r *http.Request) {
	// Implementar la lógica para manejar la solicitud de obtener todos los coordinadores
	coordinadores, err := h.ChecaDataAPI.GetCoordinadores(r.Context())
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Error retrieving coordinators")
		return
	}

	// Devolver la lista de coordinadores como respuesta
	respondJSON(w, http.StatusOK, coordinadores)
}

func (h *ChecaDataHandler) GetResponsablesHandler(w http.ResponseWriter, r *http.Request) {
	// Implementar la lógica para manejar la solicitud de obtener todos los responsables
	vars := mux.Vars(r)
	coordinador, err := strconv.Atoi(vars["coordinador"])
	responsables, err := h.ChecaDataAPI.GetResponsables(r.Context(), coordinador)
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Error retrieving responsibles")
		return
	}

	// Devolver la lista de responsables como respuesta
	respondJSON(w, http.StatusOK, responsables)
}

func (h *ChecaDataHandler) GetAreasHandler(w http.ResponseWriter, r *http.Request) {
	// Implementar la lógica para manejar la solicitud de obtener todas las áreas
	areas, err := h.ChecaDataAPI.GetAreas(r.Context())
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Error retrieving areas")
		return
	}

	// Devolver la lista de áreas como respuesta
	respondJSON(w, http.StatusOK, areas)
}
