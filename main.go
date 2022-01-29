package main

import (
	"log"
	"net/http"

	"github.com/afman42/go-crud-reactjs/handler"
	"github.com/afman42/go-crud-reactjs/post"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:afif123@tcp(127.0.0.1:3306)/go-crud?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	// db.AutoMigrate(&post.Post{})
	postRepository := post.NewRepository(db)
	postService := post.NewService(postRepository)
	postHandler := handler.NewPostHandler(postService)

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Use(cors.Default())
	r.Static("/public", "./public")
	api := r.Group("/api")

	api.GET("/post", postHandler.GetPosts)
	api.GET("/post/:id", postHandler.GetPostByID)
	api.POST("/post", postHandler.CreatePost)
	api.PUT("/post/:id", postHandler.UpdatePost)
	api.DELETE("/post/:id", postHandler.DeletePostByID)

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080"):
}
