package main

import (
	"context"

	"github.com/astaxie/beego"
	"go.opentelemetry.io/contrib/instrumentation/github.com/astaxie/beego/otelbeego"
	"go.opentelemetry.io/otel/trace"

	"github.com/uptrace/opentelemetry-go-extra/otelplay"
)

func main() {
	ctx := context.Background()

	shutdown := otelplay.ConfigureOpentelemetry(ctx)
	defer shutdown()

	// To enable tracing on template rendering, disable autorender and
	// call otelbeego.Render manually.
	beego.BConfig.WebConfig.AutoRender = false

	beego.Router("/", &IndexController{})
	beego.Router("/hello/:username", &HelloController{})

	mware := otelbeego.NewOTelBeegoMiddleWare("service-name")
	beego.RunWithMiddleWares("localhost:9999", mware)
}

type IndexController struct {
	beego.Controller
}

func (c *IndexController) Get() {
	ctx := c.Ctx.Request.Context()

	c.Data["traceURL"] = otelplay.TraceURL(trace.SpanFromContext(ctx))
	c.TplName = "index.tpl"

	if err := otelbeego.Render(&c.Controller); err != nil {
		c.Abort("500")
	}
}

type HelloController struct {
	beego.Controller
}

func (c *HelloController) Get() {
	ctx := c.Ctx.Request.Context()

	c.Data["username"] = c.Ctx.Input.Param(":username")
	c.Data["traceURL"] = otelplay.TraceURL(trace.SpanFromContext(ctx))
	c.TplName = "hello.tpl"

	if err := otelbeego.Render(&c.Controller); err != nil {
		c.Abort("500")
	}
}
