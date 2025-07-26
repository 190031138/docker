FROM golang:alpine3.21 AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

FROM alpine:latest AS release

WORKDIR /app

COPY --from=build /app ./

EXPOSE 8080

CMD ["./main"]





