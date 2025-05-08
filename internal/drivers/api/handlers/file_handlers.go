package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"infctas/internal/core/entities"
	"infctas/internal/core/ports"
	"net/http"
	"strconv"
)

type RegistryFileHandler struct {
	ports.RegistryFileAPI
}

func NewFileHandler(api ports.RegistryFileAPI) *RegistryFileHandler {
	return &RegistryFileHandler{
		RegistryFileAPI: api,
	}
}

func (h *RegistryFileHandler) CreateRegistryFileHandler(w http.ResponseWriter, r *http.Request) {
	var file entities.RegistryFileCSV
	if err := json.NewDecoder(r.Body).Decode(&file); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	err := h.RegistryFileAPI.CreateRegistryFile(r.Context(), file.Nombre, file.IdModulo, file.TipoRegistro, file.Creador)
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Error creating file")
		return
	}
}

func (h *RegistryFileHandler) GetFileRegistryByModuloHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idModulo, err := strconv.Atoi(vars["idModulo"])
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid module ID")
		return
	}
	files, err := h.RegistryFileAPI.FindFilesByModule(r.Context(), idModulo)
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Error retrieving files")
		return
	}
	respondJSON(w, http.StatusOK, files)
}

func (h *RegistryFileHandler) SetFinHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idFile, err := strconv.Atoi(vars["idFile"])
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid file ID")
		return
	}
	upper, err := strconv.Atoi(vars["upper"])
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid upper ID")
		return
	}
	file, err := h.RegistryFileAPI.SetFinRegistryFile(r.Context(), idFile, upper)
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Error setting file status")
		return
	}
	respondJSON(w, http.StatusOK, file)
}
