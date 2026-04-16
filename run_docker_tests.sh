#!/usr/bin/env bash
set -euo pipefail

# Docker Integration Test Runner
# Handles Docker permission issues automatically

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
TEST_SCRIPT="$SCRIPT_DIR/docker_integration_test.sh"

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m'

log_info() {
    echo -e "${GREEN}[INFO]${NC} $1"
}

log_warn() {
    echo -e "${YELLOW}[WARN]${NC} $1"
}

log_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# Check if Docker is available
if ! command -v docker >/dev/null 2>&1; then
    log_error "Docker is not installed or not in PATH"
    exit 1
fi

# Check Docker permissions
if docker ps >/dev/null 2>&1; then
    log_info "Docker permissions OK, running tests..."
    exec "$TEST_SCRIPT"
else
    log_warn "No Docker permissions detected"

    # Check if we can use sudo
    if command -v sudo >/dev/null 2>&1 && sudo -n true >/dev/null 2>&1 2>&1; then
        log_info "Running tests with sudo..."
        exec sudo "$TEST_SCRIPT"
    else
        log_error "Cannot run Docker commands. Please:"
        echo "  1. Add your user to the docker group: sudo usermod -aG docker \$USER"
        echo "     Then logout and login again, or run: newgrp docker"
        echo "  2. Or run this script with sudo manually: sudo $TEST_SCRIPT"
        exit 1
    fi
fi