package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"infctas/internal/core/entities"
	"infctas/internal/core/ports"
	"net/http"
	"strconv"
)

type ModuleHandler struct {
	ports.MouduleAPI
}

func NewModuleHandler(api ports.MouduleAPI) *ModuleHandler {
	return &ModuleHandler{
		MouduleAPI: api,
	}
}

func (h *ModuleHandler) CreateModuloHandler(w http.ResponseWriter, r *http.Request) {
	// Implementar la lógica para manejar la solicitud de crear un módulo
	var module entities.Module
	if err := json.NewDecoder(r.Body).Decode(&module); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	newModule, err := h.MouduleAPI.CreateModulo(r.Context(), module)
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Error creating module")
		return
	}
	// Devolver el nuevo módulo como respuesta
	respondJSON(w, http.StatusCreated, newModule)
}

func (h *ModuleHandler) GetModuloHandler(w http.ResponseWriter, r *http.Request) {
	// Implementar la lógica para manejar la solicitud de obtener un módulo
	// Extraer el ID del módulo de la solicitud
	vars := mux.Vars(r)
	idModulo, err := strconv.Atoi(vars["idModulo"])
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid module ID")
		return
	}
	module, err := h.MouduleAPI.GetModuloById(r.Context(), idModulo)
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Error retrieving module")
		return
	}
	// Devolver el módulo como respuesta
	respondJSON(w, http.StatusOK, module)
}

func (h *ModuleHandler) GetModulosHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	rol := vars["rol"]
	// Implementar la lógica para manejar la solicitud de obtener todos los módulos
	modules, err := h.MouduleAPI.GetModulos(r.Context(), rol)
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Error retrieving modules")
		return
	}

	// Devolver la lista de módulos como respuesta
	respondJSON(w, http.StatusOK, modules)
}

func (h *ModuleHandler) SetStatusModuloHandler(w http.ResponseWriter, r *http.Request) {
	// Implementar la lógica para manejar la solicitud de actualizar el estado de un módulo
	vars := mux.Vars(r)
	idModulo, err := strconv.Atoi(vars["idModulo"])
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
	// Extraer el ID del usuario que realiza la solicitud
	userID, err := strconv.Atoi(vars["user"])
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}
	// Llamar a la API para actualizar el estado del módulo
	module, err := h.MouduleAPI.SetStatusModulo(r.Context(), idModulo, statusInt, userID)
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
	idModulo, err := strconv.Atoi(vars["idModulo"])
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid module ID")
		return
	}
	coordinadorID, err := strconv.Atoi(vars["coordinador"])
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid coordinator ID")
		return
	}
	// Extraer el ID del usuario que realiza la solicitud
	userID, err := strconv.Atoi(vars["user"])
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}
	// Llamar a la API
	module, err := h.MouduleAPI.SetCoordinador(r.Context(), idModulo, coordinadorID, userID)
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Error updating coordinator")
		return
	}
	// Devolver el módulo actualizado como respuesta
	respondJSON(w, http.StatusOK, module)
}

func (h *ModuleHandler) SetResponsableHandler(w http.ResponseWriter, r *http.Request) {
	// Implementar la lógica para manejar la solicitud de actualizar los responsables de un módulo
	vars := mux.Vars(r)
	idModulo, err := strconv.Atoi(vars["idModulo"])
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid module ID")
		return
	}
	idCoordinador, err := strconv.Atoi(vars["coordinador"])
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid coordinator ID")
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
	// Extraer el ID del usuario que realiza la solicitud
	userID, err := strconv.Atoi(vars["user"])
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}
	// Llamar a la API
	module, err := h.MouduleAPI.SetResponsable(r.Context(), idModulo, idCoordinador, responsable1ID, responsable2ID, userID)
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Error updating responsible persons")
		return
	}
	// Devolver el módulo actualizado como respuesta
	respondJSON(w, http.StatusOK, module)
}

func (h *ModuleHandler) SetAreasHandler(w http.ResponseWriter, r *http.Request) {
	// Implementar la lógica para manejar la solicitud de actualizar las áreas de un módulo
	vars := mux.Vars(r)
	idModulo, err := strconv.Atoi(vars["idModulo"])
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid module ID")
		return
	}
	// Extraer el concatenado de ID's de las áreas de la solicitud
	areasIDStr := vars["areas"]

	// Extraer el ID del usuario que realiza la solicitud
	userID, err := strconv.Atoi(vars["user"])
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	// llamar a la API para actualizar las áreas del módulo
	module, err := h.MouduleAPI.SetAreas(r.Context(), idModulo, areasIDStr, userID)
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
	idModulo, err := strconv.Atoi(vars["idModulo"])
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid module ID")
		return
	}
	script := vars["script"]

	// Extraer el ID del usuario que realiza la solicitud
	userID, err := strconv.Atoi(vars["user"])
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}
	// Llamar a la API para actualizar el script del módulo
	module, err := h.MouduleAPI.SetScript(r.Context(), idModulo, script, userID)
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
	idModulo, err := strconv.Atoi(vars["idModulo"])
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid module ID")
		return
	}
	mail := vars["mail"]

	// Extraer el ID del usuario que realiza la solicitud
	userID, err := strconv.Atoi(vars["user"])
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	// Llamar a la API para actualizar el correo del módulo
	module, err := h.MouduleAPI.SetMail(r.Context(), idModulo, mail, userID)
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
	idModulo, err := strconv.Atoi(vars["idModulo"])
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid module ID")
		return
	}
	descripcion := vars["descripcion"]

	// Extraer el ID del usuario que realiza la solicitud
	userID, err := strconv.Atoi(vars["user"])
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	// Llamar a la API para actualizar la descripción del módulo
	module, err := h.MouduleAPI.SetDescripcion(r.Context(), idModulo, descripcion, userID)
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Error updating description")
		return
	}
	// Devolver el módulo actualizado como respuesta
	respondJSON(w, http.StatusOK, module)
}

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
