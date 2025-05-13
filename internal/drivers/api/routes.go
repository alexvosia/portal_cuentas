package api

import (
	"github.com/gorilla/mux"
	"infctas/internal/drivers/api/handlers"
	"net/http"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.Use(CorsMiddleware)
	return router
}

// CorsMiddleware crea un middleware para manejar CORS
func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
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
	r.HandleFunc("getPosiblesResponsables/{coordinador}", handlers.GetPosiblesResponsablesHandler).Methods("GET")
	r.HandleFunc("getAreas", handlers.GetAreasHandler).Methods("GET")
}

func RegisterFileRoutes(r *mux.Router, handlers *handlers.RegistryFileHandler) {
	r.HandleFunc("/createFile", handlers.CreateRegistryFileHandler).Methods("POST")
	r.HandleFunc("/getFilesByModulo/{idModulo}", handlers.GetFileRegistryByModuloHandler).Methods("GET")
	r.HandleFunc("/deleteFile/{idFile}/{upper}", handlers.SetFinHandler).Methods("PUT")
}

func RegisterResponsableRoutes(r *mux.Router, handlers *handlers.ResponsableHandler) {
	r.HandleFunc("/getResponsablesConfigurados/{coordinador}", handlers.GetResponsablesConfiguradosHandler).Methods("GET")
}

func RegisterCuentaRoutes(r *mux.Router, handlers *handlers.CuentaHandler) {
	r.HandleFunc("/getInfoCuenta/{cuenta}/{idArea}", handlers.GetInfoCuentaHandler).Methods("GET")
}
