package entities

type FileCSV struct {
	ID               int    `json:"id"`
	ModuleID         int    `json:"moduleid"`
	Name             string `json:"name"`
	Path             string `json:"path"`
	Size             int64  `json:"size"`
	FechaCreacion    string `json:"fechacreacion"`
	FechaEliminacion string `json:"fechaeliminacion"`
}
