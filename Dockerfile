

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

ARG			SERVICE
ARG			EXECUTABLE
ENV			SERVICE=${SERVICE}
ENV			EXECUTABLE=${EXECUTABLE}

WORKDIR		/go/${SERVICE}

COPY		--from=builder /go/src/app/${SERVICE} .
RUN			mv .secrets /secrets/

RUN			chmod -R 777 /go

ENTRYPOINT	["sh", "-c", "/go/graphql/docker-entrypoint.sh"]
