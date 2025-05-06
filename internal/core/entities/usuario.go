package entities

type User struct {
	SSFF      int    `json:"ssf" bd:"ssf"`
	Empleado  int    `json:"empleado" bd:"empleado"`
	IdRole    string `json:"rol" bd:"rol"`
	Mail      string `json:"mail" bd:"mail"`
	Nombre    string `json:"nombre" bd:"nombre"`
	Apellidos string `json:"apellido" bd:"apellido"`
}
