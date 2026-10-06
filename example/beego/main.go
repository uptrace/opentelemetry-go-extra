package main

import (
	"context"

	"github.com/beego/beego/v2/server/web"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel/trace"

	"github.com/uptrace/opentelemetry-go-extra/otelplay"
)

func main() {
	ctx := context.Background()

	shutdown := otelplay.ConfigureOpentelemetry(ctx)
	defer shutdown()

	web.Router("/", &IndexController{})
	web.Router("/hello/:username", &HelloController{})

	web.RunWithMiddleWares("localhost:9999", otelhttp.NewMiddleware("service-name"))
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
