
# Image URL to use all building/pushing image targets
IMG ?= datahub:latest

.PHONY: all test datahub license-utils
all: test datahub license-utils

# Run tests
test: generate fmt vet
	go test ./pkg/... ./cmd/... -coverprofile cover.out

# Build license-utils binary
license-utils:
	cd tools/license-utils; \
	make

# Build datahub binary
datahub: generate fmt vet
	go build -mod=vendor -ldflags "-X main.VERSION=`git rev-parse --abbrev-ref HEAD`-`git rev-parse --short HEAD``git diff --quiet || echo '-dirty'` -X 'main.BUILD_TIME=`date`' -X 'main.GO_VERSION=`go version`'" -o bin/datahub github.com/containers-ai/alameda/datahub/cmd

.PHONY: run

# Run against the configured Kubernetes cluster in ~/.kube/config
run: generate fmt vet
	go run ./cmd/main.go run

.PHONY: fmt vet generate docker-build docker-push

# Run go fmt against code
fmt:
	go fmt ./pkg/... ./cmd/...

# Run go vet against code
vet:
	go vet ./pkg/... ./cmd/...

# Generate code
generate:
	go generate ./pkg/... ./cmd/...

# Build the docker image
docker-build: test
	docker build ./.. -t ${IMG} -f Dockerfile
