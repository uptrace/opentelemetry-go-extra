FROM golang:alpine AS base
COPY . /src/
WORKDIR /src/otelgraphql/example

FROM base AS graphql-go-server
RUN go install ./server.go
CMD ["/go/bin/server"]
