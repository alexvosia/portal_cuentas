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
	r.HandleFunc("/modulo", handlers.CreateModuloHandler).Methods("POST")
	r.HandleFunc("/modulo/{idModulo}", handlers.GetModuloHandler).Methods("GET")
	r.HandleFunc("/modulos/{rol}", handlers.GetModulosHandler).Methods("GET")
	r.HandleFunc("/statusModulo/{idModulo}/{status}/{user}", handlers.SetStatusModuloHandler).Methods("PUT")
	r.HandleFunc("/coordinador/{idModulo}/{coordinador}/{user}", handlers.SetCoordinadorHandler).Methods("PUT")
	r.HandleFunc("/responsable/{idModulo}/{coordinador}/{responsable1}/{responsable2}/{user}", handlers.SetResponsableHandler).Methods("PUT")
	r.HandleFunc("/areas/{idModulo}/{areas}/{user}", handlers.SetAreasHandler).Methods("PUT")
	r.HandleFunc("/script/{idModulo}/{script}/{user}", handlers.SetScriptHandler).Methods("PUT")
	r.HandleFunc("/mail/{idModulo}/{mail}/{user}", handlers.SetMailHandler).Methods("PUT")
	r.HandleFunc("/descripcion/{idModulo}/{descripcion}/{user}", handlers.SetDescripcionHandler).Methods("PUT")
}

func RegisterChecaDataRoutes(r *mux.Router, handlers *handlers.ChecaDataHandler) {
	r.HandleFunc("/coordinadores", handlers.GetCoordinadoresHandler).Methods("GET")
	r.HandleFunc("/posiblesResponsables/{coordinador}", handlers.GetPosiblesResponsablesHandler).Methods("GET")
	r.HandleFunc("/areas", handlers.GetAreasHandler).Methods("GET")
}

func RegisterFileRoutes(r *mux.Router, handlers *handlers.RegistryFileHandler) {
	r.HandleFunc("/file", handlers.CreateRegistryFileHandler).Methods("POST")
	r.HandleFunc("/filesByModulo/{idModulo}", handlers.GetFileRegistryByModuloHandler).Methods("GET")
	r.HandleFunc("/file/{idFile}/{upper}", handlers.SetFinHandler).Methods("DELETE")
}

func RegisterResponsableRoutes(r *mux.Router, handlers *handlers.ResponsableHandler) {
	r.HandleFunc("/responsablesConfigurados/{coordinador}", handlers.GetResponsablesConfiguradosHandler).Methods("GET")
}

func RegisterCuentaRoutes(r *mux.Router, handlers *handlers.CuentaHandler) {
	r.HandleFunc("/infoCuenta/{cuenta}/{idArea}", handlers.GetInfoCuentaHandler).Methods("GET")
}
