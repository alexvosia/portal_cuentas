// Package entities: Entidades para el módulo de administracion de Modulos, Usuarios, Roles, Permisos y Archivos de Registro
package entities

import "time"

type Seccion struct {
	Id          int
	Nombre      string
	Descripcion string
	Estado      int
	Alcance     []int
	Responsable Responsables
	Script      string
	Mail        string
	Columna1    string
	Columna2    string
	Columna3    string
	Columna4    string
	Columna5    string
	Columna6    string
	Columna7    string
	Columna8    string
	Columna9    string
	Inicio      time.Time
	Fin         time.Time
	Creador     int
	Upper       int
}

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

type RegistryFileCSV struct {
	Id           int
	Nombre       string
	NombreModulo string
	TipoRegistro string
	Path         string
	Size         int
	Inicio       time.Time
	Fin          time.Time
	Creador      int
	Upper        int
}

type Responsables struct {
	Id           int
	Coordinador  ChecaEmpleado
	Responsable1 ChecaEmpleado
	Responsable2 ChecaEmpleado
}

type User struct {
	SSFF     int
	Empleado int
	Nombre   string
	IdRol    int
	Mail     string
	Fin      time.Time
	Creador  int
	Upper    int
}
