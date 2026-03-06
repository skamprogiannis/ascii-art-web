# TASK-18: Internal Refactor & Library Package

**Status:** COMPLETED
**Priority:** P1 (High)
**Dependencies:** TASK-16, TASK-17
**Estimated Time:** 30 minutes

## Objective

Refactor `core/` into idiomatic internal packages and add a public `asciiart` library package while preserving CLI behavior.

## Requirements

1. Move business logic into `internal/` packages (`parser`, `banner`, `render`, `color`)
2. Keep CLI usage `go run .` intact
3. Add `asciiart` package for library usage
4. Update tests to new package locations
5. Update docs to reflect new structure
6. Bump documentation version to v1.3

## TDD Workflow

### RED Phase

- Identify all `core/` imports and file references

### GREEN Phase

- Moved parsing to `internal/parser`
- Moved banner loading to `internal/banner`
- Moved rendering to `internal/render`
- Moved color parsing to `internal/color`
- Added `asciiart` package with `RenderString` API
- Updated `main.go` to use internal packages
- Updated unit tests to new locations

### REFACTOR Phase

- Removed obsolete `core/` directory
- Updated docs and task index to new structure
- Bumped PRD/architecture/golden tests to v1.3

## Acceptance Criteria

- [x] CLI works with `go run .`
- [x] Library API available via `asciiart.RenderString`
- [x] Tests pass with `go test ./...`
- [x] Docs updated to new package layout
