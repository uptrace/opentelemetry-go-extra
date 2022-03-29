PACKAGE_DIRS := $(shell find . -mindepth 2 -type f -name 'go.mod' -exec dirname {} \; | sort)

test:
	go test ./...
	go test ./... -short -race
	go test ./... -run=NONE -bench=. -benchmem
	env GOOS=linux GOARCH=386 go test ./...
	go vet ./...

fmt:
	gofmt -w -s ./
	gofumports -w  -local github.com/uptrace/opentelemetry-go-extra ./

go_mod_tidy:
	set -e; for dir in $(PACKAGE_DIRS); do \
	  echo "go mod tidy in $${dir}"; \
	  (cd "$${dir}" && go get -u && go mod tidy -go=1.17); \
	done
