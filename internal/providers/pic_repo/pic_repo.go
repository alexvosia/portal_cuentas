package pic_repo

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
func (M MSSQLModuleRepo) InsertResponsable(responsable entities.Responsables) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (M MSSQLModuleRepo) FindResponsableByCoordinador(idCoordinador int) (*entities.Responsables, error) {
	//TODO implement me
	panic("implement me")
}

func (M MSSQLModuleRepo) FindResponsableByID(id int) (*entities.Responsables, error) {
	row := M.DB.QueryRow("SELECT * FROM coordinador_analista WHERE ID = ?", id)
	responsable := &entities.Responsables{}
	err := row.Scan(&responsable.Id, &responsable.IdCoordinador, &responsable.IdResponsable1, &responsable.IdResponsable2, &responsable.Fin, &responsable.Creador, &responsable.Upper)
	if err != nil {
		return nil, err
	}
	return responsable, nil
}

func (M MSSQLModuleRepo) InsertModule(module entities.Module) (int, error) {
	exec, err := M.DB.Exec("INSERT INTO Modulos (Nombre, Descripcion,Coordinador) VALUES (?, ?, ?)", module.Nombre, module.Descripcion, module.Coordinador)
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
	inResp := &entities.Responsables{}
	coord := &entities.User{}
	resp1 := &entities.User{}
	resp2 := &entities.User{}
	err := row.Scan(&module.Id, &module.Nombre, &module.Descripcion, &module.Estado, &module.Alcance, &inResp.Id, &module.Script, &module.Mail, &module.Columna1, &module.Columna2, &module.Columna3, &module.Columna4, &module.Columna5, &module.Columna6, &module.Columna7, &module.Columna8, &module.Columna9, &module.Inicio, &module.Fin, &module.Creador, &module.Upper)
	if err != nil {
		return nil, err
	}

	inResp, err = M.FindResponsableByID(inResp.Id)
	if err != nil {
		return nil, err
	}

	coord, err = M.FindUserBySSFF(inResp.IdCoordinador)
	if err != nil {
		return nil, err
	}
	resp1, err = M.FindUserBySSFF(inResp.IdResponsable1)
	if err != nil {
		return nil, err
	}
	resp2, err = M.FindUserBySSFF(inResp.IdResponsable2)
	if err != nil {
		return nil, err
	}

	module.Coordinador = coord.Nombre
	module.Responsable1 = resp1.Nombre
	module.Responsable2 = resp2.Nombre

	return module, nil
}

func (M MSSQLModuleRepo) FindAllModules(role string) ([]entities.Module, error) {
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
		inResp := &entities.Responsables{}
		coord := &entities.User{}
		resp1 := &entities.User{}
		resp2 := &entities.User{}
		err := rows.Scan(&module.Id, &module.Nombre, &module.Descripcion, &module.Estado, &module.Alcance, &inResp.Id, &module.Script, &module.Mail, &module.Columna1, &module.Columna2, &module.Columna3, &module.Columna4, &module.Columna5, &module.Columna6, &module.Columna7, &module.Columna8, &module.Columna9, &module.Inicio, &module.Fin, &module.Creador, &module.Upper)
		if err != nil {
			return nil, err
		}
		inResp, err = M.FindResponsableByID(inResp.Id)
		if err != nil {
			return nil, err
		}
		coord, err = M.FindUserBySSFF(inResp.IdCoordinador)
		if err != nil {
			return nil, err
		}
		resp1, err = M.FindUserBySSFF(inResp.IdResponsable1)
		if err != nil {
			return nil, err
		}
		resp2, err = M.FindUserBySSFF(inResp.IdResponsable2)
		if err != nil {
			return nil, err
		}

		module.Coordinador = coord.Nombre
		module.Responsable1 = resp1.Nombre
		module.Responsable2 = resp2.Nombre

		modules = append(modules, module)
	}
	return modules, nil
}

func (M MSSQLModuleRepo) UpdateModule(module entities.Module) error {
	_, err := M.DB.Exec("UPDATE Modulos SET Nombre = ?, Descripcion = ?, Status = ?, CoordinadorID = ? WHERE ID = ?", module.Nombre, module.Descripcion, module.Estado, module.Coordinador, module.Id)
	if err != nil {
		return err
	}
	return nil
}

func (M MSSQLModuleRepo) FindUserBySSFF(id int) (*entities.User, error) {
	row := M.DB.QueryRow("SELECT ssff, empleado, nombre, id_rol, mail, fin, creador, upper FROM usuarios WHERE SSFF = ?", id)
	user := &entities.User{}
	err := row.Scan(&user.SSFF, &user.Empleado, &user.Nombre, &user.IdRol, &user.Mail, &user.Fin, &user.Creador, &user.Upper)
	if err != nil {
		return nil, err
	}
	return user, nil
}
