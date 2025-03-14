FROM golang:1.24-alpine AS build

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download


COPY . .

RUN apk add --no-cache make

RUN go build -o main main.go

FROM alpine:latest AS runner

WORKDIR /app

RUN apk update && apk add --no-cache curl

COPY --from=build /build/main ./main

CMD ["/app/main"]