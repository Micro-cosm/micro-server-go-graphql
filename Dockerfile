

FROM	golang:1.17 AS builder

ARG		SERVICE
ARG		EXECUTABLE
ENV		SERVICE=${SERVICE}
ENV		EXECUTABLE=${EXECUTABLE}

WORKDIR	/go/src/app/${SERVICE}
COPY	. .

RUN		go get ./...
RUN		go generate ./...
RUN		go build -o ${EXECUTABLE} ${SERVICE}.go


# FROM	golang:1.17-alpine
FROM	golang:1.17

COPY	--from=builder /go/src/app/${SERVICE} .

WORKDIR	/go/${SERVICE}

RUN		chmod -R 777 *