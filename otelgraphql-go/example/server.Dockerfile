FROM golang:alpine AS base
COPY . /src/
WORKDIR /src/instrumentation/github.com/graph-gophers/graphql-go/otelgraphql-go/example

FROM base AS graphql-go-server
RUN go install ./server.go
CMD ["/go/bin/server"]
