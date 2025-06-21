FROM golang:1.24.3-alpine3.21 AS builder

WORKDIR /app
COPY . /app
RUN CGO_ENABLED=0 go build .

FROM alpine:3.21

WORKDIR /app
COPY --from=builder /app/glance .
COPY ./config /app/config
COPY ./assets /app/assets

EXPOSE 8080
ENTRYPOINT ["/app/glance"]
