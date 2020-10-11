FROM golang:1.14 as builder

WORKDIR /app

COPY . /app

RUN go mod download

RUN GOOS=linux

RUN go build -o main

FROM ubuntu:16.04

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 80

CMD ["/app/main"]