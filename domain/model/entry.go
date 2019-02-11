package model

import (
	"time"
)

type Entry struct {
	ID        int        `gorm:"primary_key" json:"id"`
	Title     string     `gorm:"not null" json:"title" binding:"required"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt *time.Time `json:"-"`
	// When you get a Entry Model with gorm, you can get associated tags using junction table automatically.
	// When you write or update a Entry Model, you have to read　tags id beforehand.
	Tags []*Tag `gorm:"many2many:entries_tags;association_autoupdate:false;association_autocreate:false;" json:"tags"`
}
