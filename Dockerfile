

FROM		golang:1.17-bullseye AS builder

ARG			EXECUTABLE

WORKDIR		/go/src/app

ADD			. /go/src/app

RUN			go generate ./...		# gqlgen interface creation

RUN			go get -d -v ./...
RUN			go vet -v
RUN			go test -v

RUN			CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/app

RUN			mv /go/bin/app /go/bin/${EXECUTABLE}


FROM gcr.io/distroless/static
COPY --from=builder /go/bin/${EXECUTABLE} /
