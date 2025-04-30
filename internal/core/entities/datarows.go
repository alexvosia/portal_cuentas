package entities

type DataRows struct {
	ID       int64  `json:"id" bd:"id"`
	Cuenta   string `json:"cuenta" bd:"cuenta"`
	Tipo     string `json:"tipo" bd:"tipo"`
	Oferta   string `json:"oferta" bd:"oferta"`
	Columna1 string `json:"columna1" bd:"columna1"`
	Columna2 string `json:"columna2" bd:"columna2"`
	Columna3 string `json:"columna3" bd:"columna3"`
	Columna4 string `json:"columna4" bd:"columna4"`
	Columna5 string `json:"columna5" bd:"columna5"`
	Columna6 string `json:"columna6" bd:"columna6"`
	Columna7 string `json:"columna7" bd:"columna7"`
	Columna8 string `json:"columna8" bd:"columna8"`
	Columna9 string `json:"columna9" bd:"columna9"`
	FileID   int64  `json:"file_id" bd:"file_id"`
}
