package domain

import "time"

type Image struct {
	ID string
	Filename string
	URL string
	Size int64
	ContentType string
	CreatedAt time.Time
}
