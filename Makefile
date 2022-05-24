.PHONY: build
build-win:
	@go build -o PortDomainService.exe ./cmd/service/
	
build:
	@go build -o PortDomainService ./cmd/service/

build-docker:
	@docker build --rm -t port-domain-service .

run-docker: build-docker
	@docker run -it --rm port-domain-service