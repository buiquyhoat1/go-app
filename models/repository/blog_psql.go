package repository

import (
	"fmt"

	// "strconv"
	"github.com/gin-gonic/gin"
	"go_api/db"
	"go_api/form/api"
	"go_api/models"
)

// Service procides blog's behavior
type BlogRepository struct{}

// Blog is alias of entity.Blog struct
type Blog api.Blog

// Callback for after create

// func (_ BlogRepository) AfterCreate(b *models.Blog) (err error) {
// 	db := db.GetDB()
// 	db.Model(&b).Update("title", "hoat")
// 	return
// }

// GetAll is get all Blog
func (_ BlogRepository) GetAll() ([]Blog, error) {
	db := db.GetDB()
	var b []Blog

	if err := db.Order("updated_at desc").Find(&b).Error; err != nil {
		return nil, err
	}

	return b, nil
}

// CreateModel is create Blog model
func (_ BlogRepository) CreateModel(b *models.Blog) (*models.Blog, error) {
	fmt.Printf("%+v", b)
	db := db.GetDB()
	if err := db.Create(&b).Error; err != nil {
		return b, err
	}
	return b, nil
}

// GetByID is get a Blog by ID
func (_ BlogRepository) GetByID(id int) (api.Blog, error) {
	db := db.GetDB()
	var b api.Blog
	if err := db.Where("id = ?", id).First(&b).Error; err != nil {
		return b, err
	}
	return b, nil
}

// UpdateByID is update a Blog
func (_ BlogRepository) UpdateByID(id int, c *gin.Context) (api.Blog, error) {
	db := db.GetDB()
	var b api.Blog
	if err := db.Where("id = ?", id).First(&b).Error; err != nil {
		return b, err
	}
	
	if err := c.BindJSON(&b); err != nil {
		return b, err
	}
	fmt.Printf("%+V", b)
	// statusInt, _ := strconv.Atoi(b.Status)
	db.Save(&b)

	return b, nil
}

// DeleteByID is delete a Blog by ID
func (_ BlogRepository) DeleteByID(id int) error {
	db := db.GetDB()
	// var b Blog

	if err := db.Where("id = ?", id).Delete(&models.Blog{}).Error; err != nil {
		return err
	}

	return nil
}
