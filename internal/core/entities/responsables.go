package entities

type Responsables struct {
	ID             int `json:"id" db:"id"`
	CoordinadorID  int `json:"coordinador_id" bd:"coordinador_id"`
	Responsable1ID int `json:"responsable1_id" bd:"responsable1_id"`
	Responsable2ID int `json:"responsable2_id" bd:"responsable2_id"`
}
