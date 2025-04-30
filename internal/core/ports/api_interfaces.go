package ports

import (
	"context"
	"errors"
	"infctas/internal/core/entities"
)

type MouduleAPI interface {
	GetModulo(ctx context.Context, id int) (*entities.Module, error)
	GetModulos(ctx context.Context) ([]entities.Module, error)
	CreateModulo(ctx context.Context, name string, descripcion string, coordinador int, areasid []entities.Area, script string, mail string) (*entities.Module, error)
	SetStatusModulo(ctx context.Context, id int, status int) (*entities.Module, error)
	SetCoordinador(ctx context.Context, id int, coordinadorid int) (*entities.Module, error)
	SetAreas(ctx context.Context, id int, areasid []entities.Area) (*entities.Module, error)
	SetScript(ctx context.Context, id int, script string) (*entities.Module, error)
	SetMail(ctx context.Context, id int, mail string) (*entities.Module, error)
	SetDescripcion(ctx context.Context, id int, descripcion string) (*entities.Module, error)

	/* SetLayOut(ctx context.Context, id int, layoutid int) (*entities.Module, error)
	SetResponsable(ctx context.Context, id int, responsableid int, index int) (*entities.Module, error) */
}

// Errores
var (
	ErrModuleNotFound = errors.New("modulo no encontrado")
	ErrInvalidData    = errors.New("datos inválidos")
)
