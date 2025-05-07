package handlers

import (
	"github.com/gorilla/mux"
	"infctas/internal/core/ports"
	"net/http"
	"strconv"
)

type CuentaHandler struct {
	ports.CuentaAPI
}

func NewCuentaHandler(api ports.CuentaAPI) *CuentaHandler {
	return &CuentaHandler{
		CuentaAPI: api,
	}
}

func (h *CuentaHandler) GetInfoCuentaHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if vars["cuenta"] == "" || vars["idArea"] == "" {
		respondError(w, http.StatusBadRequest, "Missing cuenta or idArea parameter")
		return
	}

	cuenta, err := strconv.Atoi(vars["cuenta"])
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid cuenta parameter")
		return
	}
	idArea, err := strconv.Atoi(vars["idArea"])
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid idArea parameter")
		return
	}

	info, err := h.CuentaAPI.GetInfoCuenta(r.Context(), cuenta, idArea)
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Error retrieving account information")
		return
	}

	respondJSON(w, http.StatusOK, info)
}
