package main

import (
	"context"

	"github.com/beego/beego/v2/server/web"
	beecontext "github.com/beego/beego/v2/server/web/context"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	semconv "go.opentelemetry.io/otel/semconv/v1.43.0"
	"go.opentelemetry.io/otel/trace"

	"github.com/uptrace/opentelemetry-go-extra/otelplay"
)

func main() {
	ctx := context.Background()

	shutdown := otelplay.ConfigureOpentelemetry(ctx)
	defer shutdown()

	// otelhttp only sees the route via Request.Pattern, which Beego does not set.
	web.InsertFilter("*", web.BeforeExec, otelRouteFilter)

	web.Router("/", &IndexController{})
	web.Router("/hello/:username", &HelloController{})

	web.RunWithMiddleWares("localhost:9999", otelhttp.NewMiddleware("service-name"))
}

// otelRouteFilter copies the matched Beego route to the otelhttp span and metrics.
func otelRouteFilter(ctx *beecontext.Context) {
	route, ok := ctx.Input.GetData("RouterPattern").(string)
	if !ok || route == "" {
		return
	}

	reqCtx := ctx.Request.Context()

	span := trace.SpanFromContext(reqCtx)
	span.SetName(ctx.Request.Method + " " + route)
	span.SetAttributes(semconv.HTTPRoute(route))

	if labeler, ok := otelhttp.LabelerFromContext(reqCtx); ok {
		labeler.Add(semconv.HTTPRoute(route))
	}
}

type IndexController struct {
	web.Controller
}

func (c *IndexController) Get() {
	ctx := c.Ctx.Request.Context()

	c.Data["traceURL"] = otelplay.TraceURL(trace.SpanFromContext(ctx))
	c.TplName = "index.tpl"
}

type HelloController struct {
	web.Controller
}

func (c *HelloController) Get() {
	ctx := c.Ctx.Request.Context()

	c.Data["username"] = c.Ctx.Input.Param(":username")
	c.Data["traceURL"] = otelplay.TraceURL(trace.SpanFromContext(ctx))
	c.TplName = "hello.tpl"
}
