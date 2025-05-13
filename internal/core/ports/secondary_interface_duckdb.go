package ports

import "infctas/internal/core/entities"

type DuckDBRepo interface {
	InsertFile(file entities.RegistryFileCSV) (int, error)
	FindFileByID(id int) (*entities.RegistryFileCSV, error)
	FindFilesByModule(idModule int) ([]entities.RegistryFileCSV, error)
	SetFin(id int, fin string, upper int) error
}
