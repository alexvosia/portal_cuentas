package main

import (
	"database/sql"
	"infctas/internal/core/usecases"
	"infctas/internal/drivers/api"
	"infctas/internal/drivers/api/handlers"
	"infctas/internal/providers/pic_repo"
	"log"
	"net/http"

	_ "github.com/denisenkom/go-mssqldb"
)

func main() {
	// Configurar conexión a la base de datos
	db, err := sql.Open("sqlserver", "sqlserver://pic_owner:pic_1234@localhost:1433?database=pic")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)

	// Crear el router
	router := api.NewRouter()

	// Configurar Provider MSSQL
	repo := pic_repo.NewMSSQLModuleRepo(db)
	// Crear el servicio de repositorio de módulos
	moduleservice := usecases.NewModuleService(repo)
	// Crear el handler de Módulos
	moduleHandler := handlers.NewModuleHandler(moduleservice)
	// Configurar las rutas de modulos en la API
	api.RegisterModuleRoutes(router, moduleHandler)

	// Configurar el Provider ORCLChecaData
	checaDataRepo := pic_repo.NewORCLChecaData(db)
	// Crear el servicio de repositorio de ChecaData
	checaDataService := usecases.NewChecaDataService(checaDataRepo)
	// Crear el handler de ChecaData
	checaDataHandler := handlers.NewChecaDataHandler(checaDataService)
	// Configurar las rutas de ChecaData en la API
	api.RegisterChecaDataRoutes(router, checaDataHandler)

	// Crear el handler de Archivos

	// Iniciar el servidor
	log.Println("Iniciando servidor en :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
