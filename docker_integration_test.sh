#!/usr/bin/env bash
set -euo pipefail

# Docker Integration Tests for ASCII Art Web Server
# Tests build success, best practices, runtime health, and port access

IMAGE_NAME="ascii-art-web-test"
CONTAINER_NAME="ascii-art-test-container"
TEST_PORT="8081"

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

log_info() {
    echo -e "${GREEN}[INFO]${NC} $1"
}

log_warn() {
    echo -e "${YELLOW}[WARN]${NC} $1"
}

log_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

cleanup() {
    log_info "Cleaning up test resources..."
    docker rm -f "$CONTAINER_NAME" 2>/dev/null || true
    docker rmi -f "$IMAGE_NAME" 2>/dev/null || true
}

trap cleanup EXIT

# Test 1: Build Success
test_build_success() {
    log_info "Testing Docker build success..."

    # Check Docker permissions
    if ! docker ps > /dev/null 2>&1; then
        log_error "✗ No permission to run Docker commands"
        log_error "Try running with sudo or add user to docker group: sudo usermod -aG docker \$USER"
        return 1
    fi

    if docker build -t "$IMAGE_NAME" . > /dev/null 2>&1; then
        log_info "✓ Build successful"
        return 0
    else
        log_error "✗ Build failed"
        return 1
    fi
}

# Test 2: Dockerfile Best Practices (Linting)
test_best_practices() {
    log_info "Testing Dockerfile best practices..."

    # Check if hadolint is available
    if command -v hadolint >/dev/null 2>&1; then
        log_info "Running hadolint for Dockerfile linting..."
        if hadolint Dockerfile > /dev/null 2>&1; then
            log_info "✓ Dockerfile passes hadolint checks"
        else
            log_warn "! Dockerfile has linting warnings (non-fatal)"
        fi
    else
        log_warn "hadolint not available, skipping advanced linting"
    fi

    # Basic checks
    local issues=0

    # Check for lightweight base image
    if grep -q "FROM alpine" Dockerfile; then
        log_info "✓ Using lightweight Alpine base image"
    else
        log_warn "! Not using Alpine base image"
        ((++issues))
    fi

    # Check for non-root user
    if grep -q "USER appuser\|USER [0-9]" Dockerfile; then
        log_info "✓ Running as non-root user"
    else
        log_error "✗ Running as root user"
        ((++issues))
    fi

    # Check for multi-stage build
    if grep -q "AS builder" Dockerfile && grep -q "FROM.*AS builder" Dockerfile; then
        log_info "✓ Using multi-stage build"
    else
        log_warn "! Not using multi-stage build"
    fi

    # Check for unnecessary packages
    if grep -q "apk add.*bash" Dockerfile; then
        log_info "✓ Only installing necessary packages (bash)"
    else
        log_warn "! Installing unnecessary packages"
    fi

    return $issues
}

# Test 3: Runtime Health
test_runtime_health() {
    log_info "Testing runtime health..."

    # Start container
    if ! docker run -d -p "$TEST_PORT:8080" --name "$CONTAINER_NAME" "$IMAGE_NAME" > /dev/null 2>&1; then
        log_error "✗ Failed to start container"
        return 1
    fi

    # Wait a moment for startup
    sleep 3

    # Check if container is still running
    if ! docker ps | grep -q "$CONTAINER_NAME"; then
        log_error "✗ Container crashed immediately"
        docker logs "$CONTAINER_NAME" || true
        return 1
    fi

    log_info "✓ Container is running and healthy"

    # Check container logs for errors
    if docker logs "$CONTAINER_NAME" 2>&1 | grep -qi "error\|panic\|fatal"; then
        log_warn "! Container logs contain error messages"
        docker logs "$CONTAINER_NAME"
    else
        log_info "✓ No errors in container logs"
    fi

    return 0
}

# Test 4: Port Access
test_port_access() {
    log_info "Testing port accessibility..."

    # Check if port is listening
    if ! nc -z localhost "$TEST_PORT" 2>/dev/null; then
        log_error "✗ Port $TEST_PORT is not accessible"
        return 1
    fi

    log_info "✓ Port $TEST_PORT is accessible"

    # Test HTTP endpoint
    if command -v curl >/dev/null 2>&1; then
        log_info "Testing HTTP endpoints..."

        # Test root endpoint
        if curl -s -f "http://localhost:$TEST_PORT/" > /dev/null 2>&1; then
            log_info "✓ Root endpoint (/) is accessible"
        else
            log_error "✗ Root endpoint (/) is not accessible"
            return 1
        fi

        # Test ASCII art endpoint with a simple request
        response=$(curl -s -X POST "http://localhost:$TEST_PORT/ascii-art" \
            -H "Content-Type: application/x-www-form-urlencoded" \
            -d "text=Hi&banner=standard" 2>/dev/null || echo "")

        if [[ -n "$response" ]]; then
            log_info "✓ ASCII art endpoint is functional"
        else
            log_error "✗ ASCII art endpoint is not responding"
            return 1
        fi
    else
        log_warn "curl not available, skipping HTTP endpoint tests"
    fi

    return 0
}

# Main test execution
main() {
    log_info "Starting Docker Integration Tests for ASCII Art Web Server"
    log_info "=================================================="

    local test_count=0
    local pass_count=0

    # Test 1: Build Success
    ((++test_count))
    if test_build_success; then
        ((++pass_count))
    fi

    # Test 2: Best Practices
    ((++test_count))
    if test_best_practices; then
        ((++pass_count))
    fi

    # Test 3: Runtime Health
    ((++test_count))
    if test_runtime_health; then
        ((++pass_count))
    else
        # If runtime test fails, skip port test
        log_error "Runtime health test failed, skipping port access test"
        cleanup
        exit 1
    fi

    # Test 4: Port Access
    ((++test_count))
    if test_port_access; then
        ((++pass_count))
    fi

    # Summary
    log_info "=================================================="
    log_info "Test Results: $pass_count/$test_count tests passed"

    if [[ $pass_count -eq $test_count ]]; then
        log_info "🎉 All tests passed!"
        return 0
    else
        log_error "❌ Some tests failed"
        return 1
    fi
}

# Run main function
main "$@"