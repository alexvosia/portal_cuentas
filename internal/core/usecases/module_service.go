package usecases

import (
	"context"
	"infctas/internal/core/entities"
	"infctas/internal/core/ports"
)

type ModuleService struct {
	Repo ports.ModuleRepo
}

func NewModuleService(repo ports.ModuleRepo) *ModuleService {
	return &ModuleService{
		Repo: repo,
	}
}

func (m ModuleService) GetModulo(ctx context.Context, id int) (*entities.Module, error) {
	if id == 0 {
		return nil, ports.ErrInvalidData
	}
	module, err := m.Repo.FindModuleByID(id)
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

func (m ModuleService) CreateModulo(ctx context.Context, module entities.Module) (*entities.Module, error) {
	if module.Nombre == "" || module.Descripcion == "" || module.Alcance == "" || module.Coordinador == "" || module.Respoonsable1 == "" || module.Respoonsable2 == "" || module.Columna1 == "" || module.Creador == 0 {
		return nil, ports.ErrInvalidData
	}

	id, err := m.Repo.InsertModule(module)
	if err != nil {
		return nil, err
	}
	module.ID = id
	return &module, nil
}

func (m ModuleService) SetStatusModulo(ctx context.Context, id int, status int) (*entities.Module, error) {
	if id == 0 || status < 0 || status > 1 {
		return nil, ports.ErrInvalidData
	}
	module, err := m.Repo.FindModuleByID(id)
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

func (m ModuleService) SetCoordinador(ctx context.Context, id int, coordinadorid int) (*entities.Module, error) {
	if id == 0 || coordinadorid == 0 {
		return nil, ports.ErrInvalidData
	}
	module, err := m.Repo.FindModuleByID(id)
	if err != nil {
		return nil, err
	}
	module.CoordinadorID = coordinadorid
	err = m.Repo.UpdateModule(*module)
	if err != nil {
		return nil, err
	}
	return module, nil
}

/*
func (m ModuleService) SetResponsable(ctx context.Context, id int, responsableid int, index int) (*entities.Module, error) {
	if id == 0 || responsableid == 0 {
		return nil, ports.ErrInvalidData
	}
	module, err := m.Repo.FindModuleByID(id)
	if err != nil {
		return nil, err
	}
	if index == 1 {
		module.Responsable1ID = responsableid
	}
	if index == 2 {
		module.Responsable2ID = responsableid
	}
	err = m.Repo.UpdateModule(*module)
	if err != nil {
		return nil, err
	}
	return module, nil
}
*/

func (m ModuleService) SetAreas(ctx context.Context, id int, areas []entities.Area) (*entities.Module, error) {
	if id == 0 || len(areas) == 0 {
		return nil, ports.ErrInvalidData
	}
	module, err := m.Repo.FindModuleByID(id)
	if err != nil {
		return nil, err
	}
	module.Areas = areas
	err = m.Repo.UpdateModule(*module)
	if err != nil {
		return nil, err
	}
	return module, nil
}

func (m ModuleService) SetScript(ctx context.Context, id int, script string) (*entities.Module, error) {
	if id == 0 || script == "" {
		return nil, ports.ErrInvalidData
	}
	module, err := m.Repo.FindModuleByID(id)
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

func (m ModuleService) SetMail(ctx context.Context, id int, mail string) (*entities.Module, error) {
	if id == 0 || mail == "" {
		return nil, ports.ErrInvalidData
	}
	module, err := m.Repo.FindModuleByID(id)
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

func (m ModuleService) SetDescripcion(ctx context.Context, id int, descripcion string) (*entities.Module, error) {
	if id == 0 || descripcion == "" {
		return nil, ports.ErrInvalidData
	}
	module, err := m.Repo.FindModuleByID(id)
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

/*
func (m ModuleService) SetLayOut(ctx context.Context, id int, layoutid int) (*entities.Module, error) {
	if id == 0 || layoutid == 0 {
		return nil, ports.ErrInvalidData
	}
	module, err := m.Repo.FindModuleByID(id)
	if err != nil {
		return nil, err
	}
	module.LayOutID = layoutid
	err = m.Repo.UpdateModule(*module)
	if err != nil {
		return nil, err
	}
	return module, nil
}
*/
