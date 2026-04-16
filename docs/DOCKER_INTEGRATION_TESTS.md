# Docker Integration Tests

This directory contains comprehensive integration tests for the Docker containerization of the ASCII Art Web Server.

## Test Coverage

The integration tests verify four critical aspects of the Docker deployment:

### 1. Build Success ✅
- **What**: Verifies that the Dockerfile builds successfully without errors
- **How**: Runs `docker build` and checks for successful completion
- **Importance**: Ensures the container can be built in any environment

### 2. Best Practices (Linting) ✅
- **What**: Checks Dockerfile follows Docker best practices
- **Checks**:
  - Uses lightweight base image (Alpine)
  - Runs as non-root user for security
  - Implements multi-stage build for smaller images
  - Only installs necessary packages
- **Tools**: Uses `hadolint` if available, falls back to basic checks
- **Importance**: Ensures efficient, secure, and maintainable containers

### 3. Runtime Health ✅
- **What**: Verifies the container runs without crashing
- **How**: Starts the container and monitors if it stays alive
- **Checks**:
  - Container starts successfully
  - Container doesn't exit immediately
  - No critical errors in logs
- **Importance**: Ensures the application runs properly in containerized environment

### 4. Port Access ✅
- **What**: Verifies the web server is accessible from outside the container
- **How**: Makes HTTP requests to the exposed port
- **Checks**:
  - Port 8080 is accessible
  - Root endpoint (`/`) responds
  - ASCII art endpoint (`/ascii-art`) functions
- **Importance**: Confirms the web service works end-to-end

## Running the Tests

### Option 1: Automated Runner (Recommended)
```bash
# Automatically handles Docker permissions
./run_docker_tests.sh

# This script will:
# - Check Docker availability and permissions
# - Use sudo if needed
# - Run all integration tests
# - Provide clear error messages
```

### Option 2: Direct Bash Script
```bash
# Run all integration tests directly
./docker_integration_test.sh

# Requires Docker permissions to be set up manually
```

### Option 3: Go Test Integration
```bash
# Run as part of Go test suite
go test -run TestDockerIntegration

# Run only build test (faster)
go test -run TestDockerBuildOnly

# Run all tests including integration
go test ./...

# Skip integration tests in CI (use -short flag)
go test -short ./...
```

## Prerequisites

- Docker installed and running
- **Docker permissions**: User must be able to run Docker commands (either sudo or in docker group)
- `curl` and `nc` (netcat) for port testing (usually pre-installed on Linux/macOS)
- `hadolint` for advanced Dockerfile linting (optional, will use basic checks if not available)

### Docker Permission Setup

If you get permission errors, run one of these commands:

```bash
# Option 1: Add user to docker group (recommended)
sudo usermod -aG docker $USER
# Then logout and login again, or run: newgrp docker

# Option 2: Run tests with sudo
sudo ./docker_integration_test.sh

# Option 3: Run Go tests with sudo
sudo go test -run TestDockerIntegration
```

## Test Output

The tests provide colored output indicating:
- ✅ **Green**: Tests passed
- ⚠️ **Yellow**: Warnings (non-fatal)
- ❌ **Red**: Test failures

Example successful run:
```
[INFO] Starting Docker Integration Tests for ASCII Art Web Server
==================================================
[INFO] Testing Docker build success...
[INFO] ✓ Build successful
[INFO] Testing Dockerfile best practices...
[INFO] ✓ Using lightweight Alpine base image
[INFO] ✓ Running as non-root user
[INFO] ✓ Using multi-stage build
[INFO] Testing runtime health...
[INFO] ✓ Container is running and healthy
[INFO] Testing port accessibility...
[INFO] ✓ Port 8081 is accessible
[INFO] ✓ Root endpoint (/) is accessible
[INFO] ✓ ASCII art endpoint is functional
==================================================
[INFO] Test Results: 4/4 tests passed
[INFO] 🎉 All tests passed!
```

## Troubleshooting

### Common Issues

1. **Docker not available**
   - Install Docker Desktop or Docker Engine
   - Ensure Docker daemon is running

2. **Port already in use**
   - Tests use port 8081 to avoid conflicts with port 8080
   - Stop any services using port 8081

3. **Build failures**
   - Check Dockerfile syntax
   - Ensure all required files are present
   - Verify Go modules are properly configured

4. **Runtime failures**
   - Check container logs: `docker logs ascii-art-test-container`
   - Verify banner files and templates are copied correctly
   - Ensure the Go binary was built successfully

### Manual Testing

If automated tests fail, you can test components manually:

```bash
# Build image
./docker.sh build

# Run container
./docker.sh run

# Check if running
docker ps | grep dockerize

# Test endpoints
curl http://localhost:8080/
curl -X POST http://localhost:8080/ascii-art -d "text=Hi&banner=standard"

# View logs
./docker.sh shell  # Then check logs or run commands inside container
```

## CI/CD Integration

For continuous integration, add to your pipeline:

```yaml
# GitHub Actions example
- name: Run Docker Integration Tests
  run: |
    chmod +x docker_integration_test.sh
    ./docker_integration_test.sh
```

Or using Go tests:
```yaml
- name: Run Go Tests with Docker Integration
  run: go test -v ./...
```

## Architecture Notes

The tests are designed to be:
- **Isolated**: Use unique container/image names to avoid conflicts
- **Clean**: Automatically clean up test resources
- **Comprehensive**: Cover all aspects of container deployment
- **Fast**: Build test can run separately for quick feedback
- **Reliable**: Include proper error handling and timeouts