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
	r.HandleFunc("/getModulo/{id}", handlers.GetModuloHandler).Methods("GET")
	r.HandleFunc("/getModulos", handlers.GetModulosHandler).Methods("GET")
	r.HandleFunc("/createModulo", handlers.CreateModuloHandler).Methods("POST")
	r.HandleFunc("/setStatusModulo/{id}/{status}", handlers.SetStatusModuloHandler).Methods("PUT")
	r.HandleFunc("/setCoordinador/{id}/{coordinador}", handlers.SetCoordinadorHandler).Methods("PUT")
	r.HandleFunc("/setAreas/{id}/{areas}", handlers.SetAreasHandler).Methods("PUT")
	r.HandleFunc("/setScript/{id}/{script}", handlers.SetScriptHandler).Methods("PUT")
	r.HandleFunc("/setMail/{id}/{mail}", handlers.SetMailHandler).Methods("PUT")
	r.HandleFunc("/setDescripcion/{id}/{descripcion}", handlers.SetDescripcionHandler).Methods("PUT")

	/*
		r.HandleFunc("/setLayOut/{id}/{layout}", handlers.SetLayOutHandler).Methods("PUT")
		r.HandleFunc("/setResponsable/{id}/{responsable1}/{responsable2}", handlers.SetResponsableHandler).Methods("PUT")
	*/

}
