FROM golang:1 AS build-env
WORKDIR /go
ADD main.go /go
ARG CGO_ENABLED=0
ARG OPTS
RUN env ${OPTS} CGO_ENABLED=${CGO_ENABLED} go build -o 301

FROM scratch
COPY --from=build-env /go/301 /301
EXPOSE 80/tcp
ENTRYPOINT ["/301"]
