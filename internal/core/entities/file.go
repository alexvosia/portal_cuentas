package entities

type FileCSV struct {
	Id           int    `json:"id" bd:"id"`
	Nombre       string `json:"nombre" bd:"nombre"`
	IdModulo     int    `json:"id_modulo" bd:"id_modulo"`
	TipoRegistro string `json:"tipo_registro" bd:"tipo_registro"`
	Path         string `json:"path" bd:"path"`
	Size         int64  `json:"size" bd:"size"`
	Inicio       string `json:"inicio" bd:"inicio"`
	Fin          string `json:"fin" bd:"fin"`
	Creador      int    `json:"creador" bd:"creador"`
	Upper        int    `json:"upper" bd:"upper"`
}
