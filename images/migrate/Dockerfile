FROM golang:alpine3.20@sha256:a8836ec73ab2f18e7f9abe18cdf2580b9575226e7dabeec3fc5230f8788aa9c4

RUN apk add git

WORKDIR /migrate

COPY /storage/migrate/. /migrate/

RUN go install github.com/pressly/goose/v3/cmd/goose@v3.21.1

ENTRYPOINT [ "goose", "up" ]