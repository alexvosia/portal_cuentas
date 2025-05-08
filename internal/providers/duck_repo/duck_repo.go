package duck_repo

import (
	"database/sql"
	"infctas/internal/core/entities"
)

type DUCKFileRepo struct {
	DB *sql.DB
}

func NewDUCKFileRepo(db *sql.DB) *DUCKFileRepo {
	return &DUCKFileRepo{
		DB: db,
	}
}

func (d *DUCKFileRepo) InsertFile(file entities.RegistryFileCSV) (int, error) {
	return 0, nil
}

func (d *DUCKFileRepo) FindFileByID(id int) (*entities.RegistryFileCSV, error) {
	return nil, nil
}

func (d *DUCKFileRepo) FindFilesByModule(idModule int) ([]entities.RegistryFileCSV, error) {
	return nil, nil
}

func (d *DUCKFileRepo) SetFin(id int, fin string, upper int) error {
	return nil
}
