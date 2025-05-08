package entities

type RegistryFileCSV struct {
	Id           int    `json:"id"`
	Nombre       string `json:"nombre"`
	IdModulo     int    `json:"id_modulo"`
	TipoRegistro string `json:"tipo_registro"`
	Path         string `json:"path"`
	Size         int64  `json:"size"`
	Inicio       string `json:"inicio"`
	Fin          string `json:"fin"`
	Creador      int    `json:"creador"`
	Upper        int    `json:"upper"`
}
