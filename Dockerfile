# Stage 1 (builder):
FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY go.mod go.sum* ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o server .

# Stage 2 (final image):
FROM alpine:3.21
# Security: Run as a non-root user
RUN adduser -D appuser
USER appuser
WORKDIR /app
COPY --from=builder /app/server ./
COPY --from=builder /app/banners ./banners/
COPY --from=builder /app/templates ./templates/
COPY --from=builder /app/static ./static/

EXPOSE 8080
CMD ["./server"]

