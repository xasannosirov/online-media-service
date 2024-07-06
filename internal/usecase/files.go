package usecase

import (
	"context"
	"fmt"

	"github.com/xasannosirov/online-media-service/internal/entity"
)

// FileUseCase -.
type FileUseCase struct {
	repo FileRepo
}

// New -.
func New(r FileRepo) *FileUseCase {
	return &FileUseCase{
		repo: r,
	}
}

// Store - getting translate history from store.
func (uc *FileUseCase) Store(ctx context.Context, file entity.File) (entity.File, error) {
	files, err := uc.repo.Store(ctx, file)
	if err != nil {
		return entity.File{}, fmt.Errorf("FileUseCase - Store - s.repo.Store: %w", err)
	}

	return files, nil
}

// Remove -.
func (uc *FileUseCase) Remove(ctx context.Context, url string) error {
	err := uc.repo.Remove(ctx, url)
	if err != nil {
		return fmt.Errorf("FileUseCase - Remove - s.repo.Remove: %w", err)
	}

	return nil
}
