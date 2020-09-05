# Build the application
FROM golang:1.15-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build \
        -o rhttp \
        ./cmd/rhttp/main.go

# Run the application in an empty alpine environment
FROM alpine:latest
RUN apk update && apk upgrade && \
    apk add --no-cache git
WORKDIR /root
COPY --from=build /app/rhttp .
CMD ["./rhttp"]