

FROM golang:1.17-buster AS builder
# FROM golang:1.17 AS builder

ARG	SERVICE
ARG	EXECUTABLE
ENV	SERVICE=${SERVICE}
ENV	EXECUTABLE=${EXECUTABLE}

WORKDIR	/go/src

COPY	go.mod ./
COPY	go.sum ./
RUN		go mod download
COPY	. .


RUN		go generate ./...
RUN		mkdir ../${SERVICE}

RUN		go build -o ../${SERVICE} ${EXECUTABLE}.go

# FROM golang:1.17-alpine
FROM golang:1.17
# FROM gcr.io/distroless/base-debian10
# FROM busybox

ARG	SERVICE
ARG	EXECUTABLE
ENV	SERVICE=${SERVICE}
ENV	EXECUTABLE=${EXECUTABLE}

WORKDIR /

RUN mkdir ${SERVICE}

COPY --from=builder /go/${SERVICE} /${SERVICE}

WORKDIR		/${SERVICE}
COPY		./docker-entrypoint.sh .
COPY		.env .
