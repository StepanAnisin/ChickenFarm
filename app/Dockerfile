FROM golang:latest
RUN mkdir /app
COPY . /app
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod tidy
RUN go build -o main .
CMD ["/app/main"]