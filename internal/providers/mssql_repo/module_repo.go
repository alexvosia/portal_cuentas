package mssql_repo

import (
	"database/sql"
	"infctas/internal/core/entities"
	"log"
)

type MSSQLModuleRepo struct {
	DB *sql.DB
}

// NewMSSQLModuleRepo creates a new MSSQLModuleRepo
func NewMSSQLModuleRepo(db *sql.DB) *MSSQLModuleRepo {
	return &MSSQLModuleRepo{
		DB: db,
	}
}

func (M MSSQLModuleRepo) InsertModule(module entities.Module) (int, error) {
	exec, err := M.DB.Exec("INSERT INTO Modulos (Nombre, Descripcion, CoordinadorID) VALUES (?, ?, ?, ?, ?, ?)", module.Nombre, module.Descripcion, module.CoordinadorID)
	if err != nil {
		return 0, err
	}
	id, err := exec.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nilw
}

func (M MSSQLModuleRepo) FindModuleByID(id int) (*entities.Module, error) {
	row := M.DB.QueryRow("SELECT * FROM Modulos WHERE ID = ?", id)
	module := &entities.Module{}
	err := row.Scan(&module.ID, &module.Nombre, &module.Descripcion, &module.Estado, &module.Script, &module.Mail, &module.FechaCreacion, &module.FechaModificacion, &module.FechaEliminacion, &module.CoordinadorID)
	if err != nil {
		return nil, err
	}
	return module, nil
}

func (M MSSQLModuleRepo) FindAllModules() ([]entities.Module, error) {
	rows, err := M.DB.Query("SELECT * FROM Modulos WHERE Status != 'eliminado' AND fechaeliminacion IS NULL")
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(rows)

	var modules []entities.Module
	for rows.Next() {
		module := entities.Module{}
		err := rows.Scan(&module.ID, &module.Nombre, &module.Descripcion, &module.Estado, &module.Script, &module.Mail, &module.FechaCreacion, &module.FechaModificacion, &module.FechaEliminacion, &module.CoordinadorID)
		if err != nil {
			return nil, err
		}
		modules = append(modules, module)
	}
	return modules, nil
}

func (M MSSQLModuleRepo) UpdateModule(module entities.Module) error {
	_, err := M.DB.Exec("UPDATE Modulos SET Nombre = ?, Descripcion = ?, Status = ?, CoordinadorID = ? WHERE ID = ?", module.Nombre, module.Descripcion, module.Estado, module.CoordinadorID, module.ID)
	if err != nil {
		return err
	}
	return nil
}
