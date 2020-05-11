package server

import (
    "github.com/gin-gonic/gin"
    "go_api/controllers"
)

// Init is initialize server
func Init() {
    r := router()
    r.Run()
}

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func router() *gin.Engine {
    r := gin.Default()
    r.Use(CORS())
    u := r.Group("/users")
    {
        ctrl := controllers.UserController{}
        u.GET("", ctrl.Index)
        u.POST("", ctrl.Create)
        u.GET("/:id", ctrl.Show)
        u.PUT("/:id", ctrl.Update)
        u.DELETE("/:id", ctrl.Delete)
    }

    p := r.Group("/posts")
    {
        ctrl := controllers.PostController{}
        p.GET("", ctrl.Index)
        p.POST("", ctrl.Create)
        p.GET("/:id", ctrl.Show)
        p.PUT("/:id", ctrl.Update)
        p.DELETE("/:id", ctrl.Delete)
    }

    b := r.Group("/blogs")
    {
        ctrl := controllers.BlogController{}
        b.GET("", ctrl.Index)
        b.POST("", ctrl.Create)
        b.GET("/:id", ctrl.Show)
        b.PUT("/:id", ctrl.Update)
        b.DELETE("/:id", ctrl.Delete)
    }

    return r
}