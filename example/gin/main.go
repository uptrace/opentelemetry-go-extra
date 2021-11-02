package main

import (
	"context"
	"html/template"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel/trace"

	"github.com/uptrace/opentelemetry-go-extra/otelplay"
)

const (
	indexTmpl   = "index"
	profileTmpl = "profile"
)

func main() {
	ctx := context.Background()

	shutdown := otelplay.ConfigureOpentelemetry(ctx)
	defer shutdown()

	router := gin.Default()
	router.SetHTMLTemplate(parseTemplates())
	router.Use(otelgin.Middleware("service-name"))
	router.GET("/", indexHandler)
	router.GET("/hello/:username", helloHandler)

	if err := router.Run("localhost:9999"); err != nil {
		log.Print(err)
	}
}

func parseTemplates() *template.Template {
	indexTemplate := `
		<html>
		<p>Here are some routes for you:</p>
		<ul>
			<li><a href="/hello/world">Hello world</a></li>
			<li><a href="/hello/foo-bar">Hello foo-bar</a></li>
		</ul>
		<p><a href="{{ .traceURL }}" target="_blank">{{ .traceURL }}</a></p>
		</html>
	`
	t := template.Must(template.New(indexTmpl).Parse(indexTemplate))

	profileTemplate := `
		<html>
		<h3>Hello {{ .username }}</h3>
		<p><a href="{{ .traceURL }}" target="_blank">{{ .traceURL }}</a></p>
		</html>
	`
	return template.Must(t.New(profileTmpl).Parse(profileTemplate))
}

func indexHandler(c *gin.Context) {
	ctx := c.Request.Context()
	otelgin.HTML(c, http.StatusOK, indexTmpl, gin.H{
		"traceURL": otelplay.TraceURL(trace.SpanFromContext(ctx)),
	})
}

func helloHandler(c *gin.Context) {
	ctx := c.Request.Context()
	otelgin.HTML(c, http.StatusOK, profileTmpl, gin.H{
		"username": c.Param("username"),
		"traceURL": otelplay.TraceURL(trace.SpanFromContext(ctx)),
	})
}
