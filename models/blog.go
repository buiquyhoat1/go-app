package models

import (
    // "time"
    "github.com/jinzhu/gorm"
)

// Blog is blog models property
type Blog struct {
    gorm.Model
    // ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
    Title     string    `gorm:"size:255;not null;" json:"title"`
    Content     string    `gorm:"type:text" json:"content"`
    Author    string    `gorm:"size:255;not null;" json:"author"`
    Category  string    `gorm:"size:255;not null;" json:"category"`
    // CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
    // UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
    Status    int       `gorm:"default:0"`
}