package entities

type Permiso struct {
	Id     int    `json:"id"`
	Codigo string `json:"codigo"`
	Accion string `json:"accion"`
}

type Role struct {
	Id   int    `json:"id"`
	Role string `json:"rol"`
}

type RolePermiso struct {
	Id        int `json:"id"`
	IdPermiso int `json:"id_permiso"`
	IdRole    int `json:"id_rol"`
}
