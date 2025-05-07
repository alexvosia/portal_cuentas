package entities

type Responsables struct {
	Id             int    `json:"id"`
	IdCoordinador  int    `json:"usuario_coordinador"`
	IdResponsable1 int    `json:"usuario_responsable1"`
	IdResponsable2 int    `json:"usuario_responsable2"`
	Fin            string `json:"fin"`
	Creador        int    `json:"creador"`
	Upper          int    `json:"upper"`
}
