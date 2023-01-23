FROM golang:latest

WORKDIR /go/src/packages-api

COPY . .

RUN go build -o packages-api

EXPOSE 8080

CMD ["./packages-api"]