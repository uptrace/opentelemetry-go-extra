module github.com/uptrace/opentelemetry-go-extra/example/beego

go 1.16

replace github.com/uptrace/opentelemetry-go-extra/otelplay => ../../otelplay

require (
	github.com/astaxie/beego v1.12.3
	github.com/uptrace/opentelemetry-go-extra/otelplay v0.0.0-00010101000000-000000000000
	go.opentelemetry.io/contrib/instrumentation/github.com/astaxie/beego/otelbeego v0.26.0
	go.opentelemetry.io/otel/trace v1.1.0
)
