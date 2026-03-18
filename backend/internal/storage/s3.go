package storage

import (
	"context"
	"mime/multipart"
)

// to-do

type ImageStorage interface {
   Save(ctx context.Context, file multipart.File, filename string) (string, error)
}
