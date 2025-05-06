package entities

type Responsables struct {
	Id             int    `json:"id" db:"id"`
	IdCoordinador  int    `json:"usuario_coordinador" bd:"usuario_coordinador"`
	IdResponsable1 int    `json:"usuario_responsable1" bd:"usuario_responsable1"`
	IdResponsable2 int    `json:"usuario_responsable2" bd:"usuario_responsable2"`
	Fin            string `json:"fin" db:"fin"`
	Creador        int    `json:"creador" db:"creador"`
	Upper          int    `json:"upper" db:"upper"`
}
