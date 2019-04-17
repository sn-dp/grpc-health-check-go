.PHONY: all
all: build
APP_NAME=grpc-health-check-go
ifndef APP_NAME
$(error APP_NAME is not set)
endif

SRC_PKGS=$(shell go list ./... | grep -v "vendor")

deps:
	dep ensure

build: deps fmt vet

fmt:
	go fmt $(SRC_PKGS)

vet:
	go vet $(SRC_PKGS)

test:
	go test -v -count=1 $(SRC_PKGS)

proto-gen:
	protoc -I proto/ proto/health_check.proto --go_out=plugins=grpc:proto