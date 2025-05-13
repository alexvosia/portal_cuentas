package entities

type RegistryFileCSV struct {
	Id           int
	Nombre       string
	IdModulo     int
	TipoRegistro string
	Path         string
	Size         int
	Inicio       string
	Fin          string
	Creador      int
	Upper        int
}
