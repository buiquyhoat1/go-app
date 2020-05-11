package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go_api/form/api"
	"go_api/models"
	"go_api/models/repository"
)

// Controller is blog controlller
type BlogController struct{}

// Index action: GET /blogs
func (pc BlogController) Index(c *gin.Context) {
	var u repository.BlogRepository
	b, err := u.GetAll()
	if err != nil {
		c.AbortWithStatus(404)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, b)
	}
}

// Create action: POST /blogs
func (pc BlogController) Create(c *gin.Context) {
	var u repository.BlogRepository
	in := api.Blog{}
	if err := c.BindJSON(&in); err != nil {
		return
	}

	statusInt, _ := strconv.Atoi(in.Status)

	in2 := &models.Blog{
		Title: 	 in.Title,
		Author: in.Author,
		Content: in.Content,
		Category:  in.Category,
		Status: statusInt,
	}
	b, err := u.CreateModel(in2)
	if err != nil {
		c.AbortWithStatus(400)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(201, b)
	}
}

// Show action: Get /blogs/:id
func (pc BlogController) Show(c *gin.Context) {
	id := c.Params.ByName("id")
	var b repository.BlogRepository
	idInt, _ := strconv.Atoi(id)
	blog, err := b.GetByID(idInt)

	if err != nil {
		c.AbortWithStatus(400)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, blog)
	}
}

// Update action: Put /blogs/:id
func (pc BlogController) Update(c *gin.Context) {
	id := c.Params.ByName("id")
	var b repository.BlogRepository
	idInt, _ := strconv.Atoi(id)
	blog, err := b.UpdateByID(idInt, c)

	if err != nil {
		c.AbortWithStatus(404)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, blog)
	}
}

// Delete action: DELETE /blogs/:id
func (pc BlogController) Delete(c *gin.Context) {
	id := c.Params.ByName("id")
	var b repository.BlogRepository
	idInt, _ := strconv.Atoi(id)
	if err := b.DeleteByID(idInt); err != nil {
		c.AbortWithStatus(403)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"success": "ID" + id + "の投稿を削除しました"})
	return
}