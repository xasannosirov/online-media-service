// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	"context"

	"github.com/xasannosirov/online-media-service/internal/entity"
)

type (
	File interface {
		Store(ctx context.Context, file entity.File) (entity.File, error)
		Remove(ctx context.Context, url string) error
	}

	FileRepo interface {
		Store(ctx context.Context, file entity.File) (entity.File, error)
		Remove(ctx context.Context, url string) error
	}
)
