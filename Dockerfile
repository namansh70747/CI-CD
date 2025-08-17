FROM golang:1.25-alpine AS builder
WORKDIR /app

# Install certificates and git (if needed for modules)
RUN apk add --no-cache git ca-certificates

# Copy and build
COPY . .
RUN go build -o app main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/app .
EXPOSE 8080
CMD ["./app"]
