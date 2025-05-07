package entities

type User struct {
	SSFF     int    `json:"ssff" bd:"ssff"`
	Empleado int    `json:"empleado" bd:"empleado"`
	Nombre   string `json:"nombre" bd:"nombre"`
	IdRol    int    `json:"id_rol" bd:"id_rol"`
	Mail     string `json:"mail" bd:"mail"`
	Fin      string `json:"fin" bd:"fin"`
	Creador  int    `json:"creador" bd:"creador"`
	Upper    int    `json:"upper" bd:"upper"`
}
