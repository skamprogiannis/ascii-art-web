# Stage 1 (builder):
FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY go.mod go.sum* ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux \
       go build \
         -trimpath \
         -ldflags='-s -w' \
         -o /app/server \
         .

# Stage 2 (final image):
FROM alpine:3.21
LABEL org.opencontainers.image.title="ascii-art-web-dockerize" \
         org.opencontainers.image.description="ASCII-Art web server in Docker" \
         org.opencontainers.image.version="1.0.0" \
         org.opencontainers.image.authors="your-names" 


# Security: Run as a non-root user
RUN apk add --no-cache bash && adduser -D appuser

WORKDIR /app
COPY --from=builder /app/server ./
COPY --from=builder /app/banners ./banners/
COPY --from=builder /app/templates ./templates/
COPY --from=builder /app/static ./static/

USER appuser
EXPOSE 8080
CMD ["./server"]

