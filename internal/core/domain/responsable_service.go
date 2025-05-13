package domain

import (
	"context"
	"infctas/internal/core/entities"
	"infctas/internal/core/ports"
)

type ResponsableService struct {
	Repository ports.ResponsableRepo
}

func (r ResponsableService) GetResponsablesConfigurados(ctx context.Context, idCoordinador int) (*entities.Responsables, error) {
	//TODO implement me
	panic("implement me")
}

func NewResponsableService(repository ports.ResponsableRepo) *ResponsableService {
	return &ResponsableService{
		Repository: repository,
	}
}

func (r ResponsableService) InsertResponsable(responsable entities.Responsables) (int, error) {
	id, err := r.Repository.InsertResponsable(responsable)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r ResponsableService) FindResponsableByID(id int) (*entities.Responsables, error) {
	responsable, err := r.Repository.FindResponsableByID(id)
	if err != nil {
		return nil, err
	}
	return responsable, nil
}

func (r ResponsableService) FindResponsableByCoordinador(idCoordinador int) (*entities.Responsables, error) {
	responsable, err := r.Repository.FindResponsableByCoordinador(idCoordinador)
	if err != nil {
		return nil, err
	}
	return responsable, nil
}
