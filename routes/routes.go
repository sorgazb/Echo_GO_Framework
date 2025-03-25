package routes

import (
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
	"net/http"
)

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func SetupRoutes(e *echo.Echo) {

	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}

	e.Renderer = renderer

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index-old.html", map[string]string{
			"Title":   "Mi aplicacion Echo",
			"Heading": "!Hola Mundo!",
			"Message": "Bienvenido a mi aplicacion web con Echo y plantillas HTML.",
		})
	})

	e.GET("/saludo", func(c echo.Context) error {
		return c.String(http.StatusOK, "Saludos desde Echo")
	})

	e.GET("/saludo/:nombre", func(c echo.Context) error {
		nombre := c.Param("nombre")
		return c.String(http.StatusOK, "Saludos "+nombre+"!")
	})

	e.Static("/static", "static")
}
