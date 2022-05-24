.PHONY: install build-win build build-docker run-docker

install:
	@go get ./...

build-win: install
	@go build -o PortDomainService.exe ./cmd/service/
	
build: install
	@go build -o PortDomainService ./cmd/service/

build-docker:
	@docker build --rm -t port-domain-service .

run-docker: build-docker
	@docker run -it --rm port-domain-service