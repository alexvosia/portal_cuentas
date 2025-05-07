package checa_data

import (
	"database/sql"
	"infctas/internal/core/entities"
	"log"
	"strconv"
)

type ORCLChecaData struct {
	DB *sql.DB
}

// NewORCLChecaData creates a new ORCLChecaData
func NewORCLChecaData(db *sql.DB) *ORCLChecaData {
	return &ORCLChecaData{
		DB: db,
	}
}

func (M ORCLChecaData) FindEmpleadoBySSFF(ssff string) (*entities.User, error) {
	// Convertir el SSFF a int
	ssffInt, err := strconv.Atoi(ssff)
	if err != nil {
		return nil, err
	}
	// Realizar la consulta
	row := M.DB.QueryRow("SELECT HC_NO_EMPLEADO EMPLEADO, EMPLEADO NOMBRE, DAGE_MAIL_EMPRESA MAIL FROM CH2VW_PLAZAS WHERE DAGE_IDSSFF = ?", ssffInt)
	user := &entities.User{}
	err = row.Scan(&user.Empleado, &user.Nombre, &user.Mail)
	if err != nil {
		return nil, err
	}
	user.SSFF = ssffInt
	return user, nil
}

func (M ORCLChecaData) FindAllCoordinadores() ([]entities.ChecaCoordinador, error) {
	rows, err := M.DB.Query("SELECT DAGE_IDSSFF SSFF, EMPLEADO NOMBRE FROM CH2VW_PLAZAS cvp WHERE CPUOP_PUESTO_OP = 'Subdirector'")
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(rows)

	var coordinadores []entities.ChecaCoordinador
	for rows.Next() {
		coordinador := entities.ChecaCoordinador{}
		err = rows.Scan(&coordinador.SSFF, &coordinador.Nombre)
		if err != nil {
			return nil, err
		}
		coordinadores = append(coordinadores, coordinador)
	}
	return coordinadores, nil
}
