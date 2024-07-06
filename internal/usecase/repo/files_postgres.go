package repo

import (
	"context"
	"fmt"

	"github.com/xasannosirov/online-media-service/internal/entity"
	"github.com/xasannosirov/online-media-service/pkg/postgres"
)

const _defaultEntityCap = 64

// FileRepo -.
type FileRepo struct {
	*postgres.Postgres
}

// New -.
func New(pg *postgres.Postgres) *FileRepo {
	return &FileRepo{pg}
}

// Store -.
func (r *FileRepo) Store(ctx context.Context, f entity.File) (entity.File, error) {
	sql, args, err := r.Builder.
		Insert("files").
		Columns("file_name, file_url").
		Values(f.Filename, f.FileURL).
		ToSql()
	if err != nil {
		return entity.File{}, fmt.Errorf("FileRepo - Store - r.Builder: %w", err)
	}

	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return entity.File{}, fmt.Errorf("FileRepo - Store - r.Pool.Exec: %w", err)
	}

	return f, nil
}

// Remove -.
func (r *FileRepo) Remove(ctx context.Context, url string) error {
	sql, args, err := r.Builder.
		Delete("files").
		Where("file_url = ?", url).
		ToSql()
	if err != nil {
		return fmt.Errorf("FileRepo - Remove - r.Builder: %w", err)
	}

	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("FileRepo - Remove - r.Pool.Exec: %w", err)
	}

	return nil
}
