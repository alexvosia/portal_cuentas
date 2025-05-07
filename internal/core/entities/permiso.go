package entities

type Permiso struct {
	Id     int    `json:"id" db:"id"`
	Codigo string `json:"codigo" db:"codigo"`
	Accion string `json:"accion" db:"accion"`
}

type Role struct {
	Id   int    `json:"id" db:"id"`
	Role string `json:"rol" db:"rol"`
}

type RolePermiso struct {
	Id        int `json:"id" db:"id"`
	IdPermiso int `json:"id_permiso" db:"id_permiso"`
	IdRole    int `json:"id_rol" db:"id_rol"`
}
