package config

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/godror/godror"
	_ "github.com/marcboeker/go-duckdb"
)

// DBManager maneja las conexiones a las bases de datos
type DBManager struct {
	checaConn *sql.DB
	picConn   *sql.DB
	duckConn  *sql.DB

	checaMutex sync.Mutex
	picMutex   sync.Mutex
	duckMutex  sync.Mutex
}

var (
	instance *DBManager
	once     sync.Once
)

// GetDBManager retorna la instancia singleton del DBManager
func GetDBManager() *DBManager {
	once.Do(func() {
		instance = &DBManager{}
	})
	return instance
}

// GetConn obtiene una conexión a la base de datos según el entorno
func (m *DBManager) GetConn(entorno string) (*sql.DB, error) {
	switch entorno {
	case "checa":
		return m.getOracleConn()
	case "pic":
		return m.getMSSQLConn()
	case "duck":
		return m.getDuckDBConn()
	default:
		return nil, errors.New("entorno no válido")
	}
}

// GetHTTPServerConfig obtiene la configuración del servidor HTTP
func (m *DBManager) GetHTTPServerConfig() (string, string, string, error) {
	config, err := m.loadConfig()
	if err != nil {
		return "", "", "", err
	}

	port := config["http"]["port"]
	cert := config["http"]["cert"]
	key := config["http"]["key"]

	if port == "" || cert == "" || key == "" {
		return "", "", "", errors.New("configuración del servidor HTTP incompleta")
	}

	return port, cert, key, nil
}

// CloseAll cierra todas las conexiones activas
func (m *DBManager) CloseAll() error {
	var errs []error

	if err := m.closeOracleConn(); err != nil {
		errs = append(errs, err)
	}

	if err := m.closeMSSQLConn(); err != nil {
		errs = append(errs, err)
	}

	if err := m.closeDuckDBConn(); err != nil {
		errs = append(errs, err)
	}

	if len(errs) > 0 {
		return fmt.Errorf("errores al cerrar conexiones: %v", errs)
	}
	return nil
}

// Métodos específicos para cada base de datos
func (m *DBManager) getOracleConn() (*sql.DB, error) {
	m.checaMutex.Lock()
	defer m.checaMutex.Unlock()

	if m.checaConn != nil {
		return m.checaConn, nil
	}

	connStr, err := m.createOracleConnStr()
	if err != nil {
		return nil, err
	}

	db, err := sql.Open("godror", connStr)
	if err != nil {
		return nil, fmt.Errorf("error al conectar a Oracle: %v", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error al hacer ping a Oracle: %v", err)
	}

	m.checaConn = db
	return m.checaConn, nil
}

func (m *DBManager) getMSSQLConn() (*sql.DB, error) {
	m.picMutex.Lock()
	defer m.picMutex.Unlock()

	if m.picConn != nil {
		return m.picConn, nil
	}

	connStr, err := m.createMSSQLConnStr()
	if err != nil {
		return nil, err
	}

	db, err := sql.Open("sqlserver", connStr)
	if err != nil {
		return nil, fmt.Errorf("error al conectar a MSSQL: %v", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error al hacer ping a MSSQL: %v", err)
	}

	m.picConn = db
	return m.picConn, nil
}

func (m *DBManager) getDuckDBConn() (*sql.DB, error) {
	m.duckMutex.Lock()
	defer m.duckMutex.Unlock()

	if m.duckConn != nil {
		return m.duckConn, nil
	}

	connStr, err := m.createDuckDBConnStr()
	if err != nil {
		return nil, err
	}

	db, err := sql.Open("duckdb", connStr)
	if err != nil {
		return nil, fmt.Errorf("error al conectar a DuckDB: %v", err)
	}

	m.duckConn = db
	return m.duckConn, nil
}

// Métodos para cerrar conexiones específicas
func (m *DBManager) closeOracleConn() error {
	m.checaMutex.Lock()
	defer m.checaMutex.Unlock()

	if m.checaConn != nil {
		err := m.checaConn.Close()
		m.checaConn = nil
		return err
	}
	return nil
}

func (m *DBManager) closeMSSQLConn() error {
	m.picMutex.Lock()
	defer m.picMutex.Unlock()

	if m.picConn != nil {
		err := m.picConn.Close()
		m.picConn = nil
		return err
	}
	return nil
}

func (m *DBManager) closeDuckDBConn() error {
	m.duckMutex.Lock()
	defer m.duckMutex.Unlock()

	if m.duckConn != nil {
		err := m.duckConn.Close()
		m.duckConn = nil
		return err
	}
	return nil
}

// Métodos para crear strings de conexión
func (m *DBManager) createOracleConnStr() (string, error) {
	config, err := m.loadConfig()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("user=%s password=%s connectString=%s:%s/%s",
		config["checa"]["user"],
		config["checa"]["password"],
		config["checa"]["host"],
		config["checa"]["port"],
		config["checa"]["sid"]), nil
}

func (m *DBManager) createMSSQLConnStr() (string, error) {
	config, err := m.loadConfig()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s",
		config["pic"]["user"],
		config["pic"]["password"],
		config["pic"]["host"],
		config["pic"]["port"],
		config["pic"]["database"]), nil
}

func (m *DBManager) createDuckDBConnStr() (string, error) {
	config, err := m.loadConfig()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("file=%s", config["duck"]["file"]), nil
}

func (m *DBManager) loadConfig() (map[string]map[string]string, error) {
	err := godotenv.Load("./config.env")
	if err != nil {
		log.Printf("Warning: no se pudo cargar .env file: %v", err)
	}

	config := make(map[string]map[string]string)

	config["checa"] = map[string]string{
		"user":     getEnv("USER_NAME_CHECA", ""),
		"password": getEnv("PASSWORD_CHECA", ""),
		"sid":      getEnv("SID_CHECA", ""),
		"host":     getEnv("DB_HOST_CHECA", ""),
		"port":     getEnv("DB_PORT_CHECA", "1521"),
	}

	config["pic"] = map[string]string{
		"user":     getEnv("USER_NAME_PIC", ""),
		"password": getEnv("PASSWORD_PIC", ""),
		"host":     getEnv("DB_HOST_PIC", ""),
		"port":     getEnv("DB_PORT_PIC", "1433"),
		"database": getEnv("DB_NAME_PIC", ""),
	}

	config["duck"] = map[string]string{
		"file": getEnv("DB_FILE_DUCK", "./local_db/duckdb.db"),
	}

	config["http"] = map[string]string{
		"port": getEnv("HTTP_SERVER_PORT", "443"),
		"cert": getEnv("HTTP_SERVER_CERT_FILE", "/dev/certs/certificado.crt"),
		"key":  getEnv("HTTP_SERVER_KEY_FILE", "/dev/certs/clave-privada.key"),
	}

	// Validar configuraciones requeridas
	if config["checa"]["user"] == "" || config["checa"]["password"] == "" || config["checa"]["host"] == "" {
		return nil, errors.New("configuración de Oracle incompleta")
	}

	if config["pic"]["user"] == "" || config["pic"]["password"] == "" || config["pic"]["host"] == "" {
		return nil, errors.New("configuración de MSSQL incompleta")
	}

	if config["duck"]["file"] == "" {
		return nil, errors.New("configuración de DuckDB incompleta")
	}

	//if config["http"]["port"] == "" || config["http"]["cert"] == "" || config["http"]["key"] == "" {
	//	return nil, errors.New("configuración de HTTP incompleta")
	//}

	return config, nil
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
