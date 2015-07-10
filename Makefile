CGO_ENABLED=0
GOOS=linux
GOARCH=amd64
TAG=${TAG:-latest}
NAME=redirector
REPO=ehazlett/$(NAME)

all: build

clean:
	@rm $(NAME)

build:
	@go build -a -tags 'netgo' -ldflags '-w -linkmode external -extldflags -static' .

image: build
	@echo Building $(NAME) image $(TAG)
	@docker build -t $(REPO):$(TAG) .

release: deps build image
	@docker push $(REPO):$(TAG)

test:
	@go test -v ./...

.PHONY: all deps build clean image test release
