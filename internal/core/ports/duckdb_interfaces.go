package ports

import "infctas/internal/core/entities"

type DuckDBRepo interface {
	InsertFile(file entities.FileCSV) (int, error)
	FindFileByID(id int) (*entities.FileCSV, error)
	FindFilesByModule(idModule int) ([]entities.FileCSV, error)
	SetFin(id int, fin string, upper int) error
}
