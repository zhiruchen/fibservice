all: deps linux_build

export GOPATH	  := $(CURDIR)/_project
export GOBIN 	  := $(CURDIR)/bin
DOCKER_IMAGE_TAG  ?= $(shell git describe --abbrev=0 --tags)
DOCKER_IMAGE_NAME ?= fibservice
COMMONENVVAR      ?= GOOS=linux GOARCH=amd64
BUILDENVVAR       ?= CGO_ENABLED=0

CURRENT_GIT_GROUP := github.com/zhiruchen
CURRENT_GIT_REPO := fibservice

folder_dep:
	mkdir -p $(CURDIR)/_project/src/$(CURRENT_GIT_GROUP)
	test -d $(CURDIR)/_project/src/$(CURRENT_GIT_GROUP)/$(CURRENT_GIT_REPO) || ln -s $(CURDIR) $(CURDIR)/_project/src/$(CURRENT_GIT_GROUP)

deps: folder_dep
	mkdir -p $(CURDIR)/vendor
	glide install

build: deps
	go build -o bin/$(CURRENT_GIT_REPO) $(CURRENT_GIT_GROUP)/$(CURRENT_GIT_REPO)/main.go

linux_build:
	$(COMMONENVVAR) $(BUILDENVVAR) make build

test: deps
	go test -v $(CURRENT_GIT_GROUP)/$(CURRENT_GIT_REPO)/server

docker: linux_build
	@echo ">> build docker artifact"
	docker build --rm --no-cache -t $(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_TAG) $(CURDIR)

clean:
	@rm -rf vendor bin _project

.PHONY: deps install test add_dep clean 


