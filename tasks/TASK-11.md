# TASK-11: Banner Path Integration

**Status:** COMPLETED
**Priority:** P1 (High)
**Dependencies:** TASK-10
**Estimated Time:** 20 minutes

## Objective

Update the code to load banner files from the `banners/` directory instead of the root.

## Requirements

1. Update `core/load_banner.go` to search in the `banners/` directory.
2. Update unit tests in `core/load_banner_test.go` to reflect the new path.
3. Ensure `main.go` correctly passes the default banner filename.

## TDD Workflow

### RED Phase

- Run `go test ./core/...` and verify `TestLoadBanner` fails because it cannot find the file in root.

### GREEN Phase

- Update `core/load_banner.go`:
  ```go
  func loadBanner(filename string) ([]string, error) {
      path := "banners/" + filename
      data, err := os.ReadFile(path)
      // ...
  }
  ```
- Update unit tests to use valid filenames in `banners/`.

### REFACTOR Phase

- Verify all unit tests pass.

## Acceptance Criteria

- [x] `LoadBanner("standard.txt")` works.
- [x] `LoadBanner("shadow.txt")` works.
- [x] `LoadBanner("thinkertoy.txt")` works.
- [x] Unit tests pass.
- [x] Standard rendering still works with the new path.
