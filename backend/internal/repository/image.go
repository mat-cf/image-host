package repository

import (
	"context"
	"database/sql"

	"github.com/mat-cf/image-host/internal/domain"
)

// All struct that implements those methods satisfacts the interface
type ImageRepository interface {
	Save(ctx context.Context, image domain.Image) error
	FindAll(ctx context.Context) ([]domain.Image, error)
	FindById(ctx context.Context, id string) (domain.Image, error)
}

type PostgresImageRepository struct {
	db *sql.DB
}

// FindAll implements [ImageRepository].
func (p *PostgresImageRepository) FindAll(ctx context.Context) ([]domain.Image, error) {
	query := `SELECT id, filename, url, size, content_type, created_at FROM images ORDER BY created_ad DESC`

	rows, err := p.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var images []domain.Image
	for rows.Next() {
		var img domain.Image
		err := rows.Scan(&img.ID, &img.Filename, &img.URL, &img.Size, &img.ContentType, &img.CreatedAt)
		if err != nil {
			return nil, err
		}
		images = append(images, img)
	}
	return images, nil
}

// FindById implements [ImageRepository].
func (p *PostgresImageRepository) FindById(ctx context.Context, id string) (domain.Image, error) {
	query := `SELECT id, filename, original_name, url, size, content_type, created_at FROM images WHERE id = $1`

	var img domain.Image
	err := p.db.QueryRowContext(ctx, query, id).Scan(&img.ID, &img.Filename, &img.URL, &img.Size, &img.ContentType, &img.CreatedAt)
	return img, err
}

// Save implements [ImageRepository].
func (p *PostgresImageRepository) Save(ctx context.Context, image domain.Image) error {
	query := `
        INSERT INTO images (id, filename, url, size, content_type, created_at)
        VALUES ($1, $2, $3, $4, $5, $6)
    `
    _, err := p.db.ExecContext(ctx, query,
        image.ID,
        image.Filename,
        image.URL,
        image.Size,
        image.ContentType,
        image.CreatedAt,
    )
    return err
}

// constructor
func NewPostgresImageRepository(db *sql.DB) ImageRepository {
	return &PostgresImageRepository{db: db}
}
