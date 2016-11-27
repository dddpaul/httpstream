.PHONY: all build release

IMAGE=dddpaul/httpstream
VERSION=$(shell cat VERSION)

all: build

build:
	@go test
	@mkdir -p root/bin
	@CGO_ENABLED=0 go build -o root/bin/httpstream
	@docker build --tag=${IMAGE} .

debug:
	@docker run -it --entrypoint=sh ${IMAGE}

release: build
	@docker build --tag=${IMAGE}:${VERSION} .

deploy: release
	@docker push ${IMAGE}
	@docker push ${IMAGE}:${VERSION}
