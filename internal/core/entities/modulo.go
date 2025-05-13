package entities

import "time"

type Module struct {
	Id          int
	Nombre      string
	Descripcion string
	Estado      int
	Alcance     string
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
	Creador     User
	Upper       User
}
