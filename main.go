package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/stsh89/web-math-go/app"
	"github.com/stsh89/web-math-go/equations"
	"github.com/stsh89/web-math-go/providers/notion"

	"github.com/gin-gonic/gin"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	config := buildConfig()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"List all equations": "GET /equations",
			"Create equation":    "POST /equations",
			"Delete equation":    "DELETE /equations",
			"Get points for given equation on a range from -5 to 5": "GET /calculate",
		})
	})

	r.GET("/equations", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Equations": equations.ListEquations(logger, &config),
		})
	})

	r.POST("/equations", func(c *gin.Context) {
		var json struct {
			Term string `json:"term" binding:"required"`
		}

		c.Bind(&json)

		c.JSON(http.StatusOK, gin.H{
			"ID": equations.CreateEquation(json.Term, logger, &config),
		})
	})

	r.DELETE("/equations", func(c *gin.Context) {
		var json struct {
			Term string `json:"term" binding:"required"`
		}

		c.Bind(&json)

		c.JSON(http.StatusOK, gin.H{
			"ID": equations.DeleteEquation(json.Term, logger, &config),
		})
	})

	r.POST("/calculate", func(c *gin.Context) {
		var json struct {
			Term string             `json:"term" binding:"required"`
			Args map[string]float64 `json:"args" binding:"required"`
		}

		c.Bind(&json)

		points := equations.Calculate(json.Term, json.Args, logger, &config)

		c.JSON(http.StatusOK, gin.H{
			"points": points,
		})
	})

	r.Run()
}

func buildConfig() app.Config {
	return app.Config{
		NotionConfig: &notion.Configuration{
			ApiKey:     os.Getenv("NOTION_API_KEY"),
			DatabaseId: os.Getenv("NOTION_DATABASE_ID"),
		},
	}
}
