package ports

import "infctas/internal/core/entities"

type DuckDBRepo interface {
	SeachCuentaRows(cuenta int, idArea int) (entities.Responsables, error)
}
