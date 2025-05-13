package entities

type Permiso struct {
	Id     int
	Codigo string
	Accion string
}

type Role struct {
	Id   int
	Role string
}

type RolePermiso struct {
	Id      int
	Permiso Permiso
	Role    Role
}
