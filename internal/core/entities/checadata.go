package entities

type ChecaCoordinador struct {
	SSFF   int    `json:"ssff" bd:"ssff"`
	Nombre string `json:"nombre" bd:"nombre"`
}

type ChecaResponsable struct {
	SSFF   int    `json:"ssff" bd:"ssff"`
	Nombre string `json:"nombre" bd:"nombre"`
}

type ChecaArea struct {
	Id   int    `json:"id" db:"id"`
	Area string `json:"area" bd:"area"`
}
