BINPATH ?= build

BUILD_TIME=$(shell date +%s)
GIT_COMMIT=$(shell git rev-parse HEAD)
VERSION ?= $(shell git tag --points-at HEAD | grep ^v | head -n 1)

LDFLAGS = -ldflags "-X main.BuildTime=$(BUILD_TIME) -X main.GitCommit=$(GIT_COMMIT) -X main.Version=$(VERSION)"
LOCAL_DP_RENDERER_IN_USE = $(shell grep -c "\"github.com/ONSdigital/dp-renderer\" =" go.mod)

.PHONY: all
all: audit test build

.PHONY: audit
audit:
	go list -json -m all | nancy sleuth

.PHONY: lint
lint:
	exit

.PHONY: generate-debug
generate-debug: fetch-renderer-lib
 cd assets; go run github.com/kevinburke/go-bindata/go-bindata -prefix $(CORE_ASSETS_PATH)/assets -debug -o data.go -pkg assets locales/... templates/... $(CORE_ASSETS_PATH)/assets/locales/... $(CORE_ASSETS_PATH)/assets/templates/...
 { echo "// +build debug\n"; cat assets/data.go; } > assets/debug.go.new
 mv assets/debug.go.new assets/data.go

.PHONY: generate-prod
generate-prod: fetch-renderer-lib
 cd assets; go run github.com/kevinburke/go-bindata/go-bindata -prefix $(CORE_ASSETS_PATH)/assets -o data.go -pkg assets locales/... templates/... $(CORE_ASSETS_PATH)/assets/locales/... $(CORE_ASSETS_PATH)/assets/templates/...
 { echo "// +build production\n"; cat assets/data.go; } > assets/data.go.new
 mv assets/data.go.new assets/data.go

.PHONY: build
build: generate-prod

.PHONY: debug
debug: generate-debug

.PHONY: test
test:
	go test -race -cover ./...

.PHONY: convey
convey:
	goconvey ./...

.PHONY: test-component
test-component:
	go test -cover -coverpkg=github.com/ONSdigital/dp-frontend-area-profiles/... -component