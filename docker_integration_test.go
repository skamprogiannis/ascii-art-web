package main

import (
	"os/exec"
	"strings"
	"testing"
	"time"
)

// TestDockerIntegration runs comprehensive Docker integration tests
func TestDockerIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping Docker integration tests in short mode")
	}

	// Check if Docker is available
	if _, err := exec.LookPath("docker"); err != nil {
		t.Skip("Docker not available, skipping integration tests")
	}

	// Check Docker permissions
	cmd := exec.Command("docker", "ps")
	if err := cmd.Run(); err != nil {
		t.Skip("No permission to run Docker commands, skipping integration tests")
	}

	t.Log("Running Docker integration tests...")

	// Run the integration test script
	cmd = exec.Command("./docker_integration_test.sh")
	output, err := cmd.CombinedOutput()

	// Log the output for debugging
	t.Logf("Integration test output:\n%s", string(output))

	if err != nil {
		t.Fatalf("Docker integration tests failed: %v", err)
	}

	// Check for test failures in output
	if strings.Contains(string(output), "Some tests failed") ||
		strings.Contains(string(output), "❌") {
		t.Error("One or more Docker integration tests failed")
	}

	// Verify success message
	if !strings.Contains(string(output), "All tests passed") &&
		!strings.Contains(string(output), "🎉") {
		t.Error("Docker integration tests did not complete successfully")
	}
}

// TestDockerBuildOnly tests just the Docker build process
func TestDockerBuildOnly(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping Docker build test in short mode")
	}

	// Check if Docker is available
	if _, err := exec.LookPath("docker"); err != nil {
		t.Skip("Docker not available, skipping build test")
	}

	t.Log("Testing Docker build...")

	// Build the image
	cmd := exec.Command("docker", "build", "-t", "ascii-art-web-test-build", ".")
	output, err := cmd.CombinedOutput()

	if err != nil {
		t.Fatalf("Docker build failed: %v\nOutput: %s", err, string(output))
	}

	t.Log("✓ Docker build successful")

	// Clean up the test image
	go func() {
		time.Sleep(1 * time.Second) // Give a moment for any cleanup
		exec.Command("docker", "rmi", "-f", "ascii-art-web-test-build").Run()
	}()
}
