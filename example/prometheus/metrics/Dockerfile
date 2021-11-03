FROM golang:alpine

COPY . /src/
WORKDIR /src
RUN go install /src/main.go

CMD ["/go/bin/main"]
