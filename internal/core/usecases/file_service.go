package usecases

import (
	"context"
	"errors"
	"infctas/internal/core/entities"
	"infctas/internal/core/ports"
)

type DuckFileService struct {
	FileRepo ports.FileRepo
}

func NewFileService(fileRepo ports.FileRepo) *DuckFileService {
	return &DuckFileService{
		FileRepo: fileRepo,
	}
}

func (d *DuckFileService) CreateRegistryFile(ctx context.Context, nombre string, idModulo int, tipoRegistro string, creador int) error {
	return errors.New("not implemented")
}

func (d *DuckFileService) GetRegistryFilesByModulo(ctx context.Context, idModulo int) ([]entities.RegistryFileCSV, error) {
	return nil, errors.New("not implemented")
}

func (d *DuckFileService) SetFinRegistryFile(ctx context.Context, idFile int, upper int) (*entities.RegistryFileCSV, error) {
	return nil, errors.New("not implemented")
}

func (d *DuckFileService) InsertFile(file entities.RegistryFileCSV) (int, error) {
	id, err := d.FileRepo.InsertFile(file)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (d *DuckFileService) FindFileByID(id int) (*entities.RegistryFileCSV, error) {
	file, err := d.FileRepo.FindFileByID(id)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (d *DuckFileService) FindFilesByModule(idModule int) ([]entities.RegistryFileCSV, error) {
	files, err := d.FileRepo.FindFilesByModule(idModule)
	if err != nil {
		return nil, err
	}
	return files, nil
}

func (d *DuckFileService) SetFin(id int, fin string, upper int) error {
	err := d.FileRepo.SetFin(id, fin, upper)
	if err != nil {
		return err
	}
	return nil
}
