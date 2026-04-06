# TASK-27: Decouple Web Server from Public asciiart Package

**Status:** COMPLETED
**Priority:** P2 (Medium)
**Dependencies:** TASK-23
**Estimated Time:** 15 minutes

## Objective
Refactor the web server so `main.go` uses shared internal generation logic directly instead of importing the public `asciiart` package, while preserving the public library API.

## Requirements
1. Keep the public `asciiart` package available for library consumers.
2. Remove `main.go`'s direct dependency on `asciiart`.
3. Centralize shared generation logic in an internal package rather than duplicating it in `main`.
4. Preserve existing HTTP and JSON rendering behavior.

## Implementation Steps
- ✅ Added `internal/generator` to hold shared ASCII art generation logic.
- ✅ Updated `main.go` to call `internal/generator.RenderBundle` directly.
- ✅ Kept `asciiart` as a thin public wrapper over the shared internal implementation.
- ✅ Re-ran tests and verified the build path no longer depends on copying `asciiart/`.

## Acceptance Criteria
- [x] `main.go` no longer imports `ascii-art-web/asciiart`.
- [x] Shared generation logic exists in an internal package.
- [x] Public `asciiart` package still works as a wrapper API.
- [x] Tests pass with `CGO_ENABLED=0 go test ./...`.
- [x] Build succeeds from a Docker-equivalent source set that excludes `asciiart/`.
