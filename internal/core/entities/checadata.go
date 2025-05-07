package entities

type ChecaCoordinador struct {
	SSFF   int    `json:"ssff"`
	Nombre string `json:"nombre"`
}

type ChecaResponsable struct {
	SSFF   int    `json:"ssff"`
	Nombre string `json:"nombre"`
}

type ChecaArea struct {
	Id   int    `json:"id"`
	Area string `json:"area"`
}
