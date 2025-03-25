package routes

import (
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
	"net/http"
	"os"
	"strings"
)

func Renderer() *TemplateRenderer {
	return &TemplateRenderer{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}
}

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func SetupRoutes(e *echo.Echo) {

	e.Renderer = Renderer()
	e.Static("/static", "static")

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", map[string]string{
			"Title":   "Mi aplicacion Echo",
			"Heading": "!Hola Mundo!",
			"Message": "Bienvenido a mi aplicacion web con Echo y plantillas HTML.",
		})
	})

	e.GET("/:page", func(c echo.Context) error {
		page := c.Param("page")
		if !strings.HasSuffix(page, ".html") {
			page += ".html"
		}

		if _, err := os.Stat("templates/" + page); err == nil {
			return c.Render(http.StatusOK, page, nil)
		}

		return c.Render(http.StatusNotFound, "404.html", nil)
	})

	e.GET("/saludo", func(c echo.Context) error {
		return c.String(http.StatusOK, "Saludos desde Echo")
	})

	e.GET("/saludo/:nombre", func(c echo.Context) error {
		nombre := c.Param("nombre")
		return c.String(http.StatusOK, "Saludos "+nombre+"!")
	})

}
