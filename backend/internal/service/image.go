package service

import (
	"context"
	"fmt"
	"mime/multipart"
	"path/filepath"
	"time"

	"github.com/google/uuid"
	"github.com/mat-cf/image-host/internal/domain"
	"github.com/mat-cf/image-host/internal/repository"
	"github.com/mat-cf/image-host/internal/storage"
)

type ImageService interface {
	Upload(ctx context.Context, file multipart.File, header *multipart.FileHeader) (string, error)
}

type imageService struct {
	repo repository.ImageRepository
	storage storage.ImageStorage
}

// Upload implements [ImageService].
func (i *imageService) Upload(ctx context.Context, file multipart.File, header *multipart.FileHeader) (string, error) {
	id := uuid.NewString()

	// mount the internal filename
	ext := filepath.Ext(header.Filename)
	filename := fmt.Sprintf("%s%s", id, ext)

	url, err := i.storage.Save(ctx, file, filename)
	if err != nil {
		return "", fmt.Errorf("error saving file: %w", err)
	}

	image := domain.Image{
		ID: id, 
		Filename: filename,
		URL: url,
		Size: header.Size,
		ContentType: header.Header.Get("Content-Type"),
		CreatedAt: time.Now(),
	}

	err = i.repo.Save(ctx, image)
	if err != nil {
		return "", fmt.Errorf("error saving metadata: %w", err)
	}

	return url, nil

}

func NewImageService(repo repository.ImageRepository, storage storage.ImageStorage) ImageService {
	return &imageService{repo: repo, storage: storage}
}
