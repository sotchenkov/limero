FROM golang:1.22  as builder

WORKDIR /app
COPY . .

ENV CGO_ENABLED 0
ENV GOOS linux

RUN go mod download
RUN go build -o limero ./cmd/limero/main.go


FROM alpine:3.19

WORKDIR /app
COPY --from=builder /app/limero .

EXPOSE 7920

ENTRYPOINT [ "/app/limero" ]