FROM golang:alpine AS base

RUN apk add --update alpine-sdk

FROM base AS compiler
ADD . /src/

WORKDIR /src/

RUN make build

FROM base AS service

COPY --from=compiler /src/challenge /app/challenge
COPY --from=compiler /src/PortDomainService /app/PortDomainService

WORKDIR /app/

ENTRYPOINT [ "/app/PortDomainService" ]