module github.com/uptrace/opentelemetry-go-extra/otelgraphql/example

go 1.15

replace github.com/uptrace/opentelemetry-go-extra/otelgraphql => ../

replace github.com/uptrace/opentelemetry-go-extra/otelutil => ../../otelutil

replace github.com/uptrace/opentelemetry-go-extra/otelplay => ../../otelplay

require (
	github.com/google/uuid v1.3.0
	github.com/graph-gophers/graphql-go v1.2.0
	github.com/uptrace/opentelemetry-go-extra/otelgraphql v0.1.6
	github.com/uptrace/opentelemetry-go-extra/otelplay v0.1.6
	golang.org/x/net v0.0.0-20211123203042-d83791d6bcd9 // indirect
	golang.org/x/sys v0.0.0-20211124211545-fe61309f8881 // indirect
	google.golang.org/genproto v0.0.0-20211118181313-81c1377c94b1 // indirect
)
