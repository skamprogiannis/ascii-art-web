#!/usr/bin/env bash
set -euo pipefail

IMAGE_NAME="ascii-art-web-docker"
CONTAINER_NAME="dockerize"
PORT="${PORT:-8080}"

case "${1:-}" in
 build)
   docker image build -t "$IMAGE_NAME" .
   ;;
 run)
   docker rm -f "$CONTAINER_NAME" 2>/dev/null || true
   docker container run -d -p "$PORT:8080" --name "$CONTAINER_NAME" "$IMAGE_NAME"
   ;;
 restart)
   docker rm -f "$CONTAINER_NAME" 2>/dev/null || true
   docker image build -t "$IMAGE_NAME" .
   docker container run -d -p "$PORT:8080" --name "$CONTAINER_NAME" "$IMAGE_NAME"
   ;;
 shell)
   docker exec -it "$CONTAINER_NAME" /bin/bash
   ;;
 stop)
   docker rm -f "$CONTAINER_NAME"
   ;;
 *)
   echo "Usage: $0 {build|run|restart|shell|stop}"
   exit 1
   ;;
esac
