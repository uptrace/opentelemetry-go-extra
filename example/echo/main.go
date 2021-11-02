package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.opentelemetry.io/contrib/instrumentation/github.com/labstack/echo/otelecho"
	"go.opentelemetry.io/otel/trace"

	"github.com/uptrace/opentelemetry-go-extra/otelplay"
)

func main() {
	ctx := context.Background()

	shutdown := otelplay.ConfigureOpentelemetry(ctx)
	defer shutdown()

	e := echo.New()
	e.Use(otelecho.Middleware("service-name"))
	e.Use(middleware.Recover())
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		ctx := c.Request().Context()
		trace.SpanFromContext(ctx).RecordError(err)

		e.DefaultHTTPErrorHandler(err, c)
	}

	e.GET("/", indexHandler)
	e.GET("/hello/:username", helloHandler)

	e.Logger.Fatal(e.Start(":9999"))
}

func indexHandler(c echo.Context) error {
	ctx := c.Request().Context()

	traceURL := otelplay.TraceURL(trace.SpanFromContext(ctx))
	tmpl := `
	<html>
	<p>Here are some routes for you:</p>
	<ul>
		<li><a href="/hello/world">Hello world</a></li>
		<li><a href="/hello/foo-bar">Hello foo-bar</a></li>
	</ul>
	<p><a href="%s" target="_blank">%s</a></p>
	</html>
	`
	html := fmt.Sprintf(tmpl, traceURL, traceURL)
	return c.HTML(http.StatusOK, html)
}

func helloHandler(c echo.Context) error {
	ctx := c.Request().Context()

	traceURL := otelplay.TraceURL(trace.SpanFromContext(ctx))
	username := c.Param("username")
	tmpl := `
	<html>
	<h3>Hello %s</h3>
	<p><a href="%s" target="_blank">%s</a></p>
	</html>
	`
	html := fmt.Sprintf(tmpl, username, traceURL, traceURL)
	return c.HTML(http.StatusOK, html)
}
