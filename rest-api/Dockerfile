FROM golang:1.19-alpine as builder
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN go build -o everybody

FROM alpine
WORKDIR /app
COPY --from=builder /app/everybody .
EXPOSE 5100
CMD ["./everybody"]
