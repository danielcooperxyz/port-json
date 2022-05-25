# PortDomainService
## Pre-requisites
- Golang 1.18
- Make

## Install 
To compile an executable binary run the below command.
```
make build
```
### Windows
```
make build-win
```
### Docker
To compile a docker image
```
make build-docker
```
To compile and run a docker image
```
make run-docker
```

## Tests
You can run tests using the standard golang test framework. Convienience commands are also provided in the makefile.
```
make test
```