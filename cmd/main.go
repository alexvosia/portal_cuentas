package main

import (
	"infctas/internal/config"
	"infctas/internal/core/domain"
	"infctas/internal/drivers/api"
	"infctas/internal/drivers/api/handlers"
	"infctas/internal/providers/checa_data"
	"infctas/internal/providers/duck_repo"
	"infctas/internal/providers/pic_repo"
	"log"
	"net/http"
)

func main() {
	// Inicializar la base de datos
	dbManager := config.GetDBManager()
	dbCheca, err := dbManager.GetConn("checa")
	if err != nil {
		log.Fatalf("Error al conectar a la base de datos Checa: %v", err)
	}
	dbPic, err := dbManager.GetConn("pic")
	if err != nil {
		log.Fatalf("Error al conectar a la base de datos PIC: %v", err)
	}
	dbDuck, err := dbManager.GetConn("duck")
	if err != nil {
		log.Fatalf("Error al conectar a la base de datos Duck: %v", err)
	}

	defer func() {
		if err := dbManager.CloseAll(); err != nil {
			log.Fatalf("Error al cerrar las conexiones: %v", err)
		}
	}()

	// Crear el router
	router := api.NewRouter()

	// Configurar Provider MSSQL
	mssqlRepo := pic_repo.NewMSSQLModuleRepo(dbPic)

	// Crear el servicio de repositorio de módulos
	moduleservice := domain.NewModuleService(mssqlRepo)
	// Crear el handler de Módulos
	moduleHandler := handlers.NewModuleHandler(moduleservice)
	// Configurar las rutas de modulos en la API
	api.RegisterModuleRoutes(router, moduleHandler)

	// Crear el servicio de repositorio de responsables
	responsableService := domain.NewResponsableService(mssqlRepo)
	// Crear el handler de responsables
	responsableHandler := handlers.NewResponsableHandler(responsableService)
	// Configurar las rutas de responsables en la API
	api.RegisterResponsableRoutes(router, responsableHandler)

	// Crear el servicio de repositorio de files
	fileService := domain.NewFileService(mssqlRepo)
	// Crear el handler de files
	fileHandler := handlers.NewFileHandler(fileService)
	// Configurar las rutas de files en la API
	api.RegisterFileRoutes(router, fileHandler)

	// Configurar el Provider ORCLChecaData
	checaDataRepo := checa_data.NewORCLChecaData(dbCheca)
	// Crear el servicio de repositorio de ChecaData
	checaDataService := domain.NewChecaDataService(checaDataRepo)
	// Crear el handler de ChecaData
	checaDataHandler := handlers.NewChecaDataHandler(checaDataService)
	// Configurar las rutas de ChecaData en la API
	api.RegisterChecaDataRoutes(router, checaDataHandler)

	// Configurar el Provider DUCKFileRepo
	rowsRepo := duck_repo.NewDUCKFileRepo(dbDuck)
	// Crear el servicio de repositorio de Archivos
	rowsService := domain.NewFileService(rowsRepo)
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
