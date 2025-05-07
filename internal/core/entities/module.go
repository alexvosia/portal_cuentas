package entities

type Module struct {
	ID                int    `json:"id" db:"id"`
	Name              string `json:"name" db:"name"`
	Description       string `json:"description" db:"description"`
	Status            int    `json:"status" db:"status"`
	Script            string `json:"script" db:"script"`
	Mail              string `json:"mail" db:"mail"`
	FechaCreacion     string `json:"fechacreacion" db:"fechacreacion"`
	FechaModificacion string `json:"fechamodificacion" db:"fechamodificacion"`
	FechaEliminacion  string `json:"fechaeliminacion" db:"fechaeliminacion"`
	CoordinadorID     int    `json:"coordinadorid" db:"coordinadorid"`
	Columna1          string `json:"columna1" db:"columna1"`
	Columna2          string `json:"columna2" db:"columna2"`
	Columna3          string `json:"columna3" db:"columna3"`
	Columna4          string `json:"columna4" db:"columna4"`
	Columna5          string `json:"columna5" db:"columna5"`
	Columna6          string `json:"columna6" db:"columna6"`
	Columna7          string `json:"columna7" db:"columna7"`
	Columna8          string `json:"columna8" db:"columna8"`
	Columna9          string `json:"columna9" db:"columna9"`
	Areas             string `json:"areas" db:"areas"`
}
