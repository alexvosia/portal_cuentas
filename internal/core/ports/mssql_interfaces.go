package ports

import "infctas/internal/core/entities"

/*
/getModulo
/getModulos
/createModulo
/setStatusModulo
/setCoordinador
/setResponsable
/setAreas
/setScript
/setMail
/setDescripcion
/setLayOut
*/

type ModuleRepo interface {
	InsertModule(module entities.Module) (int, error)
	FindModuleByID(id int) (*entities.Module, error)
	FindAllModules() ([]entities.Module, error)
	UpdateModule(module entities.Module) error
}
