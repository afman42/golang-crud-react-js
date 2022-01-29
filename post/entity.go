package post

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	ID          int
	Title       string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
