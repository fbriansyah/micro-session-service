FROM golang:1.19-alpine3.16 AS builder
WORKDIR /app
COPY . .
RUN go build -o session-microservice ./cmd/

FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/session-microservice .
COPY app.env .

EXPOSE 9090

ENTRYPOINT [ "/app/session-microservice" ]