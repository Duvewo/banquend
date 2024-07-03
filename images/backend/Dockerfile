FROM golang:alpine3.20@sha256:a8836ec73ab2f18e7f9abe18cdf2580b9575226e7dabeec3fc5230f8788aa9c4 AS runner

WORKDIR /src

COPY . /src/.

RUN go mod tidy && go mod download

RUN mkdir -p /app

RUN go build -o /app/backend cmd/server/main.go

WORKDIR /app

ENTRYPOINT [ "./backend" ]