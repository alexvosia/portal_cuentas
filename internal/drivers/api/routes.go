package api

import (
	"github.com/gorilla/mux"
	"infctas/internal/drivers/api/handlers"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	return router
}

func RegisterModuleRoutes(r *mux.Router, handlers *handlers.ModuleHandler) {
	r.HandleFunc("/createModulo", handlers.CreateModuloHandler).Methods("POST")
	r.HandleFunc("/getModulo/{id}", handlers.GetModuloHandler).Methods("GET")
	r.HandleFunc("/getModulos/{rol}", handlers.GetModulosHandler).Methods("GET")
	r.HandleFunc("/setStatusModulo/{id}/{status}", handlers.SetStatusModuloHandler).Methods("PUT")
	r.HandleFunc("/setCoordinador/{id}/{coordinador}", handlers.SetCoordinadorHandler).Methods("PUT")
	r.HandleFunc("/setResponsable/{coordinador}/{responsable1}/{responsable2}", handlers.SetResponsableHandler).Methods("PUT")
	r.HandleFunc("/setAreas/{id}/{areas}", handlers.SetAreasHandler).Methods("PUT")
	r.HandleFunc("/setScript/{id}/{script}", handlers.SetScriptHandler).Methods("PUT")
	r.HandleFunc("/setMail/{id}/{mail}", handlers.SetMailHandler).Methods("PUT")
	r.HandleFunc("/setDescripcion/{id}/{descripcion}", handlers.SetDescripcionHandler).Methods("PUT")
}

func RegisterChecaDataRoutes(r *mux.Router, handlers *handlers.ChecaDataHandler) {
	r.HandleFunc("getCoordinadores", handlers.GetCoordinadoresHandler).Methods("GET")
	r.HandleFunc("getResponsables/{coordinador}", handlers.GetResponsablesHandler).Methods("GET")
	r.HandleFunc("getAreas", handlers.GetAreasHandler).Methods("GET")
}

func RegisterFileRoutes(r *mux.Router, handlers *handlers.FileHandler) {
	r.HandleFunc("/createFile", handlers.CreateFileHandler).Methods("POST")
	r.HandleFunc("/getFilesByModulo/{idModulo}", handlers.GetFilesByModuloHandler).Methods("GET")
	r.HandleFunc("/deleteFile/{id}", handlers.SetFinHandler).Methods("PUT")
}

func RegisterCuentaRoutes(r *mux.Router, handlers *handlers.CuentaHandler) {
	r.HandleFunc("/getInfoCuenta/{cuenta}/{area}", handlers.GetCuentasHandler).Methods("GET")
}
