package usecases

import (
	"infctas/internal/core/entities"
	"infctas/internal/core/ports"
)

type DuckFileService struct {
	FileRepo ports.FileRepo
}

func NewDuckFileService(fileRepo ports.FileRepo) *DuckFileService {
	return &DuckFileService{
		FileRepo: fileRepo,
	}
}

func (d *DuckFileService) InsertFile(file entities.FileCSV) (int, error) {
	id, err := d.FileRepo.InsertFile(file)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (d *DuckFileService) FindFileByID(id int) (*entities.FileCSV, error) {
	file, err := d.FileRepo.FindFileByID(id)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (d *DuckFileService) FindFilesByModule(idModule int) ([]entities.FileCSV, error) {
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
