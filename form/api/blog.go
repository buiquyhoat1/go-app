package api

import "time"

type Blog struct {
    ID      uint   `json:"id"`
    Title string `json:"title" binding:"required"`
    Author string `json:"author" binding:"required"`
    Content string  `json:"content"`
    Category string `json:"category" binding:"required"`
    Status string `json:"status"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
    DeletedAt time.Time `json:"deleted_at"`
}