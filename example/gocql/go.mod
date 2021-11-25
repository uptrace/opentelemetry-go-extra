module github.com/uptrace/opentelemetry-go-extra/example/gocql

go 1.16

replace github.com/uptrace/opentelemetry-go-extra/otelplay => ../../otelplay

require (
	github.com/gocql/gocql v0.0.0-20211015133455-b225f9b53fa1
	github.com/golang/snappy v0.0.4 // indirect
	github.com/uptrace/opentelemetry-go-extra/otelplay v0.1.6
	go.opentelemetry.io/contrib/instrumentation/github.com/gocql/gocql/otelgocql v0.27.0
	go.opentelemetry.io/otel v1.2.0
	go.opentelemetry.io/otel/trace v1.2.0
	golang.org/x/net v0.0.0-20211123203042-d83791d6bcd9 // indirect
	golang.org/x/sys v0.0.0-20211124211545-fe61309f8881 // indirect
	google.golang.org/genproto v0.0.0-20211118181313-81c1377c94b1 // indirect
)
