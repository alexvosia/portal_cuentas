package ports

import "infctas/internal/core/entities"

type ModuleRepo interface {
	InsertModule(module entities.Module) (int, error)
	FindModuleByID(id int) (*entities.Module, error)
	FindAllModules(rol string) ([]entities.Module, error)
	UpdateModule(module entities.Module) error
}

type FileRepo interface {
	InsertFile(file entities.RegistryFileCSV) (int, error)
	FindFileByID(id int) (*entities.RegistryFileCSV, error)
	FindFilesByModule(idModule int) ([]entities.RegistryFileCSV, error)
	SetFin(id int, fin string, upper int) error
}

type ResponsableRepo interface {
	InsertResponsable(responsable entities.Responsables) (int, error)
	FindResponsableByID(id int) (*entities.Responsables, error)
	FindResponsableByCoordinador(idCoordinador int) (*entities.Responsables, error)
}
