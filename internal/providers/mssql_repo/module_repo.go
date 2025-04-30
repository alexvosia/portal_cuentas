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
	exec, err := M.DB.Exec("INSERT INTO Modulos (Nombre, Descripcion, Status, CoordinadorID, Responsable1ID, Responsable2ID) VALUES (?, ?, ?, ?, ?, ?)", module.Name, module.Description, module.Status, module.CoordinadorID, module.Responsable1ID, module.Responsable2ID)
	if err != nil {
		return 0, err
	}
	id, err := exec.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (M MSSQLModuleRepo) FindModuleByID(id int) (*entities.Module, error) {
	row := M.DB.QueryRow("SELECT * FROM Modulos WHERE ID = ?", id)
	module := &entities.Module{}
	err := row.Scan(&module.ID, &module.Name, &module.Description, &module.Status, &module.Script, &module.Mail, &module.FechaCreacion, &module.FechaModificacion, &module.FechaEliminacion, &module.CoordinadorID, &module.Responsable1ID, &module.Responsable2ID, &module.LayOutID)
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
		err := rows.Scan(&module.ID, &module.Name, &module.Description, &module.Status, &module.Script, &module.Mail, &module.FechaCreacion, &module.FechaModificacion, &module.FechaEliminacion, &module.CoordinadorID, &module.Responsable1ID, &module.Responsable2ID, &module.LayOutID)
		if err != nil {
			return nil, err
		}
		modules = append(modules, module)
	}
	return modules, nil
}

func (M MSSQLModuleRepo) UpdateModule(module entities.Module) error {
	_, err := M.DB.Exec("UPDATE Modulos SET Nombre = ?, Descripcion = ?, Status = ?, CoordinadorID = ?, Responsable1ID = ?, Responsable2ID = ? WHERE ID = ?", module.Name, module.Description, module.Status, module.CoordinadorID, module.Responsable1ID, module.Responsable2ID, module.ID)
	if err != nil {
		return err
	}
	return nil
}
