package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/stsh89/web-math-go/equations"

	"github.com/gin-gonic/gin"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"List all equations": "GET /equations",
			"Create equation":    "POST /equations",
			"Find equation":      "GET /equations/:id",
			"Delete equation":    "DELETE /equations/:id",
		})
	})

	r.GET("/equations", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Equations": equations.ListEquations(logger),
		})
	})

	r.Run()
}
