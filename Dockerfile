FROM golang:alpine

WORKDIR /app

COPY . .
RUN go mod download

RUN go build -o limero ./cmd/limero/main.go

EXPOSE 7920

ENTRYPOINT [ "/app/limero" ]


