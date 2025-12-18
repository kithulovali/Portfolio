package main

import (
	"Portfolio/internal/database"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()
	g := gin.Default()

	g.LoadHTMLGlob("templates/*")

	g.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			"title": "Home",
		})
	})

	g.GET("/projects", func(c *gin.Context) {
		c.HTML(200, "projects.html", gin.H{
			"title": "Projects",
		})
	})

	g.GET("/api/projects", func(c *gin.Context) {
		projects := []map[string]string{
			{"title": "goffart", "description": "description1"},
			{"title": "jean marc", "description": "description2"},
		}
		c.JSON(200, gin.H{
			"projects": projects,
		})

	})
	g.Run(":8000")
}
