FROM golang:alpine

RUN apk add --update alpine-sdk

ADD . /src/

WORKDIR /src/

RUN make build