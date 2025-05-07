package ports

import (
	"context"
	"errors"
	"infctas/internal/core/entities"
)

// Errores
var (
	ErrModuleNotFound = errors.New("modulo no encontrado")
	ErrInvalidData    = errors.New("datos inválidos")
)

type MouduleAPI interface {
	CreateModulo(ctx context.Context, nombre string, descripcion string, areas string, idCoordinador int, idResponsable1 int, idResponsable2 int, script string, mail string, columna1 string, columna2 string, columna3 string, columna4 string, columna5 string, columna6 string, columna7 string, columna8 string, columna9 string, creador int) (*entities.Module, error)
	GetModuloById(ctx context.Context, id int) (*entities.Module, error) // Obtener nombre de coordinador y responsables
	GetModulos(ctx context.Context, role string) ([]entities.Module, error)
	SetStatusModulo(ctx context.Context, id int, status int, upper int) (*entities.Module, error)
	SetAreas(ctx context.Context, id int, areas string, upper int) (*entities.Module, error)
	SetScript(ctx context.Context, id int, script string, upper int) (*entities.Module, error)
	SetMail(ctx context.Context, id int, mail string, upper int) (*entities.Module, error)
	SetDescripcion(ctx context.Context, id int, descripcion string, upper int) (*entities.Module, error)
}

type FileAPI interface {
	CreateFile(ctx context.Context, nombre string, idModulo int, tipoRegistro string, creador int) error
	GetFilesByModulo(ctx context.Context, idModulo int) ([]entities.FileCSV, error)
	SetFin(ctx context.Context, id int, upper int) (*entities.FileCSV, error)
}

type ResponsableAPI interface {
	SetResponsable(ctx context.Context, idModulo int, idCoordinador int, idResponsable1 int, idResponsable2 int, user int) (*entities.Responsables, error)
	GetResponsableById(ctx context.Context, id int) (*entities.Responsables, error)
	GetResponsableByCoordinador(ctx context.Context, idCoordinador int) (*entities.Responsables, error)
}

type ChecaDataAPI interface {
	GetCoordinadores(ctx context.Context) ([]entities.ChecaCoordinador, error)
	GetResponsables(ctx context.Context, idCoordinador int) ([]entities.ChecaResponsable, error)
	GetAreas(ctx context.Context) ([]entities.ChecaArea, error)
}
