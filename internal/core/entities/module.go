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
	Areas             []Area `json:"areas" db:"areas"`
}
