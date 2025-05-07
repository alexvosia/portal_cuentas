package entities

type Module struct {
	Id            int    `json:"id"`
	Nombre        string `json:"nombre"`
	Descripcion   string `json:"descripcion"`
	Estado        int    `json:"estado"`
	Alcance       string `json:"alcance"`
	Coordinador   string `json:"coordinador"`
	Respoonsable1 string `json:"responsable1"`
	Respoonsable2 string `json:"responsable2"`
	Script        string `json:"script"`
	Mail          string `json:"mail"`
	Columna1      string `json:"columna1"`
	Columna2      string `json:"columna2"`
	Columna3      string `json:"columna3"`
	Columna4      string `json:"columna4"`
	Columna5      string `json:"columna5"`
	Columna6      string `json:"columna6"`
	Columna7      string `json:"columna7"`
	Columna8      string `json:"columna8"`
	Columna9      string `json:"columna9"`
	Inicio        string `json:"inicio"`
	Fin           string `json:"fin"`
	Creador       int    `json:"creador"`
	Upper         int    `json:"upper"`
}
