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
	CreateModulo(ctx context.Context, name string, descripcion string, coordinador int, areasid []entities.Area, script string, mail string) (*entities.Module, error)
	GetModulo(ctx context.Context, id int) (*entities.Module, error) // obtener nombre de coordinador y responsables
	GetModulos(ctx context.Context) ([]entities.Module, error)
	SetStatusModulo(ctx context.Context, id int, status int) (*entities.Module, error)
	SetCoordinador(ctx context.Context, id int, coordinadorid int) (*entities.Module, error) //apunta a coordinador_analista
	SetAreas(ctx context.Context, id int, areasid []entities.Area) (*entities.Module, error)
	SetScript(ctx context.Context, id int, script string) (*entities.Module, error)
	SetMail(ctx context.Context, id int, mail string) (*entities.Module, error)
	SetDescripcion(ctx context.Context, id int, descripcion string) (*entities.Module, error)

	/* SetLayOut(ctx context.Context, id int, layoutid int) (*entities.Module, error)
	SetResponsable(ctx context.Context, id int, responsableid int, index int) (*entities.Module, error) */
}

type FileAPI interface {
	CreateFile(ctx context.Context, nombre string, idModulo int, tipoRegistro string, path string, size int64, creador int) (*entities.FileCSV, error)
	GetFile(ctx context.Context, id int) (*entities.FileCSV, error)
	GetFiles(ctx context.Context) ([]entities.FileCSV, error)

	SetFin(ctx context.Context, id int, fin string, upper int) (*entities.FileCSV, error)
}

type ResponsableAPI interface {
	CreateResponsable(ctx context.Context, idModulo int, idCoordinador int, idResponsable1 int, idResponsable2 int, creador int) (*entities.Responsables, error)
	GetResponsableById(ctx context.Context, id int) (*entities.Responsables, error)
	GetResponsableByCoordinador(ctx context.Context, idCoordinador int) (*entities.Responsables, error)
}
