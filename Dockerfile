FROM golang:1.23-alpine

WORKDIR /app

COPY . .

RUN go mod tidy && go build -o server

EXPOSE 8095

CMD ["./server"]