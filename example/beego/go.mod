module github.com/uptrace/opentelemetry-go-extra/example/beego

go 1.16

replace github.com/uptrace/opentelemetry-go-extra/otelplay => ../../otelplay

require (
	github.com/astaxie/beego v1.12.3
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/prometheus/common v0.32.1 // indirect
	github.com/prometheus/procfs v0.7.3 // indirect
	github.com/shiena/ansicolor v0.0.0-20200904210342-c7312218db18 // indirect
	github.com/uptrace/opentelemetry-go-extra/otelplay v0.1.3
	go.opentelemetry.io/contrib/instrumentation/github.com/astaxie/beego/otelbeego v0.26.1
	go.opentelemetry.io/otel/trace v1.1.0
	golang.org/x/crypto v0.0.0-20210921155107-089bfa567519 // indirect
)
