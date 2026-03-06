# TASK-12: Optional Banner Selection

**Status:** COMPLETED
**Priority:** P1 (High)
**Dependencies:** TASK-11
**Estimated Time:** 30 minutes

## Objective

Support an optional second argument to select the banner style (standard/shadow/thinkertoy).

## Requirements

1. Update `main.go` to accept 1 or 2 arguments.
2. If 2 arguments are provided:
   - `os.Args[1]` is the input string.
   - `os.Args[2]` is the banner style.
3. If 1 argument is provided:
   - Default to `standard.txt`.
4. Validate that the requested banner exists.
5. Exit with usage message if more than 2 arguments are provided.
6. Support styles by name (e.g., `standard`, `shadow`, `thinkertoy`) or with `.txt` extension.

## TDD Workflow

### RED Phase

- Run `go run . "Hello" "shadow"` and verify it fails (either prints usage or fails to load).

### GREEN Phase

- Update `main.go`:

  ```go
  if len(os.Args) < 2 || len(os.Args) > 3 {
      fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
      return
  }

  bannerFile := "standard.txt"
  if len(os.Args) == 3 {
      bannerFile = os.Args[2]
      if !strings.HasSuffix(bannerFile, ".txt") {
          bannerFile += ".txt"
      }
  }
  ```

- Load the selected banner file.
- Process input normally.

### REFACTOR Phase

- Clean up the argument parsing logic.
- Ensure all 30 golden tests pass with default `standard`.
- Verify `shadow` and `thinkertoy` work as expected.

## Acceptance Criteria

- [x] `go run . "Hello" shadow` works.
- [x] `go run . "Hello" thinkertoy` works.
- [x] `go run . "Hello"` defaults to standard.
- [x] Invalid banner names are handled with a clear error message.
- [x] More than 2 arguments trigger the usage message.
- [x] AI_LOG.md updated.
