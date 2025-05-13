package domain

import (
	"context"
	"infctas/internal/core/entities"
	"infctas/internal/core/ports"
	"strconv"
	"strings"
)

type ModuleService struct {
	Repo ports.ModuleRepo
}

func NewModuleService(repo ports.ModuleRepo) *ModuleService {
	return &ModuleService{
		Repo: repo,
	}
}

func (m ModuleService) CreateModulo(ctx context.Context, module entities.Module) (*entities.Module, error) {
	if module.Nombre == "" || module.Descripcion == "" || module.Alcance == "" || module.Coordinador == "" || module.Responsable1 == "" || module.Responsable2 == "" || module.Columna1 == "" || module.Creador == 0 {
		return nil, ports.ErrInvalidData
	}

	id, err := m.Repo.InsertModule(module)
	if err != nil {
		return nil, err
	}
	module.Id = id
	return &module, nil
}

func (m ModuleService) GetModuloById(ctx context.Context, idModulo int) (*entities.Module, error) {
	if idModulo == 0 {
		return nil, ports.ErrInvalidData
	}
	module, err := m.Repo.FindModuleByID(idModulo)
	if err != nil {
		return nil, ports.ErrModuleNotFound
	}

	return module, nil
}

func (m ModuleService) GetModulos(ctx context.Context, rol string) ([]entities.Module, error) {
	modules, err := m.Repo.FindAllModules(rol)
	if err != nil {
		return nil, err
	}
	return modules, nil
}

func (m ModuleService) SetStatusModulo(ctx context.Context, idModulo int, status int, upper int) (*entities.Module, error) {
	if idModulo == 0 || status < 0 || status > 1 {

		return nil, ports.ErrInvalidData
	}
	module, err := m.Repo.FindModuleByID(idModulo)
	if err != nil {
		return nil, err
	}
	module.Estado = status
	err = m.Repo.UpdateModule(*module)
	if err != nil {
		return nil, err
	}
	return module, nil
}

func (m ModuleService) SetCoordinador(ctx context.Context, idModulo int, idCoordinador int, upper int) (*entities.Module, error) {
	if idModulo == 0 || idCoordinador == 0 {
		return nil, ports.ErrInvalidData
	}
	module, err := m.Repo.FindModuleByID(idModulo)

	if err != nil {
		return nil, err
	}
	module.Coordinador = strconv.Itoa(idCoordinador)
	err = m.Repo.UpdateModule(*module)
	if err != nil {
		return nil, err
	}
	return module, nil
}

func (m ModuleService) SetResponsable(ctx context.Context, idModulo int, idCoordinador int, idResponsable1 int, idResponsable2 int, upper int) (*entities.Module, error) {
	if idModulo == 0 || idResponsable1 == 0 {
		return nil, ports.ErrInvalidData
	}
	module, err := m.Repo.FindModuleByID(idModulo)
	if err != nil {
		return nil, err
	}
	module.Responsable1 = idResponsable1
	module.Responsable2 = idResponsable2
	err = m.Repo.UpdateModule(*module)
	if err != nil {
		return nil, err
	}
	return module, nil
}

func (m ModuleService) SetAreas(ctx context.Context, idModulo int, areas string, upper int) (*entities.Module, error) {
	if idModulo == 0 || len(areas) == 0 {
		return nil, ports.ErrInvalidData
	}
	module, err := m.Repo.FindModuleByID(idModulo)
	if err != nil {
		return nil, err
	}
	module.Alcance = areas
	err = m.Repo.UpdateModule(*module)
	if err != nil {
		return nil, err
	}
	return module, nil
}

func (m ModuleService) SetScript(ctx context.Context, idModulo int, script string, upper int) (*entities.Module, error) {
	if idModulo == 0 || script == "" {
		return nil, ports.ErrInvalidData
	}
	module, err := m.Repo.FindModuleByID(idModulo)
	if err != nil {
		return nil, err
	}
	module.Script = script
	err = m.Repo.UpdateModule(*module)
	if err != nil {
		return nil, err
	}
	return module, nil
}

func (m ModuleService) SetMail(ctx context.Context, idModulo int, mail string, upper int) (*entities.Module, error) {
	if idModulo == 0 || mail == "" {
		return nil, ports.ErrInvalidData
	}
	module, err := m.Repo.FindModuleByID(idModulo)
	if err != nil {
		return nil, err
	}
	module.Mail = mail
	err = m.Repo.UpdateModule(*module)
	if err != nil {
		return nil, err
	}
	return module, nil
}

func (m ModuleService) SetDescripcion(ctx context.Context, idModulo int, descripcion string, upper int) (*entities.Module, error) {
	if idModulo == 0 || descripcion == "" {
		return nil, ports.ErrInvalidData
	}
	module, err := m.Repo.FindModuleByID(idModulo)
	if err != nil {
		return nil, err
	}
	module.Descripcion = descripcion
	err = m.Repo.UpdateModule(*module)
	if err != nil {
		return nil, err
	}
	return module, nil
}
