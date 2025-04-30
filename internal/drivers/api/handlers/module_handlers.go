package handlers

import (
	"encoding/json"
	"infctas/internal/core/entities"
	"infctas/internal/core/ports"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

type ModuleHandler struct {
	ports.MouduleAPI
}

func NewModuleHandler(api ports.MouduleAPI) *ModuleHandler {
	return &ModuleHandler{
		MouduleAPI: api,
	}
}

func (h *ModuleHandler) GetModuloHandler(w http.ResponseWriter, r *http.Request) {
	// Implementar la lógica para manejar la solicitud de obtener un módulo
	// Extraer el ID del módulo de la solicitud
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid module ID")
		return
	}
	module, err := h.MouduleAPI.GetModulo(r.Context(), id)
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Error retrieving module")
		return
	}
	// Devolver el módulo como respuesta
	respondJSON(w, http.StatusOK, module)
}

func (h *ModuleHandler) GetModulosHandler(w http.ResponseWriter, r *http.Request) {
	// Implementar la lógica para manejar la solicitud de obtener todos los módulos
	modules, err := h.MouduleAPI.GetModulos(r.Context())
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Error retrieving modules")
		return
	}
	// Devolver la lista de módulos como respuesta
	respondJSON(w, http.StatusOK, modules)
}

func (h *ModuleHandler) CreateModuloHandler(w http.ResponseWriter, r *http.Request) {
	// Implementar la lógica para manejar la solicitud de crear un módulo
	var module entities.Module
	if err := json.NewDecoder(r.Body).Decode(&module); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	newModule, err := h.MouduleAPI.CreateModulo(r.Context(), module.Name, module.Description, module.CoordinadorID, module.Areas, module.Mail, module.Script)
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Error creating module")
		return
	}
	// Devolver el nuevo módulo como respuesta
	respondJSON(w, http.StatusCreated, newModule)
}

func (h *ModuleHandler) SetStatusModuloHandler(w http.ResponseWriter, r *http.Request) {
	// Implementar la lógica para manejar la solicitud de actualizar el estado de un módulo
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid module ID")
		return
	}
	status := vars["status"]
	statusInt, err := strconv.Atoi(status)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid status value")
		return
	}
	module, err := h.MouduleAPI.SetStatusModulo(r.Context(), id, statusInt)
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Error updating module status")
		return
	}
	// Devolver el módulo actualizado como respuesta
	respondJSON(w, http.StatusOK, module)
}

func (h *ModuleHandler) SetCoordinadorHandler(w http.ResponseWriter, r *http.Request) {
	// Implementar la lógica para manejar la solicitud de actualizar el coordinador de un módulo
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid module ID")
		return
	}
	coordinadorID, err := strconv.Atoi(vars["coordinador"])
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid coordinator ID")
		return
	}
	module, err := h.MouduleAPI.SetCoordinador(r.Context(), id, coordinadorID)
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Error updating coordinator")
		return
	}
	// Devolver el módulo actualizado como respuesta
	respondJSON(w, http.StatusOK, module)
}

/*
func (h *ModuleHandler) SetResponsableHandler(w http.ResponseWriter, r *http.Request) {
	// Implementar la lógica para manejar la solicitud de actualizar los responsables de un módulo
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid module ID")
		return
	}
	responsable1ID, err := strconv.Atoi(vars["responsable1"])
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid first responsible ID")
		return
	}
	responsable2ID, err := strconv.Atoi(vars["responsable2"])
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid second responsible ID")
		return
	}
	module, err := h.MouduleAPI.SetResponsable(r.Context(), id, responsable1ID, responsable2ID)
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Error updating responsible persons")
		return
	}
	// Devolver el módulo actualizado como respuesta
	respondJSON(w, http.StatusOK, module)
}
*/

func (h *ModuleHandler) SetAreasHandler(w http.ResponseWriter, r *http.Request) {
	// Implementar la lógica para manejar la solicitud de actualizar las áreas de un módulo
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid module ID")
		return
	}
	areasIDStr := vars["areas"]
	areasID := strings.Split(areasIDStr, ",")
	var areas []int
	for _, areaStr := range areasID {
		areaID, err := strconv.Atoi(areaStr)
		if err != nil {
			respondError(w, http.StatusBadRequest, "Invalid area ID")
			return
		}
		areas = append(areas, areaID)
	}
	module, err := h.MouduleAPI.SetAreas(r.Context(), id, areas)
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Error updating areas")
		return
	}
	// Devolver el módulo actualizado como respuesta
	respondJSON(w, http.StatusOK, module)
}

func (h *ModuleHandler) SetScriptHandler(w http.ResponseWriter, r *http.Request) {
	// Implementar la lógica para manejar la solicitud de actualizar el script de un módulo
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid module ID")
		return
	}
	script := vars["script"]
	module, err := h.MouduleAPI.SetScript(r.Context(), id, script)
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Error updating script")
		return
	}
	// Devolver el módulo actualizado como respuesta
	respondJSON(w, http.StatusOK, module)
}

func (h *ModuleHandler) SetMailHandler(w http.ResponseWriter, r *http.Request) {
	// Implementar la lógica para manejar la solicitud de actualizar el correo de un módulo
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid module ID")
		return
	}
	mail := vars["mail"]
	module, err := h.MouduleAPI.SetMail(r.Context(), id, mail)
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Error updating mail")
		return
	}
	// Devolver el módulo actualizado como respuesta
	respondJSON(w, http.StatusOK, module)
}

func (h *ModuleHandler) SetDescripcionHandler(w http.ResponseWriter, r *http.Request) {
	// Implementar la lógica para manejar la solicitud de actualizar la descripción de un módulo
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid module ID")
		return
	}
	descripcion := vars["descripcion"]
	module, err := h.MouduleAPI.SetDescripcion(r.Context(), id, descripcion)
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Error updating description")
		return
	}
	// Devolver el módulo actualizado como respuesta
	respondJSON(w, http.StatusOK, module)
}

/*
func (h *ModuleHandler) SetLayOutHandler(w http.ResponseWriter, r *http.Request) {
	// Implementar la lógica para manejar la solicitud de actualizar el layout de un módulo
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid module ID")
		return
	}
	layoutID, err := strconv.Atoi(vars["layout"])
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid layout ID")
		return
	}
	module, err := h.MouduleAPI.SetLayOut(r.Context(), id, layoutID)
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Error updating layout")
		return
	}
	// Devolver el módulo actualizado como respuesta
	respondJSON(w, http.StatusOK, module)
}
*/

// Helper para respuestas de error
func respondError(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(map[string]string{"error": message})
	if err != nil {
		return
	}
}

// Helper para respuestas JSON
func respondJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		return
	}
}
