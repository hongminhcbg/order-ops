FROM golang:1.14 as builder

WORKDIR /app

COPY . /app

RUN go mod download

RUN GOOS=linux

RUN go build -o main ./main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 80

CMD ["/app/main"]