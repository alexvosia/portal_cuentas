package main

import (
	"context"
	"crypto/tls"
	"database/sql"
	"errors"
	"infctas/internal/config"
	"infctas/internal/core/domain"
	"infctas/internal/drivers/api"
	"infctas/internal/drivers/api/handlers"
	"infctas/internal/providers/checa_data"
	"infctas/internal/providers/duck_repo"
	"infctas/internal/providers/pic_repo"
	"log"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	// Crear DBManager para cargar la configuracion de las bases de datos
	dbManager := config.GetDBManager()

	// Funcion local para cargar la configuracion de las bases de datos
	localConfigDB := func(dbManager *config.DBManager) (*sql.DB, *sql.DB, *sql.DB, error) {
		dbCheca, err := dbManager.GetConn("checa")
		if err != nil {
			log.Fatalf("Error al conectar a la base de datos Checa: %v", err)
			return nil, nil, nil, err
		}
		dbPic, err := dbManager.GetConn("pic")
		if err != nil {
			log.Fatalf("Error al conectar a la base de datos PIC: %v", err)
			return nil, nil, nil, err
		}
		dbDuck, err := dbManager.GetConn("duck")
		if err != nil {
			log.Fatalf("Error al conectar a la base de datos Duck: %v", err)
			return nil, nil, nil, err
		}
		return dbCheca, dbPic, dbDuck, nil
	}

	// Cargar la configuracion de las bases de datos
	dbCheca, dbPic, dbDuck, err := localConfigDB(dbManager)
	if err != nil {
		log.Fatalf("Error al cargar la configuracion de las bases de datos: %v", err)
		return
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

	// Obtener la configuración del servidor HTTP
	port, cert, key, err := dbManager.GetHTTPServerConfig()
	if err != nil {
		log.Fatalf("Error al obtener la configuración del servidor HTTP: %v", err)
		return
	}

	// Configuración del servidor TLS (puerto 443)
	server := &http.Server{
		Addr:    ":" + port, // Puerto HTTPS
		Handler: router,
		TLSConfig: &tls.Config{
			MinVersion: tls.VersionTLS13, // TLS más seguro
		},
	}

	// Crear un contexto con un canal de cancelación
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Canal para recibir señales de interrupción (Ctrl+C)
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	// Lanzar el servidor en una goroutine
	go func() {
		log.Println("Servidor escuchando en https://localhost:" + port)
		if err := server.ListenAndServeTLS(cert, key); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Error al iniciar el servidor: %v", err)
		}
	}()

	// Esperar a que el sistema reciba una señal de interrupción
	<-stop
	log.Println("Señal de interrupción recibida. Deteniendo el servidor...")
	log.Println("Esperando a terminar las peticiones en transito...")

	// Se crea un contexto con un cancel para el apagado del servidor
	shutdownCtx, shutdownCancel := context.WithCancel(ctx)
	defer shutdownCancel()

	// Deteniendo la aceptación de nuevas conexiones, pero esperando a que las actuales terminen
	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("Error durante el apagado del servidor: %v", err)
	}
	log.Println("Servidor detenido con éxito")
}
