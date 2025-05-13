package ports

import "infctas/internal/core/entities"

type ModuleRepo interface {
	InsertModule(module entities.Seccion) (int, error)
	FindModuleByID(id int) (*entities.Seccion, error)
	FindAllModules(rol string) ([]entities.Seccion, error)
	UpdateModule(module entities.Seccion) error
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
