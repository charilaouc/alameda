
# Image URL to use all building/pushing image targets
IMG ?= evictioner:latest

.PHONY: all test evictioner
all: test evictioner

# Run tests
test: generate fmt vet
	go test ./pkg/... ./cmd/... -coverprofile cover.out

# Build evictioner binary
evictioner: generate fmt vet
	go build -ldflags "-X main.VERSION=`git rev-parse --abbrev-ref HEAD`-`git rev-parse --short HEAD``git diff --quiet || echo '-dirty'` -X 'main.BUILD_TIME=`date`' -X 'main.GO_VERSION=`go version`'" -o bin/evictioner github.com/containers-ai/alameda/evictioner/cmd

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
