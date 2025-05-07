package entities

type User struct {
	SSFF     int    `json:"ssff"`
	Empleado int    `json:"empleado"`
	Nombre   string `json:"nombre"`
	IdRol    int    `json:"id_rol"`
	Mail     string `json:"mail"`
	Fin      string `json:"fin"`
	Creador  int    `json:"creador"`
	Upper    int    `json:"upper"`
}
