package main

import (
	"net/http"
	"io"
	"github.com/labstack/echo/v4"
	"html/template"
)

type Template struct {
    templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()
	renderer := &Template{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}

	e.Renderer = renderer
	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", nil)
	})
	e.Logger.Fatal(e.Start("0.0.0.0:8000"))
}