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
	CreateModulo(ctx context.Context, module entities.Module) (*entities.Module, error)
	GetModuloById(ctx context.Context, idModulo int) (*entities.Module, error) // Obtener nombre de coordinador y responsables
	GetModulos(ctx context.Context, role string) ([]entities.Module, error)
	SetStatusModulo(ctx context.Context, idModulo int, status int, upper int) (*entities.Module, error)
	SetCoordinador(ctx context.Context, idModulo int, idCoordinador int, upper int) (*entities.Module, error)
	SetResponsable(ctx context.Context, idModulo int, idCoordinador int, idResponsable1 int, idResponsable2 int, upper int) (*entities.Module, error)
	SetAreas(ctx context.Context, idModulo int, areas string, upper int) (*entities.Module, error)
	SetScript(ctx context.Context, idModulo int, script string, upper int) (*entities.Module, error)
	SetMail(ctx context.Context, idModulo int, mail string, upper int) (*entities.Module, error)
	SetDescripcion(ctx context.Context, idModulo int, descripcion string, upper int) (*entities.Module, error)
}

type ChecaDataAPI interface {
	GetCoordinadores(ctx context.Context) ([]entities.ChecaCoordinador, error)
	GetResponsables(ctx context.Context, idCoordinador int) ([]entities.ChecaResponsable, error)
	GetAreas(ctx context.Context) ([]entities.ChecaArea, error)
}

type FileAPI interface {
	CreateFile(ctx context.Context, nombre string, idModulo int, tipoRegistro string, creador int) error
	GetFilesByModulo(ctx context.Context, idModulo int) ([]entities.FileCSV, error)
	SetFin(ctx context.Context, idFile int, upper int) (*entities.FileCSV, error)
}

type ResponsableAPI interface {
	GetResponsablesByCoordinador(ctx context.Context, idCoordinador int) (*entities.Responsables, error)
}
