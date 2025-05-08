package main

import (
	"database/sql"
	"infctas/internal/core/usecases"
	"infctas/internal/drivers/api"
	"infctas/internal/drivers/api/handlers"
	"infctas/internal/providers/checa_data"
	"infctas/internal/providers/duck_repo"
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
	mssqlRepo := pic_repo.NewMSSQLModuleRepo(db)

	// Crear el servicio de repositorio de módulos
	moduleservice := usecases.NewModuleService(mssqlRepo)
	// Crear el handler de Módulos
	moduleHandler := handlers.NewModuleHandler(moduleservice)
	// Configurar las rutas de modulos en la API
	api.RegisterModuleRoutes(router, moduleHandler)

	// Crear el servicio de repositorio de responsables
	responsableService := usecases.NewResponsableService(mssqlRepo)
	// Crear el handler de responsables
	responsableHandler := handlers.NewResponsableHandler(responsableService)
	// Configurar las rutas de responsables en la API
	api.RegisterResponsableRoutes(router, responsableHandler)

	// Crear el servicio de repositorio de files
	fileService := usecases.NewFileService(mssqlRepo)
	// Crear el handler de files
	fileHandler := handlers.NewFileHandler(fileService)
	// Configurar las rutas de files en la API
	api.RegisterFileRoutes(router, fileHandler)

	// Configurar el Provider ORCLChecaData
	checaDataRepo := checa_data.NewORCLChecaData(db)
	// Crear el servicio de repositorio de ChecaData
	checaDataService := usecases.NewChecaDataService(checaDataRepo)
	// Crear el handler de ChecaData
	checaDataHandler := handlers.NewChecaDataHandler(checaDataService)
	// Configurar las rutas de ChecaData en la API
	api.RegisterChecaDataRoutes(router, checaDataHandler)

	// Configurar el Provider DUCKFileRepo
	rowsRepo := duck_repo.NewDUCKFileRepo(db)
	// Crear el servicio de repositorio de Archivos
	rowsService := usecases.NewFileService(rowsRepo)
	// Crear el handler de Archivos
	rowsHandler := handlers.NewFileHandler(rowsService)
	// Configurar las rutas de Archivos en la API
	api.RegisterFileRoutes(router, rowsHandler)

	// Iniciar el servidor
	log.Println("Iniciando servidor en :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
