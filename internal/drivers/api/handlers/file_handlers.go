package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"infctas/internal/core/entities"
	"infctas/internal/core/ports"
	"net/http"
	"strconv"
)

type FileHandler struct {
	ports.FileAPI
}

func NewFileHandler(api ports.FileAPI) *FileHandler {
	return &FileHandler{
		FileAPI: api,
	}
}

func (h *FileHandler) CreateFileHandler(w http.ResponseWriter, r *http.Request) {
	var file entities.FileCSV
	if err := json.NewDecoder(r.Body).Decode(&file); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	err := h.FileAPI.CreateFile(r.Context(), file.Nombre, file.IdModulo, file.TipoRegistro, file.Creador)
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Error creating file")
		return
	}
}

func (h *FileHandler) GetFilesByModuloHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idModulo, err := strconv.Atoi(vars["idModulo"])
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid module ID")
		return
	}
	files, err := h.FileAPI.GetFilesByModulo(r.Context(), idModulo)
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Error retrieving files")
		return
	}
	respondJSON(w, http.StatusOK, files)
}

func (h *FileHandler) SetFinHandler(w http.ResponseWriter, r *http.Request) {
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
	file, err := h.FileAPI.SetFin(r.Context(), idFile, upper)
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Error setting file status")
		return
	}
	respondJSON(w, http.StatusOK, file)
}
