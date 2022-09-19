FROM golang:1.12.0-alpine3.9
RUN mkdir /app
COPY . /app
WORKDIR /app
RUN apk add git
RUN go run main.go
CMD ["/app/main"]