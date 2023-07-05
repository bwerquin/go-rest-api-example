FROM golang:latest

RUN mkdir -p /app

COPY . /app

WORKDIR /app

RUN go build main.go

CMD ["/app/main"]