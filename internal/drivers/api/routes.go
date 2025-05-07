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
	r.HandleFunc("/getModulo/{idModulo}", handlers.GetModuloHandler).Methods("GET")
	r.HandleFunc("/getModulos/{rol}", handlers.GetModulosHandler).Methods("GET")
	r.HandleFunc("/setStatusModulo/{idModulo}/{status}/{user}", handlers.SetStatusModuloHandler).Methods("PUT")
	r.HandleFunc("/setCoordinador/{idModulo}/{coordinador}/{user}", handlers.SetCoordinadorHandler).Methods("PUT")
	r.HandleFunc("/setResponsable/{idModulo}/{coordinador}/{responsable1}/{responsable2}/{user}", handlers.SetResponsableHandler).Methods("PUT")
	r.HandleFunc("/setAreas/{idModulo}/{areas}/{user}", handlers.SetAreasHandler).Methods("PUT")
	r.HandleFunc("/setScript/{idModulo}/{script}/{user}", handlers.SetScriptHandler).Methods("PUT")
	r.HandleFunc("/setMail/{idModulo}/{mail}/{user}", handlers.SetMailHandler).Methods("PUT")
	r.HandleFunc("/setDescripcion/{idModulo}/{descripcion}/{user}", handlers.SetDescripcionHandler).Methods("PUT")
}

func RegisterChecaDataRoutes(r *mux.Router, handlers *handlers.ChecaDataHandler) {
	r.HandleFunc("getCoordinadores", handlers.GetCoordinadoresHandler).Methods("GET")
	r.HandleFunc("getPosiblesResponsables/{coordinador}", handlers.GetResponsablesHandler).Methods("GET")
	r.HandleFunc("getAreas", handlers.GetAreasHandler).Methods("GET")
}

func RegisterFileRoutes(r *mux.Router, handlers *handlers.FileHandler) {
	r.HandleFunc("/createFile", handlers.CreateFileHandler).Methods("POST")
	r.HandleFunc("/getFilesByModulo/{idModulo}", handlers.GetFilesByModuloHandler).Methods("GET")
	r.HandleFunc("/deleteFile/{idFile}/{upper}", handlers.SetFinHandler).Methods("PUT")
}

func ResisterResponsableRoutes(r *mux.Router, handlers *handlers.ResponsableHandler) {
	r.HandleFunc("/getResponsablesConfigurados/{coordinador}", handlers.GetResponsablesByCoordinadorHandler).Methods("GET")
}

func RegisterCuentaRoutes(r *mux.Router, handlers *handlers.CuentaHandler) {
	r.HandleFunc("/getInfoCuenta/{cuenta}/{idArea}", handlers.GetInfoCuentaHandler).Methods("GET")
}
