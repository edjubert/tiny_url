FROM golang:1.23-alpine AS build_base

WORKDIR /src

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o ./out/app ./cmd/tiny_url/main.go

FROM alpine:3.12

WORKDIR /app

COPY --from=build_base /src/out/app /app/tiny_url
COPY --from=build_base /src/data /app/data
COPY --from=build_base /src/config/ /app/config

RUN chmod +x tiny_url

EXPOSE 3000

ENTRYPOINT ["./tiny_url"]
