# TASK-03: Banner Loading

**Agent:** previous | **LLM:** unknown

**Status:** COMPLETED
**Priority:** P0 (Critical)
**Golden Tests:** Infrastructure (enables GT-014+)

## Objective

Load and parse banner file into memory with error handling.

## Requirements

1. Read banner files from `banners/` directory (Updated in TASK-11)
2. Parse into string array (855 lines)
3. Handle file not found errors
4. Load once at startup

## TDD Workflow

### RED Phase

- Test loading a banner that doesn't exist.

### GREEN Phase

- Implement `LoadBanner` in `core/load_banner.go`.
- Use `os.ReadFile`.

### REFACTOR Phase

- Normalize line endings for cross-platform support.

## Acceptance Criteria

- [x] `core/load_banner.go` created.
- [x] Loads standard banner successfully.
- [x] Handles missing files gracefully.
