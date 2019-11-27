FROM golang:1.12.0-alpine3.9 as build

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o main .

EXPOSE 8080

CMD ["/main"]
