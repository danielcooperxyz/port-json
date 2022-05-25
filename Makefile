.PHONY: install build-win build build-docker run-docker

install:
	@go get ./...

build-win: install
	@go build -o PortDomainService.exe ./cmd/service/

test: install
	@go test ./...
	
build: install
	@go build -o PortDomainService ./cmd/service/

build-docker:
	@docker build --rm -t port-domain-service .

run-docker-int: build-docker
	@docker run -it --rm --entrypoint=/bin/ash port-domain-service

run-docker: build-docker
	@docker run -it --rm port-domain-service