BUILD_FLAGS = CGO_ENABLED=0 GOOS=linux GOARCH=amd64

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: all
all: build docker clean

.PHONY: build
build:
	$(BUILD_FLAGS) go build -o app

.PHONY: docker
docker:
	docker build -t say .

.PHONY: clean
clean:
	rm -rf app