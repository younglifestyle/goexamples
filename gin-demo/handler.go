package main

import (
	"text/template"

	"github.com/gin-gonic/gin"
)

func page404(c *gin.Context) {
	t, _ := template.ParseFiles("./static/error.html")
	t.Execute(c.Writer, c.Request)
}
