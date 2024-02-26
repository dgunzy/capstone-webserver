FROM golang:1.21 AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o myapp .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/myapp .

COPY --from=builder /app/routing /root/routing

EXPOSE 8080

CMD ["./myapp"]
