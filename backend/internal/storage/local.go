package storage

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"os"
)

type localStorage struct {
	basePath string
	baseURL  string
}

// Save implements [ImageStorage].
func (l *localStorage) Save(ctx context.Context, file multipart.File, filename string) (string, error) {
	err := os.MkdirAll(l.basePath, os.ModePerm)
	if err != nil {
		return "", fmt.Errorf("error creating folder: %w", err)
	}

	dst, err := os.Create(fmt.Sprintf("%s/%s", l.basePath, filename))
	if err != nil {
		return "", fmt.Errorf("error creating file: %w", err)
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		return "", fmt.Errorf("error copying file: %w", err)
	}

	url := fmt.Sprintf("%s/uploads/%s", l.baseURL, filename)
	return url, nil
}

func NewLocalStorage(basePath, baseURL string) ImageStorage {
	return &localStorage{basePath: basePath, baseURL: baseURL}
}
