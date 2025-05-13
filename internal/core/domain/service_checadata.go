package domain

import (
	"context"
	"infctas/internal/core/entities"
	"infctas/internal/core/ports"
)

type ChecaDataService struct {
	Data ports.ChecaDataAPI
}

func NewChecaDataService(data ports.ChecaDataAPI) *ChecaDataService {
	return &ChecaDataService{
		Data: data,
	}
}

func (c ChecaDataService) GetCoordinadores(ctx context.Context) ([]entities.ChecaCoordinador, error) {
	coordinadores, err := c.Data.FindAllCoordinadores()
	if err != nil {
		return nil, err
	}
	return coordinadores, nil
}

func (c ChecaDataService) GetResponsables(ctx context.Context, idCoordinador int) (interface{}, error) {
	if idCoordinador == 0 {
		return nil, ports.ErrInvalidData
	}
	responsables, err := c.Data.FindResponsableByCoordinador(idCoordinador)
	if err != nil {
		return nil, err
	}
	return responsables, nil
}

func (c ChecaDataService) GetAreas(ctx context.Context) ([]entities.ChecaArea, error) {
	areas, err := c.Data.FindAllAreas()
	if err != nil {
		return nil, err
	}
	return areas, nil
}
