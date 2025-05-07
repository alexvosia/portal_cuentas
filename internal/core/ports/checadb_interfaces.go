package ports

import "infctas/internal/core/entities"

type ChecaDataRepo interface {
	FindCoordinadores() ([]entities.ChecaCoordinador, error)
	FindResponsablesByCoordinado(idCoordinador int) ([]entities.ChecaResponsable, error)
	FindAreas() ([]entities.ChecaArea, error)
}
