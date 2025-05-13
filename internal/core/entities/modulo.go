package entities

import "time"

type Module struct {
	Id          int          `json:"id"`
	Nombre      string       `json:"nombre"`
	Descripcion string       `json:"descripcion"`
	Estado      int          `json:"estado"`
	Alcance     string       `json:"alcance"`
	Coordinador Responsables `json:"coordinador"`
	Script      string       `json:"script"`
	Mail        string       `json:"mail"`
	Columna1    string       `json:"columna1"`
	Columna2    string       `json:"columna2"`
	Columna3    string       `json:"columna3"`
	Columna4    string       `json:"columna4"`
	Columna5    string       `json:"columna5"`
	Columna6    string       `json:"columna6"`
	Columna7    string       `json:"columna7"`
	Columna8    string       `json:"columna8"`
	Columna9    string       `json:"columna9"`
	Inicio      time.Time    `json:"inicio"`
	Fin         time.Time    `json:"fin"`
	Creador     User         `json:"creador"`
	Upper       User         `json:"upper"`
	repo        ChecaDataRepop
}
