module github.com/uptrace/opentelemetry-go-extra/example/mongo-driver

go 1.16

replace github.com/uptrace/opentelemetry-go-extra/otelplay => ../../otelplay

require (
	github.com/go-stack/stack v1.8.1 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/stretchr/objx v0.1.1 // indirect
	github.com/uptrace/opentelemetry-go-extra/otelplay v0.1.7
	github.com/youmark/pkcs8 v0.0.0-20201027041543-1326539a0a0a // indirect
	go.mongodb.org/mongo-driver v1.8.1
	go.opentelemetry.io/contrib/instrumentation/go.mongodb.org/mongo-driver/mongo/otelmongo v0.28.0
	go.opentelemetry.io/otel v1.3.0
	golang.org/x/crypto v0.0.0-20211209193657-4570a0811e8b // indirect
)
