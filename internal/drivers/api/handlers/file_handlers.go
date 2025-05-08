package handlers

import (
	"fmt"
	"github.com/gorilla/mux"
	"infctas/internal/core/entities"
	"infctas/internal/core/ports"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
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
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}
	// Procesar el formulario multipart
	err := r.ParseMultipartForm(32 << 20) // Límite de memoria a 32 MB
	if err != nil {
		http.Error(w, "Error al parsear el formulario", http.StatusBadRequest)
		fmt.Println("Error al parsear el formulario:", err)
		return
	}
	// Obtener el archivo CSV del formulario
	file, header, err := r.FormFile("csv_file") // "csv_file" debe ser el nombre del campo del archivo
	if err != nil {
		http.Error(w, "Error al obtener el archivo", http.StatusBadRequest)
		fmt.Println("Error al obtener el archivo:", err)
		return
	}
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			http.Error(w, "Error al cerrar el archivo", http.StatusInternalServerError)
			fmt.Println("Error al cerrar el archivo:", err)
		}
	}(file)

	// Obtener los otros valores de configuración del formulario
	var fileRegistry entities.RegistryFileCSV
	idModulo, err := strconv.Atoi(r.FormValue("id_modulo"))
	if err != nil {
		http.Error(w, "Error al convertir id_modulo a entero", http.StatusBadRequest)
		fmt.Println("Error al convertir id_modulo a entero:", err)
		return
	}
	fileRegistry.IdModulo = idModulo
	fileRegistry.Nombre = r.FormValue("id_modulo") + `_` + r.FormValue("nombre")
	fileRegistry.TipoRegistro = r.FormValue("tipo_registro")
	idCreador, err := strconv.Atoi(r.FormValue("id_creador"))
	if err != nil {
		http.Error(w, "Error al convertir id_creador a entero", http.StatusBadRequest)
		fmt.Println("Error al convertir id_creador a entero:", err)
		return
	}
	fileRegistry.Creador = idCreador

	// Guardar el archivo en el sistema de archivos
	uploadPath := fmt.Sprintf("./files/%s", header.Filename)
	err = os.MkdirAll(uploadPath, os.ModeDir|0755) // Crear el directorio si no existe
	if err != nil {
		http.Error(w, "Error al crear el directorio de destino", http.StatusInternalServerError)
		fmt.Println("Error al crear el directorio:", err)
		return
	}

	filename := header.Filename
	filePath := filepath.Join(uploadPath, filename)

	// Crear el archivo de destino en el servidor
	dst, err := os.Create(filePath)
	if err != nil {
		http.Error(w, "Error al crear el archivo de destino", http.StatusInternalServerError)
		fmt.Println("Error al crear el archivo de destino:", err)
		return
	}
	defer func(dst *os.File) {
		err := dst.Close()
		if err != nil {
			http.Error(w, "Error al cerrar el archivo de destino", http.StatusInternalServerError)
			log.Println("Error al cerrar el archivo de destino:", err)
		}
	}(dst)

	// Leer el archivo subido en chunks y escribirlo en el archivo de destino
	buffer := make([]byte, 4096) // Tamaño del buffer (ajusta según necesidad)
	totalBytes := 0
	for {
		bytesRead, err := file.Read(buffer)
		if err != nil {
			if err.Error() == "EOF" {
				break // Fin del archivo
			}
			http.Error(w, "Error al leer el archivo", http.StatusInternalServerError)
			log.Println("Error al leer el archivo:", err)
			return
		}
		totalBytes += bytesRead
		if _, err := dst.Write(buffer[:bytesRead]); err != nil {
			http.Error(w, "Error al escribir en el archivo de destino", http.StatusInternalServerError)
			log.Println("Error al escribir en el archivo de destino:", err)
			return
		}
	}

	err = h.RegistryFileAPI.CreateRegistryFile(r.Context(), fileRegistry.Nombre, fileRegistry.IdModulo, fileRegistry.TipoRegistro, fileRegistry.Creador)
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
	files, err := h.RegistryFileAPI.GetRegistryFilesByModulo(r.Context(), idModulo)
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
